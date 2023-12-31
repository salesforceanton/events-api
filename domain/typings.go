package domain

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}

type Event struct {
	Id            int    `json:"id" db:"id"`
	Title         string `json:"title" db:"title" binding:"required"`
	StartDatetime string `json:"startDatetime" db:"startdatetime" binding:"required"`
	TimezoneId    string `json:"timezoneId" db:"timezoneid"`
	OrganizerId   int    `json:"organizerId" db:"organizerid"`
	Description   string `json:"description" db:"description"`
}

type SaveEventRequest struct {
	Title         string `json:"title" db:"title" binding:"required"`
	StartDatetime string `json:"startDatetime" db:"startdatetime" binding:"required"`
	TimezoneId    string `json:"timezoneId" db:"timezoneid"`
	Description   string `json:"description" db:"description"`
}
