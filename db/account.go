package db

import (
	"errors"
	"github.com/ggdream/mochat-server/model/user"
)


func AddUser(ac *user.Account) error {
	_, err := Get().Exec("insert into user(username, password) values (?,?)", ac.Username, ac.Password)
	if err != nil {
		return err
	}
	return nil
}

func SearchUser(ac *user.Account) error {
	row := Get().QueryRow("select password from user where username=?", ac.Username)

	var pwd string
	if err := row.Scan(&pwd); err != nil {
		return err
	}
	if pwd != ac.Password {
		return errors.New("wrong password")
	}
	return nil
}
