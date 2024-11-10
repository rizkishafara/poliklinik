package models

import (
	"log"
	"poliklinik/db"
	"poliklinik/utils"
)

func Login(username, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datauser, err := dbEngine.QueryString("SELECT id, username FROM users WHERE username = ? ", username)

	if err != nil {
		log.Println("error get user", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	if datauser == nil {
		Respon.Status = 404
		Respon.Message = "user not found"
		return Respon
	}

	datares := make(map[string]interface{})
	if datauser != nil {
		datares["username"] = datauser[0]["username"]
		datares["id"] = datauser[0]["id"]
	}
	Respon.Status = 200
	Respon.Data = datares
	Respon.Message = "success"
	return Respon
}
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