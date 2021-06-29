package migrate

import (
	"hotel/connect"
	"hotel/model"
)

func CreateTable() {
	db := connect.Connect()
	db.AutoMigrate(&model.Role{}, &model.User{}, &model.Authentication{}, &model.Conversation{}, &model.Messenger{}, &model.ImageHotel{}, &model.ImageRoom{})
	db.AutoMigrate(&model.Option{}, &model.Times{}, &model.Price{}, &model.Room{}, &model.Hotel{}, &model.Bill{}, &model.Rate{})
}
