package request

type OpeningRequestDto struct { //nolint
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}
