package encrypter

import "os"

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

func (enc *Encrypter) Encrypt(stringForEncrypt string) string {
	return ""
}

func (enc *Encrypter) Decrypt(encryptedString string) string {
	return ""
}
