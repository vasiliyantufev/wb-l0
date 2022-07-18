package app

import (
	"encoding/json"
	"fmt"
	"github.com/vasiliyantufev/wb-l0/internal/models"
)

func ParseMessages(data []byte) (*models.Orders, error) {
	ord := models.Orders{}
	err := json.Unmarshal(data, &ord)
	if err != nil {
		return nil, err
	}
	if ord.Entry != "WBIL" {
		return nil, fmt.Errorf("wrong or missing entry")
	}
	return &ord, nil
}
