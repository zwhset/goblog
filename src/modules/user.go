package modules

import (
	"errors"
	"log"
	"time"
)

type User struct {
	Uid        int
	Username   string
	Password   string
	RePassword string
	Created    string
}

func (this *User) CheckUser() bool {
	db := Mysql()
	// check user
	query := "SELECT username FROM user where username=?"
	err := db.QueryRow(query, this.Username)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (this *User) Login() (bool, error) {
	db := Mysql()

	if (len(this.Username) == 0) || (len(this.Password) == 0) { // 检查帐号或密码是否为空
		return false, errors.New("帐号密码不能为空")
	}

	query := "SELECT password FROM user where username=?" //执行查询判断
	var passwd string                                     // select db passwd
	err := db.QueryRow(query, this.Username).Scan(&passwd)
	if err != nil {
		return false, errors.New("帐号不存在")
	}

	if passwd == this.Password { //检查密码是否正确
		return true, nil
	} else {
		return false, errors.New("密码不正确")
	}
}

func (this *User) Register() (bool, error) {
	db := Mysql()

	if (len(this.Username) == 0) || (len(this.Password) == 0) || (len(this.RePassword) == 0) { // 检查帐号或密码是否为空
		return false, errors.New("帐号密码不能为空")
	}

	if !this.CheckUser() { //check user esxit
		return false, errors.New("用户名已经存在，请更换")

	} else {
		query := "INSERT INTO user(username,password,created) VALUES(?,?,?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return false, err
		}
		// format time
		this.Created = time.Now().Format("2006-01-02 03:04:05")

		_, err = stmt.Exec(this.Username, this.Password, this.Created)
		if err != nil {
			log.Fatal(err)
			return false, err
		} else {
			return true, nil
		}
	}
}
