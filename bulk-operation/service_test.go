package bulkoperation

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"testing"
)

func Test_processFile(t *testing.T) {
	t.Run("should return success when file can be parsed", func(t *testing.T) {

		body := new(bytes.Buffer)
		mw := multipart.NewWriter(body)

		filePath := "bulk_create_update_sku_ops.csv"
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatal(err)
		}
		w, err := mw.CreateFormFile("file", filePath)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}

		processFile(file)
	})
}
