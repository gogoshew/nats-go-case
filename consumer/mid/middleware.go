package mid

import (
	"L0_Case/consumer/models"
	"encoding/json"
	"errors"
)

func ValidateMessage(message []byte) (*models.Order, error) {
	msg := new(models.Order)

	err := json.Unmarshal(message, &msg)
	if err != nil {
		return nil, errors.New("incorrect type of message")
	}
	return msg, nil
}
