package schema

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"hextechdocs-be/model"
)

var QueryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getMenuItems": &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(CategoryType)),
					Description: "Get all categories for the menu",
					Resolve:     GetMenuItems,
				},
				"getDocumentDetails": &graphql.Field{
					Type:        DocumentType,
					Description: "Get a document based on its ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					DeprecationReason: "Unused resource, replaced by getDocumentBySlug",
					Resolve:           GetDocumentDetails,
				},
				"getDocumentBySlug": &graphql.Field{
					Type:        DocumentType,
					Description: "Get a document based on its URL slug",
					Args: graphql.FieldConfigArgument{
						"slug": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: GetDocumentDetails,
				},
				"getDocumentsForSubcategory": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(DocumentType)),
					Description: "Get all documents in the given subcategory",
					Args: graphql.FieldConfigArgument{
						"subcategoryId": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					Resolve: GetDocumentsForSubcategory,
				},
				"getMarkersForSubcategory": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(MarkerType)),
					Description: "Get all markers in the given subcategory",
					Args: graphql.FieldConfigArgument{
						"subcategoryId": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					Resolve: GetMarkersForSubcategory,
				},
				"narrowSubcategory": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(DocumentType)),
					Description: "Get all documents in the given subcategory that have a specific marker on them",
					Args: graphql.FieldConfigArgument{
						"subcategoryId": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
						"markerId": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					Resolve: NarrowSubcategory,
				},
				"performSearch": &graphql.Field{
					Type: graphql.NewList(graphql.NewNonNull(DocumentType)),
					Description: "Get all documents that match the given search query",
					Args: graphql.FieldConfigArgument{
						"query": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: PerformSearch,
				},
			},
		})

func GetMenuItems(_ graphql.ResolveParams) (interface{}, error) {
	cats, err := model.GetCategories()

	if err != nil {
		return nil, err
	}
	return cats, nil
}

func GetDocumentDetails(params graphql.ResolveParams) (interface{}, error) {
	if params.Args["id"] != nil {
		doc, err := model.GetDocumentById(params.Args["id"].(int64))
		if doc == nil {
			return nil, err
		}
		return *doc, err
	} else if params.Args["slug"] != nil {
		doc, err := model.GetDocumentBySlug(params.Args["slug"].(string))
		if doc == nil {
			return nil, err
		}
		return *doc, err
	}

	return nil, errors.New("invalid parameter")
}

func GetDocumentsForSubcategory(params graphql.ResolveParams) (interface{}, error) {
	subcategory, err := model.GetSubcategoryByUuid(params.Args["subcategoryId"].(string))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	documents, err := model.GetDocumentsForSubcategory(subcategory.Id)

	if err != nil {
		return nil, err
	}

	return documents, nil
}

func GetMarkersForSubcategory(params graphql.ResolveParams) (interface{}, error) {
	subcategory, err := model.GetSubcategoryByUuid(params.Args["subcategoryId"].(string))

	if err != nil {
		return nil, err
	}

	documents, err := model.GetDocumentsForSubcategory(subcategory.Id)

	if err != nil {
		return nil, err
	}

	markers := make([]model.Marker, 0)

	for _, document := range documents {
		for _, marker := range document.GetMarkers() {
			if !contains(markers, marker) {
				markers = append(markers, marker)
			}
		}
	}

	return markers, nil
}

func NarrowSubcategory(params graphql.ResolveParams) (interface{}, error) {
	markerId := params.Args["markerId"].(string)
	subcategory, err := model.GetSubcategoryByUuid(params.Args["subcategoryId"].(string))


	if err != nil {
		return nil, err
	}

	documents, err := model.GetDocumentsForSubcategory(subcategory.Id)

	if err != nil {
		return nil, err
	}

	filteredDocuments := make([]model.Document, 0)

	for _, document := range documents {
		for _, marker := range document.GetMarkers() {
			if marker.Uuid == markerId {
				filteredDocuments = append(filteredDocuments, document)
			}
		}
	}

	return filteredDocuments, nil
}

func PerformSearch(params graphql.ResolveParams) (interface{}, error) {
	documents, err := model.SearchForDocuments(params.Args["query"].(string))

	if err != nil {
		return nil, err
	}

	return documents, nil
}

func contains(s []model.Marker, e model.Marker) bool {
	for _, a := range s {
		if a.Id == e.Id {
			return true
		}
	}
	return false
}