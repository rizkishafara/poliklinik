package controllers

import (
	"log"
	"net/http"
	"poliklinik/models"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func PageDokter(c echo.Context) error {
	data := pongo2.Context{
		"title": "Login",
	}
	return c.Render(http.StatusOK, "views/mst_data/dokter.html", data)
}
func GetDokter(c echo.Context) error {
	resp := models.GetDokter()
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func CreateDokter(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_hp := c.FormValue("no_hp")

	resp := models.CreateDokter(nama, alamat, no_hp)
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK,resp)
}
func UpdateDokter(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_hp := c.FormValue("no_hp")

	resp := models.UpdateDokter(id, nama, alamat, no_hp)
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func DeleteDokter(c echo.Context) error {
	id := c.FormValue("id")

	resp := models.DeleteDokter(id)
	if resp.Status != 200 {
		log.Println(resp.Message)
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
