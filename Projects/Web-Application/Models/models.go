package models

//Users type Export different type of user from database
type Users struct {
	ID       int    `gorm:"type: INT; PRIMARY_KEY; AUTO_INCREMENT"`
	Email    string `gorm:"type: TEXT: NOT NULL"`
	Password string `gorm:"type: TEXT; NOT NULL"`
}
