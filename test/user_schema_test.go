package graphqltest

import (
	"gographql/pkg/graphqlexecutor"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

type GraphQlTest struct {
	Query     string
	Schema    graphql.Schema
	Expected  interface{}
	Variables map[string]interface{}
}

var Tests = []GraphQlTest{}

func init() {
	Tests = []GraphQlTest{
		{
			Query: `
				query UserNameQuery {
					user(id:"01610666-2c23-49ad-8a8c-5d44f1b1bd5a"){
						name, id
					}
				}
			`,
			Schema: graphqlexecutor.Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"user": map[string]interface{}{
						"name": "Accessers15",
						"id":   "01610666-2c23-49ad-8a8c-5d44f1b1bd5a",
					},
				},
			},
		},
	}
}

func TestQuery(t *testing.T) {
	for _, test := range Tests {
		params := graphql.Params{
			Schema:         test.Schema,
			RequestString:  test.Query,
			VariableValues: test.Variables,
		}
		testGraphql(test, params, t)
	}
}

func testGraphql(test GraphQlTest, p graphql.Params, t *testing.T) {
	result := graphql.Do(p)
	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}
	if !reflect.DeepEqual(result, test.Expected) {
		t.Fatalf("wrong result, query: %v, graphql result diff: %v", test.Query, testutil.Diff(test.Expected, result))
	}
}
