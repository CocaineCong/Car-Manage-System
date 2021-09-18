package model

import (
	"CarDemo1/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Society struct {
	gorm.Model
	CategoryID  uint
	CategoryName string
	EnglishName string
	Title 	string
	Content string
	Picture string
	UserID  	uint
	UserName 	string
	UserAvatar  string
	Price 	string
	Status 	string  // 0 表示未交易  1 表示已交易  2 表示交易完成
}

func (society *Society) View() uint64 {
	contStr,_ := cache.RedisClient.Get(cache.ProductViewKey(society.ID)).Result()
	count,_:=strconv.ParseUint(contStr,10,64)
	return count
}

func (society *Society) AddView() {
	cache.RedisClient.Incr(cache.ProductViewKey(society.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(society.ID)))
}


