package controller

import (
	"fmt"
	Dbconnet "goec/Db.connet"
	"goec/model"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func SetupItemController(e *echo.Echo) {
	e.GET("/items", GetAllItem)
	e.POST("/items", SaveItem)
	e.GET("items/:id", GetItem)
	e.PUT("/items/:id", UpdateItem)
	e.DELETE("/items/:id", DeleteItem)
}

func GetAllItem(c echo.Context) error {
	items := []model.Item{}

	// h.DB.Find(&items)
	Dbconnet.GetDatabase().Find(&items)

	return c.JSON(http.StatusOK, items)
}

func GetItem(c echo.Context) error {
	id := c.Param("id")
	items := model.Item{}

	if err := Dbconnet.GetDatabase().Find(&items, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, items)
}

///// มีวิธีอื่น///หาvalidation
func SaveItem(c echo.Context) error {
	items := model.Item{}

	if err := c.Bind(&items); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// var validate *validator.Validate
	validate := validator.New()
	if err := validate.Struct(items); err != nil {
		fmt.Println(items)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := Dbconnet.GetDatabase().Create(&items).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func UpdateItem(c echo.Context) error {
	id := c.Param("id")
	itembody := model.Item{}
	item := model.Item{}

	if err := c.Bind(&itembody); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(itembody); err != nil {
		fmt.Println(itembody)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := Dbconnet.GetDatabase().First(&item, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	item.Name = itembody.Name
	item.Price = itembody.Price

	if err := Dbconnet.GetDatabase().Save(&item).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	// if err := h.DB.Model(&item).Updates(Item{Name: itembody.Name, Price: itembody.Price}).Error; err != nil {
	// 	return c.NoContent(http.StatusInternalServerError)
	// }

	return c.JSON(http.StatusOK, item)
}

func DeleteItem(c echo.Context) error {
	id := c.Param("id")
	// item := Item{}

	// if err := h.DB.Find(&item, id).Error; err != nil {
	// 	return c.String(http.StatusNotFound, "IdNotFound")
	// }

	result := Dbconnet.GetDatabase().Delete(&model.Item{}, id)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "IdNotFound")
	}

	return c.String(http.StatusOK, "Deletesuseccfull")
}
