package main

import (
	"fmt"
	"image-steganography/image"
	"log"
)

func main() {
	// Encode Image
	binaryData, err := image.ImageToBinary("input_image.png")
	if err != nil {
		log.Fatal(err)
	}

	obfuscatedText := image.BinaryToObfuscatedText(binaryData)
	err = image.SaveEncodedText("encoded_image.txt", obfuscatedText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image encoded and saved to encoded_image.txt")

	// Decode Image
	obfuscatedText, err = image.LoadEncodedText("encoded_image.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = image.ObfuscatedTextToImage(obfuscatedText, "reconstructed_image.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image decoded and saved to reconstructed_image.png")
}
