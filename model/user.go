package models

import "errors"

type User struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	IsActive bool `json:"is_active"`
}

var users =  []User{
	{
		Username: "user1",
		Fullname: "user one",
		Email: "sinclairolajuwon@gmail.com",
		IsActive:false,
	},

	{
		Username:"user2",
		Fullname: "user two",
		Email: "olajuwonsinclair@gmail.com",
		IsActive: true,
	},
}

func ListUsers()([]User){
	return users
}

func GetUser(username string)(*User, error){
	for _,user := range users{
		if user.Username == username{
			return &user,nil
		}
	}
	return nil,errors.New("not found")
}

func CreateUsers(user User) error{
	users = append(users, user)
	return nil
}

type Country struct{
	Name string
}

var Countries = []Country{
	{
		Name: "Nigeria",
	},
	{
		Name: "America",
	},
	{
		Name: "Ukraine",
	},
	{
		Name: "Germany",
	},
	{
		Name: "London",
	},
	{
		Name: "Norway",
	},
	{
		Name: "Sweedeen",
	},
}

func ListCountries()([]Country,error){
	return Countries,nil
}


