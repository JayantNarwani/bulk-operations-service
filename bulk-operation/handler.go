package bulkoperation

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func UploadFile(c echo.Context) error {

	service := c.FormValue("service")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Error("Error while uploading the file. ")
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		log.Error("Error while opening the file")
		return err
	}
	defer file.Close()
	processFile(file)
	return c.String(http.StatusAccepted, fmt.Sprintf("Request accpeted for service %v, will be processed soon.", service))
}
