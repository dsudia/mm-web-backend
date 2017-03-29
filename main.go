package main

import (
	"github.com/astaxie/beego/arm"
	_ "github.com/lib/pq"
)

funct init() {
	orm.RegisterDriver("pq", orm.DRPostgres)
	orm.RegisterDataBase("default", "pq", "root:root@/my_db?charset=utf8", 30)
}

func main() {
	
}
