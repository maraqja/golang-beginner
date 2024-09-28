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

func NewVault() *Vault {
	file, err := files.ReadFile("data.json") // проверяем, существует ли vault (в конструкторе)
	if err != nil {                          // просто создаем новый пустой vault
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault) // парсим данные файла в JSON (в переменную vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *Vault) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
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

func (vault *Vault) ToBytes() ([]byte, error) { // метод для преобразования аккаунта-структуры в итоговый байт массив для создания файла
	file, err := json.Marshal(vault) // принимает переменную любого типа, которую нужно сериализовать в JSON, и возвращает два значения: сериализованные данные в виде байтового массива ([]byte) и ошибку (error)
	if err != nil {
		return nil, errors.New("CANT_PARSE_TO_JSON")
	}
	return file, nil
}
