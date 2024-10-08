package account

import (
	"encoding/json"
	"errors"
	"password_app/encrypter"
	"password_app/output"
	"strings"
	"time"

	"github.com/fatih/color"
)

// хранилище аккаунтов
type Vault struct {
	Accounts  []Account `json:"accounts"` // слайс аккаунтов
	UpdatedAt time.Time `json:"updatedAt"`
}

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type VaultWithDb struct {
	Vault
	db  Db // нельзя использовать указатель на интерфейс
	enc *encrypter.Encrypter
}

// Интерфейсы нельзя использовать с указателем
// Корректной реализаицией интерфейса являются все struct, реализующие методы интерфейса
// Не нужно писать никаких implements Db итд
func NewVault(db Db, enc *encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read() // проверяем, существует ли vault (в конструкторе)
	if err != nil {        // просто создаем новый пустой vault
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	decryptedData := enc.Decrypt(file) // при загрузке из файла сохраняем в память расшифрованные данные
	var vault Vault
	err = json.Unmarshal(decryptedData, &vault) // парсим данные файла в JSON (в переменную vault)
	if err != nil {
		output.PrintError(err)
	}
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)
	vault.UpdatedAt = time.Now()

	vault.save()
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	foundAccounts := []Account{}
	for _, account := range vault.Accounts {

		if checker(account, str) { // проверяем по включению подстроки
			foundAccounts = append(foundAccounts, account)
		}
	}
	return foundAccounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	newAccounts := []Account{}
	isDeleted := false
	for _, account := range vault.Accounts {
		if !strings.Contains(account.Url, url) {
			newAccounts = append(newAccounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = newAccounts

	vault.save()
	return isDeleted
}

func (vault *Vault) ToBytes() ([]byte, error) { // метод для преобразования аккаунта-структуры в итоговый байт массив для создания файла
	file, err := json.Marshal(vault) // принимает переменную любого типа, которую нужно сериализовать в JSON, и возвращает два значения: сериализованные данные в виде байтового массива ([]byte) и ошибку (error)
	if err != nil {
		return nil, errors.New("CANT_PARSE_TO_JSON")
	}
	return file, nil
}

func (vault *VaultWithDb) save() { // внутренний метод (поэтому с маленькой буквы, он не экспортируется из-за этого)
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encryptedData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError(err)
	}
	vault.db.Write(encryptedData)
}
