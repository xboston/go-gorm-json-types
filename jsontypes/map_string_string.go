package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// MapStringString -  map[string]string
type MapStringString map[string]string

func (s *MapStringString) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

func (s MapStringString) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
