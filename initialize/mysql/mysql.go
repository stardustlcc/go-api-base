package mysql

import (
	"fmt"
	"go-api-base/pkg/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	GetDB() *gorm.DB
	DBClose() error
}

type dbRepo struct {
	db *gorm.DB
}

func (d *dbRepo) GetDB() *gorm.DB {
	return d.db
}

func (d *dbRepo) DBClose() error {
	return d.db.DB().Close()
}

func NewMysql() (Repo, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	sqlConf := global.Conf.Database
	db, err := gorm.Open(sqlConf.DBType, fmt.Sprintf(s,
		sqlConf.UserName,
		sqlConf.PassWord,
		sqlConf.Host,
		sqlConf.DBName,
		sqlConf.Charset,
		sqlConf.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.Conf.Server.RunModel == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)

	//设置数据库系统连接数
	db.DB().SetMaxIdleConns(sqlConf.MaxIdleConns)
	db.DB().SetMaxOpenConns(sqlConf.MaxOpenConns)

	return &dbRepo{db}, nil
}
