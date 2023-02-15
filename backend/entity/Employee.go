package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~ต้องไม่เป็นค่าว่าง"`
	Email      string
	EmployeeID string `valid:"matches(^[JMS]\\d{8}$)~รหัสพนักงานขึ้นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจำนวน 8 ตัว"`
}
