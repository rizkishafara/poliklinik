package models

import (
	"poliklinik/db"
	"poliklinik/utils"
)

func GetPasien() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	result, err := dbEngine.QueryString(`SELECT * FROM pasien`)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	Respon.Data = result
	return Respon
}
func CreatePasien(nama, alamat, no_hp string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`INSERT INTO pasien (nama, alamat, no_hp) VALUES (?, ?, ?)`, nama, alamat, no_hp)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func UpdatePasien(id, nama, alamat, no_hp string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`UPDATE pasien SET nama = ?, alamat = ?, no_hp = ? WHERE id = ?`, nama, alamat, no_hp, id)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func DeletePasien(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`DELETE FROM pasien WHERE id = ?`, id)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
