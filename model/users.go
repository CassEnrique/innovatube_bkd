/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package model

type User struct {
	Pk       int    `gorm:"column:pk_user; PrimaryKey;" json:"pk_user,omitempty"`
	Username string `gorm:"column:username;" json:"username,omitempty"`
	Names    string `gorm:"column:names;" json:"names,omitempty"`
	Surnames string `gorm:"column:surnames;" json:"surnames,omitempty"`
	Email    string `gorm:"column:email;" json:"email,omitempty"`
	Password string `gorm:"column:password;" json:"password,omitempty"`
	Token    string `gorm:"-" json:"token,omitempty"`
}

func (User) TableName() string {
	return "tr_users"
}

type Users []User
