package main

import (
	"fmt"
	dbutils "gographql/pkg/dbutils"
	models "gographql/pkg/graphqlstructs"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// TO DO : Add configuration support

	var db = dbutils.DbConnection()

	if db.Migrator().HasTable(&models.User{}) == false {
		db.Migrator().CreateTable(&(models.User{}))
	}

	if db.Migrator().HasTable(&models.BusinessEntity{}) == false {
		db.Migrator().CreateTable(&(models.BusinessEntity{}))
	}
	
	if db.Migrator().HasTable(&models.UserAccesses{}) == false {
		db.Migrator().CreateTable(&(models.UserAccesses{}))
	}

	db.Exec(dbutils.UsersInsert)
	db.Exec(dbutils.EntityInsert)
	db.Exec(dbutils.UserAccessesInsert)

	var users []models.User
	db.Find(&users)

	fmt.Println("There are", len(users), "user records in the table.")

}
