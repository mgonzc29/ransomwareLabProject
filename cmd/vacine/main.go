package main

import (
	"fmt"
	"os"
	"ransomware/comunUtils"
	"ransomware/vacineUtils"
)

const key = "clave"

func main() {
	files, errListFiles := comunUtils.ListFiles(`C:\`, []string{"jpg", "jpeg", "png", "gif", "docx", "xlsx", "pptx", "pdf"})
	if errListFiles != nil {
		panic(errListFiles)
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
