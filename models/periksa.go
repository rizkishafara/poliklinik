package models

import (
	"poliklinik/db"
	"poliklinik/utils"
)

func GetPeriksa() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	result, err := dbEngine.QueryString(`SELECT 
											periksa.id AS id,
											dokter.nama AS nm_dokter,
											pasien.nama AS nm_pasien,
											periksa.obat,
											periksa.tgl_periksa,
											periksa.catatan
										FROM "periksa"
										LEFT JOIN dokter ON dokter.id = periksa.id_dokter
										LEFT JOIN pasien ON pasien.id = periksa.id_pasien 
										ORDER BY periksa.id ASC`)

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
func CreatePeriksa(id_dokter, id_pasien, obat, tanggal, catatan string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`INSERT INTO periksa (id_dokter, id_pasien, obat, tgl_periksa, catatan) VALUES (?, ?, ?, ?, ?)`, id_dokter, id_pasien, obat, tanggal, catatan)

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
func UpdatePeriksa(id, id_dokter, id_pasien, obat, tanggal, catatan string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`UPDATE periksa SET id_dokter = ?, id_pasien = ?, obat = ?, tgl_periksa = ?, catatan = ? WHERE id = ?`, id_dokter, id_pasien, obat, tanggal, catatan, id)

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
func DeletePeriksa(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.Exec(`DELETE FROM periksa WHERE id = ?`, id)

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
