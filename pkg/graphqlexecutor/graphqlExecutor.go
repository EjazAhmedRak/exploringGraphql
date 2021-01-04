package graphqlexecutor

import (
	"fmt"
	dbutils "gographql/pkg/dbutils"
	"gographql/pkg/graphqlstructs"
	"gographql/pkg/graphqltypes"

	"github.com/graphql-go/graphql"
)

// Data : User struct to read/write user information
var Data map[string]graphqlstructs.User

// Db DB Connection
var Db = dbutils.DbConnection()

// Schema : User struct to read/write user information
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: QueryType,
	},
)

// ExecuteQuery : User struct to read/write user information
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

// QueryType : User struct to read/write user information
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: graphqltypes.UserType,
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: userqueryresolver,
			},
		},
	},
)

var userqueryresolver = func(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOk := p.Args["uuid"].(string)
	if isOk {
		searchParam := graphqlstructs.User{ UUID: idQuery}
		user := graphqlstructs.User{}
		Db.Where(&searchParam).Find(&user)
		return user, nil
	}
	return nil, nil
}
