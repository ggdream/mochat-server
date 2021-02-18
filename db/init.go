package db

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql" // mysql驱动

	"github.com/ggdream/mochat-server/config"
)


var db	*sql.DB

// Get 获取mysql句柄
func Get() *sql.DB {
	return db
}

// Init 初始化连接mysql
func Init() (err error) {
	db, err = sql.Open("mysql", config.MySQLURI)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1 << 7)
	db.SetMaxIdleConns(1 << 5)

	if err = create(); err != nil {
		return err
	}

	return db.Ping()
}


const (
	sqlTableUser = `create table if not exists user (
    id INT(5) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(10) NOT NULL,
    password VARCHAR(10) NOT NULL
)ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;`

	sqlTableMsg = `create table if not exists msg (
    id INT(10) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    ` + "`from`" + `VARCHAR(10) NOT NULL,
    ` + "`to`" + `VARCHAR(10) NOT NULL
)ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;`
)

func create() error {
	if _, err := Get().Exec(sqlTableUser); err != nil {
		return err
	}

	if _, err := Get().Exec(sqlTableMsg); err != nil {
		return err
	}

	return nil
}
