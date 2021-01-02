package schema

import (
	"github.com/graphql-go/graphql"
	"hextechdocs-be/model"
)

var AuthorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"github": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
					Description: "The Github handle of the author",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Author).GithubUsername, nil
					},
				},
				"name": &graphql.Field{
					Type: graphql.String,
					Description: "The Github display name of the author, returns null if the user decided to keep it private",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						author := p.Source.(model.Author)
						if author.ShouldDisplayName {
							return author.GithubName, nil
						} else {
							return nil, nil
						}
					},
				},
			},
		})