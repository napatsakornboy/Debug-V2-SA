package entity

import (
	"time"

	"gorm.io/gorm"
)

type DOCTOR struct {
	gorm.Model
	Name     string
	Title    string
	Password string
	BASKETS  []BASKET `gorm:"foreignKey:DOCTOR_ID"`
}

type BASKET struct {
	gorm.Model
	Amount   int
	Add_time time.Time

	DOCTOR_ID *uint
	DOCTOR    DOCTOR `gorm:"references:id"`

	WHERE_ID *uint
	WHERE    WHERE `gorm:"references:id"`

	MEDICINE_ID *uint
	MEDICINE    MEDICINE `gorm:"references:id"`

	Symptom_ID *uint
	Symptom    Symptom `gorm:"references:id"`

	//รอแอดของพืชที่เป็น FK heck_ID
}

type WHERE struct {
	gorm.Model
	Name    string   `gorm:"uniqueIndex"`
	BASKETS []BASKET `gorm:"foreignKey:WHERE_ID"`
}

type MEDICINE struct {
	gorm.Model
	NameTH  string
	Name    string
	How     string
	So      string
	Unit    string
	BASKETS []BASKET `gorm:"foreignKey:MEDICINE_ID"`
}
