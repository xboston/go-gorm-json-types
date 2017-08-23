package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// ArrayArrayString - [][]string
type ArrayArrayString [][]string

func (s *ArrayArrayString) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

func (s ArrayArrayString) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
