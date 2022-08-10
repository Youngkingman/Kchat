package global

import (
	"github.com/Youngkingman/Kchat/kchat/pkg/dbutil"
	"github.com/jmoiron/sqlx"
)

var (
	MySQL *sqlx.DB
	Redis dbutil.RedisCli
)
