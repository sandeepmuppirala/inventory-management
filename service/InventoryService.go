package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"inventory-management/models"
	"inventory-management/persistence"
	"inventory-management/utils"
	"net/http"
	"strconv"
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
		utils.BuildJsonResponse(w, item, http.StatusOK)
	} else {
		log.Info("View all items in inventory")
		inventory := persistence.ViewItemsInInventory()
		utils.BuildJsonResponse(w, inventory, http.StatusOK)
	}
}

func AddItemToInventory(w http.ResponseWriter, r *http.Request) {
	log.Info("Add item to inventory..")
	item := models.InventoryItem{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !utils.IsValidCategory(item.Category) {
		utils.BuildErrorResponse(w, "IM-004", "Invalid item category: "+item.Category, http.StatusBadRequest)
		return
	}
	itemAdded := persistence.AddItemToInventory(item)
	utils.BuildJsonResponse(w, itemAdded, http.StatusCreated)
}

func GetItemQuantityForOrder(w http.ResponseWriter, r *http.Request) {
	item := models.InventoryItem{}
	err := json.NewDecoder(r.Body).Decode(&item)
	quantityQueryParam := r.FormValue("quantity")

	if item.ID, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		var errorMessage = "Item ID in the URL should be a number"
		log.Info("Item ID in the URL should be a number")
		utils.BuildErrorResponse(w, "IM-002", errorMessage, http.StatusBadRequest)
		return
	}

	if quantityQueryParam == "" {
		item.Quantity = 1
	} else if item.Quantity, err = strconv.Atoi(quantityQueryParam); err != nil {
		var errorMessage = "Item Quantity in the URL should be a number"
		log.Info("Item Quantity in the URL should be a number. Value received: ", quantityQueryParam)
		utils.BuildErrorResponse(w, "IM-003", errorMessage, http.StatusBadRequest)
		return
	}

	log.Info("Get Item Quantity For Order with ID: " + strconv.Itoa(item.ID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if item.ID > 0 && item.Quantity > 0 {
		itemDetails := persistence.GetItemQuantityForOrder(item.ID, item.Quantity)
		utils.BuildJsonResponse(w, itemDetails, http.StatusOK)
	} else {
		errorMessage := ""
		if item.ID < 0 && item.Quantity < 0 {
			errorMessage = "Invalid input data. Item ID and Quantity cannot be less than one"
		} else if item.ID < 0 {
			errorMessage = "Invalid input data. Item ID cannot be less than one"
		} else {
			errorMessage = "Invalid input data. Item Quantity cannot be less than one"
		}
		utils.BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
	}
}

func UpdateQuantityInInventory(w http.ResponseWriter, r *http.Request) {
	item := models.InventoryItem{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if item.ID, err = strconv.Atoi(mux.Vars(r)["id"]); err != nil {
		var errorMessage = "Item ID in the URL should be a number"
		log.Info("Item ID in the URL should be a number")
		utils.BuildErrorResponse(w, "IM-002", errorMessage, http.StatusBadRequest)
		return
	}

	log.Info("Update quantity of the item in the inventory with ID: " + strconv.Itoa(item.ID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if item.ID > 0 && item.Quantity > 0 {
		itemUpdated := persistence.UpdateQuantityInInventory(item.ID, item.Quantity)
		if itemUpdated.ID != 0 {
			utils.BuildJsonResponse(w, itemUpdated, http.StatusOK)
		} else {
			utils.BuildErrorResponse(w, "IM-004", "Failed to update quantity for ID: "+strconv.Itoa(item.ID), http.StatusConflict)
		}
	} else {
		errorMessage := ""
		if item.ID < 0 && item.Quantity < -1 {
			errorMessage = "Invalid input data. Item ID cannot be less than one and Quantity cannot be less than zero"
		} else if item.ID < 0 {
			errorMessage = "Invalid input data. Item ID cannot be less than one"
		} else {
			errorMessage = "Invalid input data. Item Quantity cannot be less than zero"
		}
		utils.BuildErrorResponse(w, "IM-001", errorMessage, http.StatusBadRequest)
	}
}

func RemoveItemFromInventory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Remove item from inventory..")
}
