package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) { // хотим распечатать любой тип вообще (any)
	value, ok := value.(int)
	if ok {
		color.Red("2 СПОСОБ: Код ошибки: %d", value)
	}
	switch t := value.(type) { // получаем тип значения (можно использовать только внутри swtich !!!)
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", value)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвстный тип ошибки")
	}

}
