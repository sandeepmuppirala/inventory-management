package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"inventory-management/models"
)

func ViewItemsInInventory() []models.Inventory {
	db := getDatabaseConnection()
	inventory := []models.Inventory{}
	defer db.Close()
	db.Order("\"ID\" asc").Find(&inventory)
	fmt.Println("{}", inventory)
	return inventory
}

func ViewAnItemInInventory(id int) models.Inventory {
	db := getDatabaseConnection()
	item := models.Inventory{}
	defer db.Close()
	db.First(&item, id)
	fmt.Println("{}", item)
	return item
}

func AddItemToInventory(inputItem models.InventoryItem) interface{} {
	db := getDatabaseConnection()
	item := models.Inventory{Category: inputItem.Category, Product: inputItem.Product, Quantity: inputItem.Quantity}
	defer db.Close()
	result := db.Create(&item)
	if result.RowsAffected == 1 {
		log.Debug("Item added to the inventory")
		return item
	} else {
		log.Debug("Constraint violated")
		err := models.Error{}
		err.ErrorCode = "IM-105"
		err.ErrorMessage = "Duplicate product!"
		return err
	}
}

func GetItemQuantityForOrder(id int, quantity int) interface{} {
	db := getDatabaseConnection()
	item := models.Inventory{}
	defer db.Close()
	db.First(&item, id)
	if item.ID > 0 {
		if quantity <= item.Quantity {
			// take quantity and update database
			// result := db.First(&item, id)	// this should be an update call
			item.Quantity = item.Quantity - quantity
			result := db.Save(&item)
			if result.RowsAffected == 1 {
				log.Debug("Quantity updated in the inventory")
				item.Quantity = quantity
			}
			return item
		} else {
			// log insufficient quantity in database, check view items and make this call again
			log.Debug("Insufficient quantity in the inventory")
			err := models.Error{}
			err.ErrorCode = "IM-101"
			err.ErrorMessage = "Insufficient quantity in the inventory"
			return err
		}
	}
	return item
}

func UpdateQuantityInInventory(id int, quantity int) models.Inventory {
	db := getDatabaseConnection()
	item := models.Inventory{}
	defer db.Close()
	db.First(&item, id)
	if item.ID > 0 {
		item.Quantity = quantity
		result := db.Save(&item)
		if result.RowsAffected == 1 {
			log.Debug("Quantity updated in the inventory")
		}
	}
	return item
}

func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=ec2-18-204-74-74.compute-1.amazonaws.com	port=5432 user=nfyuczppvznwgy dbname=d1kblkin8bt9sd password=37075b05a82723e8d48a08daa051c85c68a8aa08364ab9bbdde75008172e3462")
	db.SingularTable(true)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
