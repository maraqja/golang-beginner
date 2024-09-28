package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()   // откладываем выполнение этой функции в самый конец функции (в самый конец stack frame этой функции)
	defer fmt.Println(1) // будет позже
	defer fmt.Println(2) // будет раньше
	// тот кого 1-го вызвали с defer - тот будет позже выполнен

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(("Запись успешна"))
}
