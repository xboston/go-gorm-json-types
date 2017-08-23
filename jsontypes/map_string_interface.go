package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// MapStringInterface - map[string]interface{}
type MapStringInterface map[string]interface{}

func (s *MapStringInterface) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)

	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

func (s MapStringInterface) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
