package modules

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	Uid      int
	Username string
	Password string
	Created  string
}

func (this *User) Login() (bool, error) {
	db := Mysql()

	// 检查帐号或密码是否为空
	if (len(this.Username) == 0) || (len(this.Password) == 0) {
		return false, errors.New("帐号密码不能为空")
	}

	//执行查询判断
	query := "SELECT password FROM user where username=?"
	var passwd string // select db passwd
	err := db.QueryRow(query, this.Username).Scan(&passwd)

	if err != nil {
		fmt.Println(this.Username, this.Password)
		return false, errors.New("帐号不存在")
	}

	if passwd == this.Password {
		return true, nil
	} else {
		return false, errors.New("密码不正确")
	}
}

func (this *User) CheckUser() bool {
	db := Mysql()
	// check user
	query := "SELECT username FROM user where username=?"
	err := db.QueryRow(query, this.Username)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (this *User) Register() bool {
	db := Mysql()
	//check user esxit
	if this.CheckUser() {
		return false //username esxit
	} else {
		query := "INSERT INTO user(username,password,created) VALUES(?,?,?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return false
		}
		_, err = stmt.Exec(this.Username, this.Password, this.Created)
		if err != nil {
			log.Fatal(err)
			return false
		} else {
			return true
		}
	}
}
