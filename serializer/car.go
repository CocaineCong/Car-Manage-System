package serializer

import "CarDemo1/model"

type Car struct {
	ID 		uint 	`json:"id"`
	CarName	string  `json:"car_name"`
	CarNum 		string  `json:"car_num"`
	CarImages	string  `json:"car_images"`
	CarBoss    string  `json:"car_boss"`
	CarBossId	uint  `json:"car_boss_id"`
	CarBossName string `json:"car_boss_name"`
}

//序列化车辆
func BuildCar(item model.Car) Car {
	return Car{
		ID : item.ID,
		CarName : item.CarName,
		CarNum : item.CarNum,
		CarImages : item.CarImages,
		CarBossId : item.CarBossId,
		CarBossName:item.CarBossName,
	}
}

//车辆列表
func BuildCars(items []model.Car) (Cars []Car) {
	for _, item := range items {
		Car := BuildCar(item)
		Cars = append(Cars, Car)
	}
	return Cars
}
