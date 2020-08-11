package SqlControl

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	db     *sql.DB
	config mysql.Config
}

func (m *MysqlDB) Open() (err error) {
	log.Println("conn: ", m.config.FormatDSN())
	m.db, err = sql.Open("mysql", m.config.FormatDSN())
	if err != nil {
		log.Println(err)
		m.db = nil
		return errors.New("Open fail")
	}
	return nil
}

func (m *MysqlDB) Close() {
	m.db.Close()
}

func (m *MysqlDB) Ping() {
	err := m.db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func (m *MysqlDB) Query(cmd string) (rows *sql.Rows, err error) {
	rows, err = m.db.Query(cmd)
	return rows, err
}

func (m *MysqlDB) Insert(cmds []string) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, cmd := range cmds {
		_, err = tx.Exec(cmd)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return err
}

func NewDataBaseConnect(user string, password string) MysqlDB {
	db := MysqlDB{}
	db.config = mysql.Config{
		User:                 user,
		Passwd:               password,
		Addr:                 "localhost",
		Net:                  "tcp",
		DBName:               "mydata",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return db
}
