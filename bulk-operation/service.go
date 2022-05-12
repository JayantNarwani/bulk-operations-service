package bulkoperation

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"mime/multipart"
	"strconv"
)

type SKUUpdates struct {
	SKUNumber       int  `csv:"sku_number"`
	LeadTime        int  `csv:"lead_time"`
	ShelfLife       int  `csv:"shelf_life"`
	ReorderQuantity int  `csv:"reorder_qty"`
	BufferDay1      int  `csv:"buffer_day_1"`
	BufferDay2      int  `csv:"buffer_day_2"`
	BufferDay3      int  `csv:"buffer_day_3"`
	BufferDayDelay1 int  `csv:"buffer_day_delay_1"`
	BufferDayDelay2 int  `csv:"buffer_day_delay_2"`
	BufferDayDelay3 int  `csv:"buffer_day_delay_3"`
	IsAutoUpdate    bool `csv:"is_auto_update"`
}

func processFile(file multipart.File) {
	// convert into array of structs
	var skuUpdates []SKUUpdates

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
			skuNumber, _ := strconv.Atoi(row[0])
			leadTime, _ := strconv.Atoi(row[1])
			shelfLife, _ := strconv.Atoi(row[2])
			reorderQuantity, _ := strconv.Atoi(row[3])
			bufferDay1, _ := strconv.Atoi(row[4])
			bufferDay2, _ := strconv.Atoi(row[5])
			bufferDay3, _ := strconv.Atoi(row[6])
			bufferDayDelay1, _ := strconv.Atoi(row[7])
			bufferDayDelay2, _ := strconv.Atoi(row[8])
			bufferDayDelay3, _ := strconv.Atoi(row[9])
			isAutoUpdate, _ := strconv.ParseBool(row[10])

			skuUpdate := SKUUpdates{
				SKUNumber:       skuNumber,
				LeadTime:        leadTime,
				ShelfLife:       shelfLife,
				ReorderQuantity: reorderQuantity,
				BufferDay1:      bufferDay1,
				BufferDay2:      bufferDay2,
				BufferDay3:      bufferDay3,
				BufferDayDelay1: bufferDayDelay1,
				BufferDayDelay2: bufferDayDelay2,
				BufferDayDelay3: bufferDayDelay3,
				IsAutoUpdate:    isAutoUpdate,
			}

			skuUpdates = append(skuUpdates, skuUpdate)

			fmt.Printf("%v : %v ", header[i], row[i])
		}
	}
	// Upload to S3

	// Perform validations

	// Send for processing one by one

}
