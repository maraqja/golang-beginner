package cloud

import "fmt"

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {

	return nil, nil
}

func (db *CloudDb) Write(content []byte) {
	fmt.Println(("Запись успешна"))
}
