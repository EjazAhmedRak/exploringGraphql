package graphqltypes

import (
	"github.com/graphql-go/graphql"
)

// BusinessEntityType : User struct to read/write user information
var BusinessEntityType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BusinessEntity",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"entityType": &graphql.Field{
				Type: graphql.String,
			},
			"entityName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
