package schema

import (
	"github.com/graphql-go/graphql"
	"hextechdocs-be/model"
	"strconv"
	"time"
)

var DocumentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Document",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.ID),
					Description: "A unique document ID",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Document).Uuid, nil
					},
				},
				"slug": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "A URL friendly document name",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						doc := p.Source.(model.Document)
						docSlug := strconv.FormatInt(doc.Id, 10) + "." + doc.Slug
						return docSlug, nil
					},
				},
				"category": &graphql.Field{
					Type: graphql.NewNonNull(CategoryType),
					Description: "The category of this document",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						doc := p.Source.(model.Document)
						return doc.GetCategory(), nil
					},
				},
				"subcategory": &graphql.Field{
					Type: graphql.NewNonNull(SubcategoryType),
					Description: "The subcategory of this document",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						doc := p.Source.(model.Document)
						return doc.GetSubcategory(), nil
					},
				},
				"title": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "The user friendly title of the document",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Document).Title, nil
					},
				},
				"authors": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(AuthorType)),
					Description: "A list of every contributor",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						doc := p.Source.(model.Document)
						return doc.GetAuthors(), nil
					},
				},
				"content": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "The contents of this document in markdown",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Document).Content, nil
					},
				},
				"tags": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
					Description: "A list of every tag attached to the document",
					DeprecationReason: "Unused variable",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Document).Tags, nil
					},
				},
				"markers": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(MarkerType)),
					Description: "A list of every marker attached to this document",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						doc := p.Source.(model.Document)
						return doc.GetMarkers(), nil
					},
				},
				"createdAt": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "Unix timestamp on when the document was created",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Document).CreatedAt.Format(time.RFC3339), nil
					},
				},
				"updatedAt": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "Unix timestamp on when the document was last updated",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Document).UpdatedAt.Format(time.RFC3339), nil
					},
				},
			},
		})