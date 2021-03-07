package main

import (
	"inventory-management/constants"
	"inventory-management/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main (){
	handler := mux.NewRouter()

	// handler for viewing one, many or all the inventory items  
	handler.HandleFunc(constants.InventoryControllerMapping, service.ViewItemsInInventory).Methods("GET")

	// handler for getting purchase order quantity of an item from the inventory
	handler.HandleFunc(constants.InventoryItemsControllerMapping + constants.IDPathParam , service.GetItemQuantityForOrder).Methods("GET")
	// handler for updating one or more quantity of an item in the inventory
	handler.HandleFunc(constants.InventoryItemsControllerMapping + constants.IDPathParam , service.UpdateQuantityInInventory).Methods("PUT")


	// handler for adding an item to the inventory
	handler.HandleFunc(constants.InventoryItemsControllerMapping, service.AddItemToInventory).Methods("POST")
	// handler for deleting an item from the inventory
	// This should only be used when the item and all the quantities of that item needs to be removed from the inventory
	handler.HandleFunc(constants.InventoryItemsControllerMapping + constants.IDPathParam , service.RemoveItemFromInventory).Methods("DELETE")

	http.ListenAndServe(constants.DeployedPath, handler)
}