package repo

import (
	"L0_Case/consumer/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase interface {
	InsertRow(order *models.Order)
	GetRowById(id uint) (*models.Order, error)
	GetAllRows() (*[]models.Order, error)
}

type Database struct {
	db *gorm.DB
}

func GormConnect(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("err from GormOpen: %s", err)
	}

	err = db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Items{})
	if err != nil {
		return nil, fmt.Errorf("err from Automigrate: %s", err)
	}

	return &Database{db: db}, nil
}

func (db *Database) InsertRow(order *models.Order) (uint, error) {
	err := db.db.Debug().Create(&order).Error
	if err != nil {
		return 0, errors.New("error from InsertRow")
	}
	return order.ID, nil
}

func (db *Database) GetRowById(id uint) (*models.Order, error) {
	order := models.Order{ID: 0}

	err := db.db.Debug().Model(&models.Order{}).Where("id = ?", id).Take(&order).Error
	if err != nil {
		return nil, errors.New("can't get order by id")
	}

	if order.ID == 0 {
		return nil, errors.New("something went wrong, can't get order by id")
	}

	return &order, nil
}

func (db *Database) GetAllRows() (*[]models.Order, error) {
	var orders = new([]models.Order)

	err := db.db.Debug().Preload("Delivery").Preload("Payment").Preload("Items").Model(&models.Order{}).Limit(7).Find(&orders).Error

	if err != nil {
		return nil, errors.New("can't get all rows from db")
	}
	return orders, nil
}
