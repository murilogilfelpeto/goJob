package service

import "github.com/murilogilfelpeto/goJob/schemas"

func DeleteOpening(id int) error {
	var opening schemas.Opening
	err := database.First(&opening, id).Error
	if err != nil {
		return err
	}
	err = database.Delete(&opening).Error
	if err != nil {
		return err
	}
	return nil
}
