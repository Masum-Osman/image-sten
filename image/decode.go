package image

import (
	_ "image/png"
	"io/ioutil"
)

func ObfuscatedTextToImage(obfuscatedText string, outputPath string) error {
	binaryData, err := ObfuscatedTextToBinary(obfuscatedText)
	if err != nil {
		return err
	}

	return BinaryToImage(binaryData, outputPath)
}

func SaveEncodedText(filePath string, text string) error {
	return ioutil.WriteFile(filePath, []byte(text), 0644)
}

func LoadEncodedText(filePath string) (string, error) {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(text), nil
}
