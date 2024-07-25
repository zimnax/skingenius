package image

import (
	"fmt"
	"testing"
)

func Test_ReadImage(t *testing.T) {
	tt := []struct {
		imagePath string
	}{
		{
			imagePath: "Anti Aging Daily Rose Face Moisturizer.jpg",
		},
	}

	for _, s := range tt {
		imgBytes, err := ReadImageToBase64V2(s.imagePath)
		if err != nil {
			t.Fatalf("failed to read image from file:[%s], error: %v ", s.imagePath, err.Error())
		}
		fmt.Println("imgBytes: ", imgBytes)
	}
}
