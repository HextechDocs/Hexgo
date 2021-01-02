package schema

import (
	"github.com/graphql-go/graphql"
	"hextechdocs-be/model"
)

var CategoryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Category",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.ID),
					Description: "The unique ID of this category",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Category).Uuid, nil
					},
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "A user friendly name for the category",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Category).DisplayName, nil
					},
				},
				"slug": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "A URL friendly name for the category",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Category).Slug, nil
					},
				},
				"iconUrl": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "The absolute URL clients should use when fetching the icon for this category",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Category).LogoUrl, nil
					},
				},
				"readmeUrl": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "The absolute URL for the category description document clients should use",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Category).ReadmeUrl, nil
					},
				},
				"subcategories": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(SubcategoryType)),
					Description: "A list of every subcategory for this category",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						category := p.Source.(model.Category)
						return category.GetSubcategories(), nil
					},
				},
			},
		})