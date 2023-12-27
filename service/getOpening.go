package service

import "github.com/murilogilfelpeto/goJob/schemas"

func FindAll() ([]schemas.Opening, error) {
	var openings []schemas.Opening
	err := database.Find(&openings).Error
	if err != nil {
		return nil, err
	}
	return openings, nil
}

func FindById(id int) (schemas.Opening, error) {
	var opening schemas.Opening
	err := database.First(&opening, id).Error
	if err != nil {
		return schemas.Opening{}, err
	}
	return opening, nil
}
