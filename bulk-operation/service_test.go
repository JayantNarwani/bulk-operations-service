package bulkoperation

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"testing"
)

func Test_processFile(t *testing.T) {
	t.Run("should return success when file can be parsed", func(t *testing.T) {

		file, err := excelize.OpenFile("bulk_create_update_sku_ops.csv")
		if err != nil {
			panic(err)
		}

		processFile(file)
	})
}
