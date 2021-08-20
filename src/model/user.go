package model

/* create table user (
user_id int auto_increment,
user_name varchar(50),
user_password varchar(50),
user_phone varchar(50),
user_birth_day Date,
user_address varchar(50),
PRIMARY KEY (user_id)
	);*/
type User struct {
	UserId int `json:"user_id"`
	UserName string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserPhone string `json:"user_phone"`
	UserBirthDay string `json:"user_birth_day"`
	UserAddress string `json:"user_address"`
}
