package types

type Client struct {
	Id               int    `json:"client_id" db:"client_id"`
	Name             string `json:"name" binding:"required" db:"name"`
	Surname          string `json:"surname" binding:"required" db:"surname"`
	Birthday         string `json:"birthday" binding:"required" db:"birthday"`
	Gender           bool   `json:"gender" binding:"required" db:"gender"`
	RegistrationDate string `json:"registration_date" binding:"required" db:"registration_date"`
	Country          string `json:"country" binding:"required" db:"country"`
	City             string `json:"city" binding:"required" db:"city"`
	Street           string `json:"street" binding:"required" db:"street"`
}

type ClientDTO struct {
	Id               int    `json:"client_id" db:"client_id"`
	Name             string `json:"name" binding:"required" db:"name"`
	Surname          string `json:"surname" binding:"required" db:"surname"`
	Birthday         string `json:"birthday" binding:"required" db:"birthday"`
	Gender           bool   `json:"gender" binding:"required" db:"gender"`
	RegistrationDate string `json:"registration_date" binding:"required" db:"registration_date"`
	AdressID         int    `json:"adress_id" db:"adress_id"`
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
