package service

import "github.com/murilogilfelpeto/goJob/schemas"

func UpdateOpening(id int, updatedOpening schemas.Opening) (schemas.Opening, error) {
	opening, err := FindById(id)
	if err != nil {
		logger.Errorf("Error while getting opening for id %v: %v", id, err)
		return schemas.Opening{}, err
	}
	logger.Info("Opening found. Updating it...")
	opening.Role = updatedOpening.Role
	opening.Company = updatedOpening.Company
	opening.Link = updatedOpening.Link
	opening.Location = updatedOpening.Location
	opening.Remote = updatedOpening.Remote
	opening.Salary = updatedOpening.Salary

	err = database.Save(&opening).Error
	if err != nil {
		logger.Errorf("Error while updating opening: %v", err)
		return schemas.Opening{}, err
	}

	return opening, nil
}
