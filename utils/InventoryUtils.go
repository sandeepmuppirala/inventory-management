package utils

import (
	"encoding/json"
	"inventory-management/models"
	"net/http"
)

func BuildJsonResponse(w http.ResponseWriter, inventory interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(inventory); err != nil {
		panic(err)
	}
}

func BuildErrorResponse(w http.ResponseWriter, errorCode string, errorMessage string, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	err := models.Error{}
	err.ErrorCode = errorCode
	err.ErrorMessage = errorMessage
	if jsonEncodingErr := json.NewEncoder(w).Encode(err); jsonEncodingErr != nil {
		panic(jsonEncodingErr)
	}
}

func IsValidCategory(category string) bool {
	switch category {
	case
		"Shoes",
		"Exercise",
		"Books",
		"Bags",
		"Sunglasses":
		return true
	}
	return false
}
