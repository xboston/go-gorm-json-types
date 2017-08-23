package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// StringsInterfaceMap -  мапа строк
type StringsInterfaceMap map[string]interface{}

// Scan - формирование массива
func (s *StringsInterfaceMap) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

// Value - формирование байтового массива
func (s StringsInterfaceMap) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
