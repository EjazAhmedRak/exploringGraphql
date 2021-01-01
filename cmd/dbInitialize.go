package main

import (
	"fmt"
	dbutils "gographql/pkg/dbUtils"
	models "gographql/pkg/graphqlstructs"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// TO DO : Add configuration support

	var db = dbutils.DbConnection()

	if db.Migrator().HasTable(&models.User{}) == false {
		db.Migrator().CreateTable(&(models.User{}))
	}

	db.Exec(dbutils.UsersInsert)

	var users []models.User
	db.Find(&users)

	fmt.Println("There are", len(users), "user records in the table.")

}
