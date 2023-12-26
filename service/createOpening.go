package service

import (
	"github.com/murilogilfelpeto/goJob/schemas"
)

func Save(opening schemas.Opening) (schemas.Opening, error) {
	logger.Infof("Saving opening: %v", opening)
	transaction := database.Begin()
	err := transaction.Create(&opening).Error
	if err != nil {
		logger.Errorf("Error while saving opening: %v", err)
		transaction.Rollback()
		return schemas.Opening{}, err
	}

	transaction.Commit()
	return opening, nil

}
