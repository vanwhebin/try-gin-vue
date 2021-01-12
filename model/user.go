package model

import (
	"gorm.io/gorm"
	//"vanwhebin/try-gin-vue/common"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type varchar(20); not null"`
	Telephone string `gorm:"type varchar(110);not null;unique"`
	Password  string `gorm:"type size:255"`
}


func IsTelephoneExists(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


//func CreateUser(name string, telephone string, password string) *gorm.DB{
//	// 创建用户
//	newUser := User{
//		Name:      name,
//		Telephone: telephone,
//		Password:  password,
//	}
//	result := common.DB.Create(&newUser)
//	fmt.Println(result)
//	return result
//}
