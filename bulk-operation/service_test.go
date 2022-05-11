package bulkoperation

import (
	"io/ioutil"
	"testing"
)

func Test_processFile(t *testing.T) {
	t.Run("should return success when file can be parsed", func(t *testing.T) {

		file, _ := ioutil.ReadFile("bulk_create_update_sku_ops.csv")

		// convert byte into multi part file

		processFile(file)
	})
}
