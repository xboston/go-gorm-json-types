package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// ArrayStrings - массив строк
type ArrayStrings []string

// Scan - формирование массива
func (s *ArrayStrings) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

// Value - формирование байтового массива
func (s ArrayStrings) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
