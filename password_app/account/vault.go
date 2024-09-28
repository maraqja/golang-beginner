package account

import (
	"encoding/json"
	"errors"
	"time"
)

// хранилище аккаунтов
type Vault struct {
	Accounts  []Account `json:"accounts"` // слайс аккаунтов
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) { // метод для преобразования аккаунта-структуры в итоговый байт массив для создания файла
	file, err := json.Marshal(vault) // принимает переменную любого типа, которую нужно сериализовать в JSON, и возвращает два значения: сериализованные данные в виде байтового массива ([]byte) и ошибку (error)
	if err != nil {
		return nil, errors.New("CANT_PARSE_TO_JSON")
	}
	return file, nil
}
