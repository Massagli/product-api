package model

type User struct{
	ID int `json:"id_user"`
	User_name string `json:"user_name"` 
	User_email string `joson:"user_email"`
}