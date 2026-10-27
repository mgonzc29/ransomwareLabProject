package decript

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

const key = "clave"

// toDo: deriveKey convierte una contraseña de cualquier longitud en 32 bytes (AES-256).
func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func Decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(deriveKey(key))
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()

	auxText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(auxText) == 0 || len(auxText)%bs != 0 {
		return "", errors.New("longitud inválida")
	}

	auxOut := make([]byte, len(auxText))
	for i := 0; i < len(auxText); i += bs {
		block.Decrypt(auxOut[i:i+bs], auxText[i:i+bs])
	}

	sobrante := int(auxOut[len(auxOut)-1])
	if sobrante == 0 || sobrante > bs || sobrante > len(auxOut) {
		return "", errors.New("sobrante inválido")
	}
	auxOut = auxOut[:len(auxOut)-sobrante]

	return string(auxOut), nil
}

func DecryptFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}

	decrypted, err := Decrypt(string(data))
	if err != nil {
		return fmt.Errorf("error al cifrar: %w", err)
	}

	if err := os.WriteFile(path, []byte(decrypted), 0600); err != nil {
		return fmt.Errorf("error al escribir el archivo cifrado: %w", err)
	}

	return nil
}
