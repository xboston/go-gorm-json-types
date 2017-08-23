package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// StringsMap -  мапа строк
type StringsMap map[string]string

// Scan - формирование массива
func (s *StringsMap) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

// Value - формирование байтового массива
func (s StringsMap) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
