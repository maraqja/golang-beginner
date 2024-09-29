package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	secret := os.Getenv("SECRET") // получаем переменную окружения
	if secret == "" {
		panic("Нет значения для SECRET в переменных окружения")
	}
	return &Encrypter{
		Key: secret,
	}
}

func (enc *Encrypter) Encrypt(stringForEncrypt []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // создаем блок для шифрования
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block) // создаем GCM с помощью блока
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize()) // формируем случайный nonce, который используем для шифрования
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	// В результате теперь в nonce хранится некоторое уникальное значение для шифрования
	return aesGCM.Seal(nonce, nonce, stringForEncrypt, nil) // шифруем тут
}

func (enc *Encrypter) Decrypt(encryptedString []byte) string {
	return ""
}
