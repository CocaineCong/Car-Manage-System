package model

import (
	"CarDemo1/conf"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
	"strings"
	"time"
)

type Trainer struct {
	Content 	string	`bson:"content"`	//内容
	StartTime 	int64 	`bson:"startTime"` 	//创建时间
	EndTime   	int64 	`bson:"endTime"`   	//过期时间
	Read		uint	`bson:"read"`		//已读?
}

type Result struct {
	StartTime 	int64
	Msg			string
	From		string
}

func InsertOne(database string, id string, content string, read uint, expire int64) (err error) {
	collection := conf.MongoDBClient.Database(database).Collection(id)
	// 插入一条
	comment := Trainer{content, time.Now().Unix(), time.Now().Unix() + expire, read}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}


func FindMany(database string, sendId string, id string, time int64, pageSize int) (results []Result, err error) {
	// 查询多条（分页）

	var resultsMe []Trainer
	var resultsYou []Trainer
	sendIdCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	idCollection := conf.MongoDBClient.Database(database).Collection(id)
	filter := bson.M{"startTime": bson.M{"$lt": time}}
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	idTimeCursor, err := idCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	//conf.MongoDBClient.Database(database).Collection(sendId).Find(context.TODO(), filter)
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou)		// sendId 对面发过来的
	err = idTimeCursor.All(context.TODO(), &resultsMe)			// Id 发给对面的
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return
}

func FirstFind(database string, sendId string, id string) (results []Result, err error) {
	// 首次查询（把所有未读都取出来--对方发过来的）
	var resultsMe []Trainer
	var resultsYou []Trainer
	sendIdCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	idCollection := conf.MongoDBClient.Database(database).Collection(id)
	filter := bson.M{"read": bson.M{"$all": []uint{0}}}
	sendIdCursor, err := sendIdCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"startTime", 1}}), options.Find().SetLimit(1))
	if sendIdCursor == nil {
		return
	}
	var unreads []Trainer
	err = sendIdCursor.All(context.TODO(), &unreads)
	if err != nil {
		log.Println(err)
	}
	if len(unreads) > 0{
		timeFilter := bson.M{"startTime": bson.M{"$gte": unreads[0].StartTime}}		// 最旧的未读信息的时间戳为起点全部返回
		sendIdTimeCursor, _ := sendIdCollection.Find(context.TODO(), timeFilter)
		idtimeCursor, _ := idCollection.Find(context.TODO(), timeFilter)
		err = sendIdTimeCursor.All(context.TODO(), &resultsYou)		// sendId 对面发过来的
		err = idtimeCursor.All(context.TODO(), &resultsMe)			// Id 发给对面的
		results, err = AppendAndSort(resultsMe, resultsYou)
	}else {
		results, err = FindMany(database, sendId, id, 9999999999, 10)
	}
	overTimeFilter := bson.D{
		{"$and", bson.A{
			bson.D{{"endTime", bson.M{"$lt": time.Now().Unix()}}},
			bson.D{{"read", bson.M{"$eq": 1}}},
		}},
	}
	_, _ = sendIdCollection.DeleteMany(context.TODO(), overTimeFilter)
	_, _ = idCollection.DeleteMany(context.TODO(), overTimeFilter)
	// 将所有未读设置为已读
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{"$set": bson.M{"read": 1}})
	// 刷新过期时间
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{"$set": bson.M{"endTime": time.Now().Unix() + int64(60*60*24*30*3)}})
	return
}

func AppendAndSort(resultsMe []Trainer, resultsYou []Trainer) (results []Result, err error) {
	for _, r := range resultsMe{
		result := Result{
			StartTime: r.StartTime,
			Msg:  fmt.Sprintf("{\"content\": \"%s\", \"code\": %d, \"create_at\": %d}", r.Content, r.Read, r.StartTime),
			From: "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou{
		result := Result{
			StartTime: r.StartTime,
			Msg:  fmt.Sprintf("{\"content\": \"%s\", \"code\": %d, \"create_at\": %d}", r.Content, r.Read, r.StartTime),
			From: "you",
		}
		results = append(results, result)
	}
	// 最后排个序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}

func FirstUnread(database string, name string) (trainer Trainer, id string, err error)  {
	collection := conf.MongoDBClient.Database(database).Collection(name)
	filter := bson.M{"read": bson.M{"$all": []uint{0}}}

	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(1))
	var unreads []Trainer
	err = cursor.All(context.TODO(), &unreads)
	if len(unreads) > 0 {
		trainer = unreads[0]
	}else{
		err = errors.New("0")
	}
	id = strings.Split(name, "->")[0]

	return
}