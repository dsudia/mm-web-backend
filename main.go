package main

import (
	"github.com/astaxie/beego/arm"
	_ "github.com/lib/pq"
)

funct init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "dbname=test", 30)
	orm.RegisterModel(new(Educator), new(School), new(SchoolMatchingProfile), new(Match))
}

func main() {
	o := orm.NewOrm()
}
