package ws

import (
	"CarDemo1/conf"
	"CarDemo1/model"
	. "CarDemo1/service/ws/model"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

func MessageIndex(c *gin.Context) {
	//id := "->1"
	//log.Println("RUSH")
	id:="->"+c.Param("id")
	//fmt.Println(id)
	db := conf.MongoDBClient.Database(conf.MongoDBName)
	names, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "\"code: 500\", \"msg\": \"error\"")
	}
	//fmt.Println(names)		// [1->2 2->1 4->1 2->4 3->1 2->3]
	var connectNames []string
	for _, name := range names{
		if strings.HasSuffix(name, id){
			connectNames = append(connectNames, name)
		}
	}
	//fmt.Println(connectNames)	// [2->1 4->1 3->1]
	type TrainerId struct {
		Trainer Trainer		`json:"msg"`
		Id		string		`json:"id"`
		AvatarUrl string 	`json:"avatar_url"`
		UserName string `json:"user_name"`
	}
	//var trainers []Trainer
	var ans []TrainerId
	for _, name := range connectNames{
		var user model.User
		trainer, id, err := FirstUnread(conf.MongoDBName, name)
		model.DB.First(&user, id)
		newTrainerId := TrainerId{
			Trainer: trainer,
			Id:      id,
			AvatarUrl:user.Avatar,
			UserName:user.UserName,
		}
		if err == nil{
			ans = append(ans, newTrainerId)
		}else {
			continue
		}
	}
	c.JSON(http.StatusOK, ans)
}