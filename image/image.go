package image

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"os"
)

func ImageToBinary(imagePath string) ([]byte, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	binaryData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return binaryData, nil
}

func BinaryToImage(binaryData []byte, outputPath string) error {
	img, _, err := image.Decode(bytes.NewReader(binaryData))
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := png.Encode(outFile, img); err != nil {
		return err
	}

	return nil
}
