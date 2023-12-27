package mapper

import (
	"github.com/murilogilfelpeto/goJob/handler/dto/request"
	"github.com/murilogilfelpeto/goJob/handler/dto/response"
	"github.com/murilogilfelpeto/goJob/schemas"
)

func OpeningRequestToOpening(request request.OpeningRequestDto) schemas.Opening {
	return schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
}

func OpeningToOpeningResponse(opening schemas.Opening) response.OpeningResponseDto {
	return response.OpeningResponseDto{
		ID:        opening.ID,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
	}
}
