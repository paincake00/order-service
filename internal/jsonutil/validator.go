package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/paincake00/order-service/internal/delivery/dto"
)

var validate = validator.New()

func ValidateOrderJSON(raw []byte) (*dto.OrderDTO, error) {
	var dtoVar dto.OrderDTO

	decoder := json.NewDecoder(bytes.NewReader(raw))

	if err := decoder.Decode(&dtoVar); err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	if err := validate.Struct(dtoVar); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return &dtoVar, nil
}
