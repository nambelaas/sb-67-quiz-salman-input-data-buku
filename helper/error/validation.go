package error

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidationCheck(err error) any {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var details []map[string]string
		for _, fe := range ve {
			details = append(details, map[string]string{
				"field": fe.Field(),
				"tag":   fe.Tag(),
				"value": fe.Param(),
			})
		}

		errorData := map[string]any{
			"Message": "gagal validasi data",
			"Detail":  details,
		}

		return errorData
	}

	return nil
}
