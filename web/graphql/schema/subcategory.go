package schema

import (
	"github.com/graphql-go/graphql"
	"hextechdocs-be/model"
)

var SubcategoryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Subcategory",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.ID),
					Description: "A globally unique subcategory ID",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Subcategory).Uuid, nil
					},
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "A user friendly name for the subcategory",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Subcategory).DisplayName, nil
					},
				},
				"slug": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "A URL friendly name for the subcategory",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Subcategory).Slug, nil
					},
				},
			},
		})