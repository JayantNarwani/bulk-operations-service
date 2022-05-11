package bulkoperation

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

func processFile(file multipart.File) {
	// Read file
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		log.Fatal(err)
	}
	header := rows[0]
	for _, row := range rows[1:] {
		for i := 0; i < len(row); i++ {
			fmt.Printf("%v : %v ", header[i], row[i])
		}

		fmt.Println("")
	}
	// Upload to S3

	// Perform validations

	// Send for processing one by one

}
