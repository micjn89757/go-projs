package model

import (
	"errors"
	"vote-gin/utils/msgcode"

	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int 			`db:"id"`
	Username     string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Role int `db:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(username string) int {
	var user User 
	var err error

	sqlStr := "select id from user where username = ?"
	err = db.Get(user, sqlStr, username)

	if err != nil {
		return msgcode.ERROR_USER_NOT_EXIST
	}

	return msgcode.SUCCESS
}

// TODO GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User 
	var err error 

	sqlStr := "select username, password from user where id = ?"

	err = db.Get(&user, sqlStr, id)

	if err != nil {
		return user, msgcode.ERROR_USER_NOT_EXIST
	}

	return user, msgcode.SUCCESS
}


// TODO GetUsers 查询用户列表  模糊查询，分页
// TODO CreateUser 新增用户
// TODO EditUser 修改用户信息
// TODO ChangePassword 修改用户密码
// TODO DeleteUser 删除用户

// 后台登陆验证
func CheckLogin(name string, password string) (User, int) {
	var user User
	var err error
	var passwordErr error

	sqlStr := "select username, password from user where username = ?"
	err = db.Get(&user, sqlStr, name)

	if err != nil {
		return user, msgcode.ERROR_USER_NOT_EXIST 
	}

	// passwordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))  

	// TODO 实验性
	if password != user.Password {
		passwordErr = errors.New("password err")
	}

	if passwordErr != nil {
		return user, msgcode.ERROR_PASSWORD_WRONG
	}

	return user, msgcode.SUCCESS
}