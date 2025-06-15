package error

import (
	"encoding/json"
	"errors"
)

func EncodeError(message string) error {
	encode, _ := json.Marshal(map[string]string{
		"Message": message,
	})

	return errors.New(string(encode))
}
