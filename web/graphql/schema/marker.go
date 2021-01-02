package schema

import (
	"github.com/graphql-go/graphql"
	"hextechdocs-be/model"
)

var MarkerType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Marker",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.ID),
					Description: "Unique and URL friendly marker ID",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Marker).Uuid, nil
					},
				},
				"displayName": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "The display name clients should use when displaying this marker",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Marker).DisplayName, nil
					},
				},
			},
		})