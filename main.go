package main

import (
	controller "goec/Controller"
	Dbconnet "goec/Db.connet"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	Dbconnet.Initialize()
	controller.SetupItemController(e)

	e.Logger.Fatal(e.Start(":8000"))
}
