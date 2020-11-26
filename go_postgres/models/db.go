package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var ormObject orm.Ormer

func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=postgres password=1123 dbname=db_go host=localhost sslmode=disable")
	orm.RegisterModel(new(Users))

	name := "default"
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	ormObject = orm.NewOrm()
}

func GetOrmObject() orm.Ormer {
	return ormObject
}
