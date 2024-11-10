package models

import (
	"poliklinik/db"
	"poliklinik/utils"
)

func GetDokter() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	result, err := dbEngine.QueryString(`SELECT * FROM dokter`)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	Respon.Data = result
	return Respon
}
func CreateDokter(nama, alamat, no_hp string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`INSERT INTO dokter (nama, alamat, no_hp) VALUES (?, ?, ?)`, nama, alamat, no_hp)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func UpdateDokter(id, nama, alamat, no_hp string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`UPDATE dokter SET nama = ?, alamat = ?, no_hp = ? WHERE id = ?`, nama, alamat, no_hp, id)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func DeleteDokter(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`DELETE FROM dokter WHERE id = ?`, id)

	if err != nil {
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
