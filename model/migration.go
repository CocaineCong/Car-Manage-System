package model

//执行数据迁移

func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Information{}).
		AutoMigrate(&Society{}).
		AutoMigrate(&Comment{}).
		AutoMigrate(&Report{}).
		AutoMigrate(&Category{}).
		AutoMigrate(&Carousel{}).
		AutoMigrate(&Car{}).
		AutoMigrate(&Type{}).
		AutoMigrate(&Admin{})
}
