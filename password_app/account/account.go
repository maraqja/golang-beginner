package account // чтобы создать отдельный пакет, нужно вынести файл в папку с названием как у пакета
// Соглашение - все что написано с большой буквы - ЭКСПОРТИРУЕМОЕ, иначе - нет
import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_*!")

type Account struct { // Описываем тип стракта
	login    string
	password string
	Url      string
}

func newAccount(login, password, UrlString string) (*Account, error) { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
	// Хорошо в случае если нужно валидировать входные данные
	if login == "" { // логин не был передан: дефолтное значение непереданноо строки - пустая строка
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(UrlString)
	if err != nil {
		// return nil, err
		return nil, errors.New("INVALID_Url")
	}
	newAcc := &Account{ // возвращаем именно ссылку на структуру, тк иначе мы создадим структуру в функции и еще потом скопириуем в переменную при вызове функции
		login:    login,
		password: password,
		Url:      UrlString,
	}
	if password == "" { // Если не передали пароль, то генерим
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}

// Указываем что функция outputPassword - метод struct Account
func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.Url)
	acc.Url = "test"
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// В GO нет как такового наследования, но его заменяет композиция
type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account   // Встраивание - появляются embeded fields
}

func NewAccountWithTimeStamp(login, password, UrlString string) (*AccountWithTimeStamp, error) { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
	// Хорошо в случае если нужно валидировать входные данные
	if login == "" { // логин не был передан: дефолтное значение непереданноо строки - пустая строка
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(UrlString)
	if err != nil {
		// return nil, err
		return nil, errors.New("INVALID_Url")
	}
	newAcc := &AccountWithTimeStamp{ // возвращаем именно ссылку на структуру, тк иначе мы создадим структуру в функции и еще потом скопириуем в переменную при вызове функции
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{ // Приходится объявлять так с помощью встраивания
			login:    login,
			password: password,
			Url:      UrlString,
		},
	}
	if password == "" { // Если не передали пароль, то генерим
		// newAcc.Account.generatePassword(10) // ориг запись
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}
