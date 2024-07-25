package image

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

//func ReadImageToBase64(imagePath string) (string, error) {
//	img, err := getImageFromFilePath(imagePath)
//	if err != nil {
//		return "", fmt.Errorf("failed to read image from file:[%s], error: %v ", imagePath, err.Error())
//	}
//
//	buf := new(bytes.Buffer)
//	if encodeErr := jpeg.Encode(buf, img, nil); encodeErr != nil {
//		return "", fmt.Errorf("failed to encode image to buffer, error: %v", encodeErr)
//	}
//
//	imgBase64Str := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
//
//	return imgBase64Str, nil
//}
//
//func getImageFromFilePath(filePath string) (image.Image, error) {
//	f, err := os.Open(filePath)
//	if err != nil {
//		return nil, err
//	}
//	defer f.Close()
//	image, _, err := image.Decode(f)
//	return image, err
//}

func ReadImageToBase64V2(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read image from file:[%s], error: %v ", path, err.Error())
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding, nil
}
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
