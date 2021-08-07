package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	/*e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})*/

	//กำหนด Route ก่อนเลย พร้อมให้ call ไปยัง func ต่างๆ
	h := itemList{}
	h.Initialize()

	e.GET("/items", h.GetAllItem)
	e.POST("/items", h.SaveItem)
	e.GET("items/:id", h.GetItem)
	e.PUT("/items/:id", h.UpdateItem)
	e.DELETE("/items/:id", h.DeleteItem)

	e.Logger.Fatal(e.Start(":8000"))
}

type itemList struct {
	DB *gorm.DB
}

//ให้เชื่อมต่อฐานข้อมูลเมื่อ Initialize
func (h *itemList) Initialize() {
	db, err := gorm.Open("mysql", "root:helloworld@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Item{})

	h.DB = db
}

type Item struct {
	Id    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name" validate:"required"` ////เเท็กเอาไว้ระรับค่า เเละส่งค่าตามที่ระบุ "เป็นkeyของjson"เเละ"เป็นชื่อคอลั่มในdatabase"
	Price int    `json:"price" validate:"required"`
}

func (h *itemList) GetAllItem(c echo.Context) error {
	items := []Item{}

	h.DB.Find(&items)

	return c.JSON(http.StatusOK, items)
}

func (h *itemList) GetItem(c echo.Context) error {
	id := c.Param("id")
	items := Item{}

	if err := h.DB.Find(&items, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, items)
}

///// มีวิธีอื่น///หาvalidation
func (h *itemList) SaveItem(c echo.Context) error {
	items := Item{}

	if err := c.Bind(&items); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// var validate *validator.Validate
	validate := validator.New()
	if err := validate.Struct(items); err != nil {
		fmt.Println(items)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Create(&items).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *itemList) UpdateItem(c echo.Context) error {
	id := c.Param("id")
	itembody := Item{}
	item := Item{}

	if err := c.Bind(&itembody); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.First(&item, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	item.Name = itembody.Name
	item.Price = itembody.Price

	if err := h.DB.Save(&item).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	// if err := h.DB.Model(&item).Updates(Item{Name: itembody.Name, Price: itembody.Price}).Error; err != nil {
	// 	return c.NoContent(http.StatusInternalServerError)
	// }

	return c.JSON(http.StatusOK, item)
}

func (h *itemList) DeleteItem(c echo.Context) error {
	id := c.Param("id")
	// item := Item{}

	// if err := h.DB.Find(&item, id).Error; err != nil {
	// 	return c.String(http.StatusNotFound, "IdNotFound")
	// }

	result := h.DB.Delete(&Item{}, id)
	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "IdNotFound")
	}

	return c.String(http.StatusOK, "Deletesuseccfull")
}
