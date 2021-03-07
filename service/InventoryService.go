package service

import (
	"encoding/json"
	"fmt"
	"inventory-management/models"
	"inventory-management/persistence"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	// "gopkg.in/go-playground/validator.v9"
)

func ViewItemsInInventory(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		log.Info("View item from inventory with ID: " + id)
		inputID, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		item := persistence.ViewAnItemInInventory(inputID)
		BuildJsonResponse(w, item, http.StatusOK)
	} else {
		log.Info("View all items in inventory")
		inventory := persistence.ViewItemsInInventory()
		BuildJsonResponse(w, inventory, http.StatusOK)
	}
}

func AddItemToInventory(w http.ResponseWriter, r *http.Request) {
	//v := validator.New()
	log.Info("Add item to inventory..")
    item := models.InventoryItem{}
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	itemAdded := persistence.AddItemToInventory(item)
	BuildJsonResponse(w, itemAdded, http.StatusCreated)
}

func GetItemQuantityForOrder(w http.ResponseWriter, r *http.Request) {
	item := models.InventoryItem{}
	err := json.NewDecoder(r.Body).Decode(&item)
	quantityQueryParam := r.FormValue("quantity")

	if item.ID, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		var errorMessage = "Item ID in the URL should be a number"
		log.Info("Item ID in the URL should be a number")
		BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
		return
	}

	if quantityQueryParam == "" {
		item.Quantity = 1
	} else if item.Quantity, err = strconv.Atoi(quantityQueryParam); err != nil {
		var errorMessage = "Item Quantity in the URL should be a number"
		log.Info("Item Quantity in the URL should be a number. Value received: " , quantityQueryParam)
		BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
		return
	}

	log.Info("Get Item Quantity For Order with ID: " + strconv.Itoa(item.ID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if item.ID > 0 && item.Quantity > 0 {
		itemDetails := persistence.GetItemQuantityForOrder(item.ID, item.Quantity)
		BuildJsonResponse(w, itemDetails, http.StatusOK)
	} else {
		errorMessage := ""
		if item.ID < 0 && item.Quantity < 0 {
			errorMessage = "Invalid input data. Item ID and Quantity cannot be less than one"
		} else if item.ID < 0 {
			errorMessage = "Invalid input data. Item ID cannot be less than one"
		} else {
			errorMessage = "Invalid input data. Item Quantity cannot be less than one"
		}
		BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
	}
}

func UpdateQuantityInInventory(w http.ResponseWriter, r *http.Request) {
	item := models.InventoryItem{}
    err := json.NewDecoder(r.Body).Decode(&item)
	if item.ID, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		var errorMessage = "Item ID in the URL should be a number"
		log.Info("Item ID in the URL should be a number")
		BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
		return
	}
	
	log.Info("Update quantity of the item in the inventory with ID: " + strconv.Itoa(item.ID))
    
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if item.ID > 0 && item.Quantity > 0 {
		itemUpdated := persistence.UpdateQuantityInInventory(item.ID, item.Quantity)
		BuildJsonResponse(w, itemUpdated, http.StatusCreated)
	} else {
		errorMessage := ""
		if item.ID < 0 && item.Quantity < 0 {
			errorMessage = "Invalid input data. Item ID and Quantity cannot be less than one"
		} else if item.ID < 0 {
			errorMessage = "Invalid input data. Item ID cannot be less than one"
		} else { 
			errorMessage = "Invalid input data. Item Quantity cannot be less than one"
		}
		BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
	}
}

func RemoveItemFromInventory(w http.ResponseWriter, r *http.Request) {
	// if resource unavailable return 404, but if resource was found and deleted return 204 no content
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Remove item from inventry..")
}

func BuildJsonResponse(w http.ResponseWriter, inventory interface{}, httpStatus int){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(inventory); err != nil {
		panic(err)
	}
}

func BuildErrorResponse(w http.ResponseWriter, errorCode string, errorMessage string, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
	err := models.Error {}
	err.ErrorCode = errorCode
	err.ErrorMessage = errorMessage
	if jsonEncodingErr := json.NewEncoder(w).Encode(err); jsonEncodingErr != nil {
		panic(jsonEncodingErr)
	}
}