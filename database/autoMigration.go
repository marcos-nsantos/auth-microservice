package database

import "github.com/marcos-nsantos/e-commerce/auth-service/model"

func AutoMigrateUser() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(model.User{}); err != nil {
		return err
	}

	return nil
}
