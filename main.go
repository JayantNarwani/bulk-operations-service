package main

import (
	bulkoperation "bulk-operations-service/bulk-operation"
	"bulk-operations-service/ping"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	initHandlers(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func initHandlers(e *echo.Echo) {
	e.GET("/ping", ping.Ping)
	e.POST("/uploadFile", bulkoperation.UploadFile)

}
