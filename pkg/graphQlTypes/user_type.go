package graphqltypes

import (
	dbutils "gographql/pkg/dbutils"
	models "gographql/pkg/graphqlstructs"

	"github.com/graphql-go/graphql"
)

// Db DB Connection
var Db = dbutils.DbConnection()

// UserType : User struct to read/write user information
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"accesses" : &graphql.Field{
				Type: graphql.NewList(UserAccessType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User) ; ok {
						return accessQueryResolver(user)
					}
					return nil, nil
				},
			},
		},
	},
)

var accessQueryResolver = func(user models.User) ([]models.UserAccesses, error) {
		searchParam := models.UserAccesses{}
		searchParam.UserUUID = &user.UUID
		userAccesses := []models.UserAccesses{}
		Db.Where(&searchParam).Find(&userAccesses)
		return userAccesses, nil
}
