package models

type User struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"varchar(255)" json:"username"`
	Password string `gorm:"varchar(255)" json:"password"`
}
