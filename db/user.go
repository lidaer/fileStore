package db

import (
	"fmt"
)

// User : 用户表model
type User struct {
	ID           int64
	UserName     string
	PassWord     string `gorm:"user_pwd"`
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string `gorm:"column:last_active"`
	Status       int
}

func (u User) TableName() string {
	return "tbl_user"
}

type Token struct {
	UserName  string
	UserToken string
}

func (token Token) TableName() string {
	return "tbl_user_token"
}

// UserSignup : 通过用户名及密码完成user表的注册操作
func UserSignup(username string, passwd string) bool {
	user := User{UserName: username, PassWord: passwd}
	err := DB.Create(&user).Error
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	return true
}

// UserSignin : 判断密码是否一致
func UserSignin(username string, encpwd string) bool {

	user := User{UserName: username}
	err := DB.First(&user).Error
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	if user.PassWord == encpwd {
		return true
	} else {
		return false
	}
}

// UpdateToken : 刷新用户登录的token
func UpdateToken(username string, token string) bool {
	err := DB.Model(&Token{}).Where("UserName=?", username).Update("UserToken", token).Error
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	return true
}

// GetUserInfo : 查询用户信息
func GetUserInfo(username string) (*User, error) {
	user := User{}
	user.UserName = username
	err := DB.First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		// error不为nil, 返回时user应当置为nil
		//return user, err
		return nil, err
	}
	return &user, nil
}

// UserExist : 查询用户是否存在
func UserExist(username string) (bool, error) {
	user := User{}
	result := DB.First(&user, "UserName=?", username)
	err := result.Error
	if result.Error != nil {
		fmt.Println(err.Error())
		// error不为nil, 返回时user应当置为nil
		//return user, err
		return false, err
	}
	return true, err
}
