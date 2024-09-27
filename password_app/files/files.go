package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()   // откладываем выполнение этой функции в самый конец функции (в самый конец stack frame этой функции)
	defer fmt.Println(1) // будет позже
	defer fmt.Println(2) // будет раньше
	// тот кого 1-го вызвали с defer - тот будет позже выполнен

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(("Запись успешна"))
}
