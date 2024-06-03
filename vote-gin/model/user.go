package model

type User struct {
	ID int	`db:"id"`
	Name         string	`db:"name"`
	Password     string	`db:"password"`
}

// QueryUser 查询单个用户
func QueryUserInfo(username string, password string) (code int) {
	var err error 
	var usr User 

	sqlStr := "select name, password from users where name = ? and password = ?"
	err = db.Get(&usr, sqlStr, username, password)
	if err != nil {
		sugar.Errorf("not find user:%w\n", err)
		return 1000 // 表示未查找到
	}

	return 1001  // 表示查找到
}