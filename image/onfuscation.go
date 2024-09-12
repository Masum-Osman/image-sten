package image

func BinaryToObfuscatedText(binaryData []byte) string {
	var obfuscatedText []byte
	for _, b := range binaryData {
		obfuscatedText = append(obfuscatedText, b+10) // Simple obfuscation
	}
	return string(obfuscatedText)
}

func ObfuscatedTextToBinary(obfuscatedText string) ([]byte, error) {
	var binaryData []byte
	for i := 0; i < len(obfuscatedText); i++ {
		binaryData = append(binaryData, obfuscatedText[i]-10) // Reverse obfuscation
	}
	return binaryData, nil
}
