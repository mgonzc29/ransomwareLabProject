package main

import (
	"fmt"
	"os"
	"ransomware/decript"
	"ransomware/utils"
)

func main() {
	files, errListFiles := utils.ListFiles(`C:\`, []string{"jpg", "png"})
	if errListFiles != nil {
		panic(errListFiles)
	}
	fmt.Println("Descifrando los siguientes archivos:")
	for _, file := range files {
		errDecrypt := decript.DecryptFile(file)
		if errDecrypt != nil {
			fmt.Fprintf(os.Stderr, "Error desencriptando %s: %v\n", file, errDecrypt)
		}
		fmt.Println(file)
	}
}
