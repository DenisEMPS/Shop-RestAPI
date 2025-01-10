package types

type ClientDAO struct {
	Name             string `json:"name" db:"name"`
	Surname          string `json:"surname" db:"surname"`
	Birthday         string `json:"birthday" db:"birthday"`
	Gender           bool   `json:"gender" db:"gender"`
	RegistrationDate string `json:"registration_date" db:"registration_date"`
	Country          string `json:"country" db:"country"`
	City             string `json:"city" db:"city"`
	Street           string `json:"street" db:"street"`
}

type CreateClient struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Gender   bool   `json:"gender" binding:"required"`
	Country  string `json:"country" binding:"required"`
	City     string `json:"city" binding:"required"`
	Street   string `json:"street" binding:"required"`
}
