package models

//User type Export different type of user from database
type User struct {
	ID       int    `gorm:"type: INT; PRIMARY_KEY; AUTO_INCREMENT"`
	Name     string `gorm:"type: TEXT: NOT NULL"`
	Password string `gorm:"type: TEXT; NOT NULL"`
}

//Data struct is data comming from database mydata.db
type Data struct {
	ID       int    `gorm:"type: INT; PRIMARY_KEY; AUTO_INCREMENT"`
	Name     string `gorm:"type: TEXT: NOT NULL"`
	Email    string `gorm:"type: TEXT; NOT NULL"`
	Password string `gorm:"type: TEXT; NOT NULL"`
}
