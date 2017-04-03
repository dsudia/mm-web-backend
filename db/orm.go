package db

import (
	"github.com/astaxie/beego/orm"
	// driver for database
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "dbname=test", 30)
	orm.RegisterModel(new(Educator), new(School), new(SchoolMatchingProfile), new(Match))
}

// O is base orm to be used elsewhere
var O = orm.NewOrm()
