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
		fmt.Printf("wrong result, unexoected errors: %v", result.Errors)
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
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: idqueryresolver,
			},
		},
	},
)

var idqueryresolver = func(p graphql.ResolveParams) (interface{}, error) {
	var users []graphqlstructs.User
	idQuery, isOk := p.Args["id"].(string)
	if isOk {
		rows, dberr := Db.Raw("SELECT user_uuid as id, adfs_name_id as email, rid_last_name as name FROM users_tbl WHERE user_uuid = ?", idQuery).Rows()
		if dberr != nil {
			return nil, dberr
		}
		for rows.Next() {
			// ScanRows scan a row into user
			var user graphqlstructs.User
			Db.ScanRows(rows, &user)
			// do something
			users = append(users, user)
		}
		return users[0], nil
	}
	return nil, nil
}
