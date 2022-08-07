package mysql

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Youngkingman/Kchat/kchat/pkg/setting"
	_ "github.com/go-sql-driver/mysql" //import mysql impolementation for sql
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

type IDB interface {
	GetDB() *sqlx.DB
	Prefix(str string) string
	UnPrefix(str string) string
	GetPrefix() string
}

type sqlServer struct {
	Config setting.DatabaseSettingS
	DB     *sqlx.DB
}

//init 初始化sql
func (sql *sqlServer) init(config setting.DatabaseSettingS) error {
	var dbonce sync.Once
	var db *sqlx.DB
	var err error
	dbonce.Do(func() {
		db, err = sqlx.Open(
			config.DBType,
			fmt.Sprintf(
				"%s:%s@tcp(%s)/%s",
				config.UserName,
				config.Password,
				config.Host,
				config.DBName,
			),
		)
		if err != nil {
			log.Printf("get mysql database error: %s", err)
		} else {
			//db.SetConnMaxLifetime(time.Duration(config.ConnLifeTime) * time.Second)
			db.SetMaxIdleConns(config.MaxIdleConns)
			db.SetMaxOpenConns(config.MaxOpenConns)
			db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
		}
	})

	sql.DB = db
	return err
}

//GetDB GetDB
func (sql *sqlServer) GetDB() *sqlx.DB {
	return sql.DB
}

func New(config setting.DatabaseSettingS) (sql IDB, err error) {
	s := &sqlServer{
		Config: config,
	}
	err = s.init(config)
	sql = s
	return
}

//Prefix change the relative sql to real sql with prefix
func (sql sqlServer) Prefix(str string) string {
	return strings.Replace(str, "#__", sql.Config.TablePrefix, -1)
}

//UnPrefix change the real sql with prefix to relative one
func (sql sqlServer) UnPrefix(str string) string {
	return strings.Replace(str, sql.Config.TablePrefix, "#__", 1)
}

//GetPrefix get sql prefix
func (sql sqlServer) GetPrefix() string {
	return sql.Config.TablePrefix
}
