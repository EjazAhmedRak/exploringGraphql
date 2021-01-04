package graphqltypes

import (
	"github.com/graphql-go/graphql"
)

// UserAccessType : User struct to read/write user information
var UserAccessType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Access",
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
