package controllers

import (
	"net/http"
	"poliklinik/models"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func PagePeriksa(c echo.Context) error {
	data := pongo2.Context{
		"title": "Periksa",
	}
	return c.Render(http.StatusOK, "views/periksa/periksa.html", data)
}

func GetPeriksa(c echo.Context) error {
	resp := models.GetPeriksa()
	if resp.Status != 200 {
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func CreatePeriksa(c echo.Context) error {
	id_pasien := c.FormValue("pasien")
	id_dokter := c.FormValue("dokter")
	obat := c.FormValue("obat")
	tgl_periksa := c.FormValue("tgl_periksa")
	catatan := c.FormValue("catatan")

	resp := models.CreatePeriksa(id_dokter, id_pasien, obat, tgl_periksa, catatan)
	if resp.Status != 200 {
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func UpdatePeriksa(c echo.Context) error {
	id := c.FormValue("id")
	id_pasien := c.FormValue("pasien")
	id_dokter := c.FormValue("dokter")
	obat := c.FormValue("obat")
	tgl_periksa := c.FormValue("tgl_periksa")
	catatan := c.FormValue("catatan")

	resp := models.UpdatePeriksa(id, id_dokter, id_pasien, obat, tgl_periksa, catatan)
	if resp.Status != 200 {
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func DeletePeriksa(c echo.Context) error {
	id := c.FormValue("id")

	resp := models.DeletePeriksa(id)
	if resp.Status != 200 {
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
