package account // чтобы создать отдельный пакет, нужно вынести файл в папку с названием как у пакета
// Соглашение - все что написано с большой буквы - ЭКСПОРТИРУЕМОЕ, иначе - нет
import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_*!")

type Account struct { // Описываем тип стракта
	Login string `json:"login"` // это тег структуры - нужен для преобразования в JSON и другие форматы (пробел разделяет теги)
	// Благодаря этим тегам можем в runtime получать доп метаинформацию (с помощью встроенной библиотеки Reflect)
	Password  string    `json:"password"` // все поля с тегами должны экспортироваться
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, UrlString string) (*Account, error) { // Это типо можно использовать как конструктор стракта (название просто по соглашению)
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
		Login:     login,
		Password:  password,
		Url:       UrlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" { // Если не передали пароль, то генерим
		// newAcc.Account.generatePassword(10) // ориг запись
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}

// Указываем что функция outputPassword - метод struct Account
func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}
