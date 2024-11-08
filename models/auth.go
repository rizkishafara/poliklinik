package models

import (
	"log"
	"poliklinik/db"
	"poliklinik/utils"
)

func Register(username, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	passnew, err := utils.HashPassword(password)
	if err != nil {
		log.Println("error hash password", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	_, err = dbEngine.QueryString("INSERT INTO users (username, password) VALUES (?,?)", username, passnew)

	if err != nil {
		log.Println("error insert user", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
