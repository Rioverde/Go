package db

//Student struct is use uppercase naming if you want to use your variable, function in other packages
type Student struct {
	ID         int     `gorm:"INT; PRIMARY KEY; AUTO_INCREMENT"`
	Name       string  `gorm:"TEXT, NOT NULL"`
	Surname    string  `gorm:"TEXT"`
	Department string  `gorm:"TEXT"`
	Gpa        float64 `gorm:"REAL"`
}
