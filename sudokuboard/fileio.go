package sudokuboard

import (
	"encoding/json"
	"io/ioutil"
)

func LoadFromJsonFile(path string) (*Board, error) {
	file, err := ioutil.ReadFile(path)
	if err == nil {
		x := Board{}
		err = json.Unmarshal(file, &x)
		return &x, err
	}
	return nil, err
}

func (b *Board) SaveToJsonFile(path string) error {
	file, err := json.MarshalIndent(&b, "", " ")
	if err == nil {
		err = ioutil.WriteFile(path, file, 0644)
	}
	return err
}
