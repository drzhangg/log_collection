package init

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func InitMysql() (db *gorm.DB, err error) {

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", ConfigInit.Mysql.Username, ConfigInit.Mysql.Password, ConfigInit.Host, ConfigInit.Port, ConfigInit.Database)
	db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		err = fmt.Errorf("connect mysql failed, err:%v", err)
		log.Println(err)
		return
	}
	log.Println("mysql init success")
	return
}
