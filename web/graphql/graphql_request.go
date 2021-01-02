package graphql

import (
	"github.com/graphql-go/graphql"
)
import "hextechdocs-be/web/graphql/schema"

var graphqlSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: schema.QueryType,
	},
)

func ExecuteQuery(operationName string, query string, variables map[string]interface{}) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         graphqlSchema,
		OperationName:  operationName,
		RequestString:  query,
		VariableValues: variables,
	})

	return result
}
