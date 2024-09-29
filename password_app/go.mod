module password_app

go 1.23.1

require github.com/fatih/color v1.17.0 // после использования в коде при помощи go mod tidy убираем что эта зависимость indirect (он была indirect (непрямой) тк при первом импорте она еще не была использована)

// наоборот также работает - если перестать использовать зависимость, то после go mod tidy зависимость будет удалена из этого файла

require github.com/joho/godotenv v1.5.1

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
