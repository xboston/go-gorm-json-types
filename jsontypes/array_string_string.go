package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// ArrayArrayStrings - [][]string
type ArrayArrayStrings [][]string

func (s *ArrayArrayStrings) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

func (s ArrayArrayStrings) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
