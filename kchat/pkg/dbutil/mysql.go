package dbutil

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

var TablePrefix string

// 初始化sql链接，幂等操作，仅仅连接一次
func NewSQL(config *setting.DatabaseSettingS) (db *sqlx.DB, err error) {
	var dbonce sync.Once
	dbonce.Do(func() {
		// ："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
		db, err = sqlx.Open(
			config.DBType,
			fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=%s",
				config.UserName,
				config.Password,
				config.Host,
				config.DBName,
				config.Charset,
			),
		)
		if err != nil {
			log.Printf("get mysql database error: %s", err)
		} else {
			//db.SetConnMaxLifetime(time.Duration(config.ConnLifeTime) * time.Second)
			db.SetMaxIdleConns(config.MaxIdleConns)
			db.SetMaxOpenConns(config.MaxOpenConns)
			db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
			TablePrefix = config.TablePrefix
		}
	})

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return
}

//global func for users
func Prefix(str string) string {
	return strings.Replace(str, "#__", TablePrefix, -1)
}

//UnPrefix change the real sql with prefix to relative one
func UnPrefix(str string) string {
	return strings.Replace(str, TablePrefix, "#__", 1)
}

//GetPrefix get sql prefix
func GetPrefix() string {
	return TablePrefix
}
