package controllers

import (
	"log"
	"net/http"
	"poliklinik/models"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func PagePasien(c echo.Context) error {
	data := pongo2.Context{
		"title": "Pasien",
	}
	return c.Render(http.StatusOK, "views/mst_data/pasien.html", data)
}

func GetPasien(c echo.Context) error {
	resp := models.GetPasien()
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func CreatePasien(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_hp := c.FormValue("no_hp")

	resp := models.CreatePasien(nama, alamat, no_hp)
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func UpdatePasien(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_hp := c.FormValue("no_hp")

	resp := models.UpdatePasien(id, nama, alamat, no_hp)
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func DeletePasien(c echo.Context) error {
	id := c.FormValue("id")

	resp := models.DeletePasien(id)
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
