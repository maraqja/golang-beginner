package account

import (
	"encoding/json"
	"errors"
	"password_app/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

// хранилище аккаунтов
type Vault struct {
	Accounts  []Account `json:"accounts"` // слайс аккаунтов
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func NewVault(db *files.JsonDb) *VaultWithDb {
	file, err := db.Read() // проверяем, существует ли vault (в конструкторе)
	if err != nil {        // просто создаем новый пустой vault
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault) // парсим данные файла в JSON (в переменную vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &VaultWithDb{
		Vault: vault,
		db:    *db,
	}
}

func (vault *VaultWithDb) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)
	vault.UpdatedAt = time.Now()

	vault.save()
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []Account {
	foundAccounts := []Account{}
	for _, account := range vault.Accounts {
		// // проверяем по полному совпадению
		// if account.Url == url {
		if strings.Contains(account.Url, url) { // проверяем по включению подстроки
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
	if err != nil {
		color.Red(err.Error())
	}
	vault.db.Write(data)
}
