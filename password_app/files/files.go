package files

import (
	"fmt"
	"os"
	"password_app/output"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		output.PrintError(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		output.PrintError(err)
		return
	}
	defer file.Close() // откладываем выполнение этой функции в самый конец функции (в самый конец stack frame этой функции)
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}
	fmt.Println(("Запись успешна"))
}
