/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package model

type Login struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type JwtLoginUser struct {
	Pk int `json:"pk,omitempty"`
	User string `json:"user,omitempty"`
	Email string `json:"email,omitempty"`
}
