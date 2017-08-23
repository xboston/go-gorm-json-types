package jsontypes

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// MapStringsInterface - map[string]interface{}
type MapStringsInterface map[string]interface{}

func (s *MapStringsInterface) Scan(value interface{}) error {

	asBytes, ok := value.([]byte)

	if !ok {
		return errors.New("Scan source is not []byte")
	}

	return jsoniter.Unmarshal(asBytes, s)
}

func (s MapStringsInterface) Value() (driver.Value, error) {
	return jsoniter.Marshal(&s)
}
