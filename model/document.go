package model

import (
	"errors"
	"strings"
	"time"
)

type Document struct {
	Id            int64
	Uuid          string       `xorm:"char(36) notnull unique index"`
	Slug          string       `xorm:"varchar(255) notnull index"`
	Subcategory   *Subcategory `xorm:"-"`
	SubcategoryId int64
	Title         string `xorm:"varchar(255) notnull"`
	Tags          []string
	Content       string   `xorm:"text notnull"`
	Path          string   `xorm:"varchar(255) notnull"`
	Hash          string   `xorm:"varchar(255) notnull"`
	Hidden        bool     `xorm:"default false"`
	Authors       []Author `xorm:"-"`
	AuthorIds     []int64
	Markers       []Marker `xorm:"-"`
	MarkersIds    []int64
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
	Version       int32     `xorm:"version"`
}

func NewDocument(document *Document) (int64, error) {
	return x.Insert(document)
}

func UpdateDocument(document *Document) (int64, error) {
	return x.ID(document.Id).Update(document)
}

func (d *Document) GetSubcategory() *Subcategory {
	if d.Subcategory == nil {
		var subcategory Subcategory
		success, err := x.Where("id = ?", d.SubcategoryId).Get(&subcategory)

		if !success || err != nil {
			return nil
		}

		d.Subcategory = &subcategory
		return &subcategory
	} else {
		return d.Subcategory
	}
}

func (d *Document) GetCategory() *Category {
	return d.GetSubcategory().GetParentCategory()
}

func (d *Document) GetAuthors() []Author {
	if d.Authors == nil {
		authors := make([]Author, len(d.AuthorIds))
		session := x.NewSession()

		for index, value := range d.AuthorIds {
			var author Author
			_, err := session.Where("id = ?", value).Get(&author)
			if err != nil {
				authors[index] = author
			}
		}

		d.Authors = authors
		return authors
	} else {
		return d.Authors
	}
}

func (d *Document) GetMarkers() []Marker {
	if d.Markers == nil {
		markers := make([]Marker, len(d.MarkersIds))

		for index, value := range d.MarkersIds {
			var marker Marker
			_, err := x.Where("id = ?", value).Get(&marker)

			if err == nil {
				markers[index] = marker
			}
		}

		d.Markers = markers
		return markers
	} else {
		return d.Markers
	}
}

func GetDocumentById(id int64) (*Document, error) {
	var document Document
	success, err := x.Where("id = ?", id).Get(&document)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, nil
	}

	return &document, nil
}

func GetDocumentBySlug(slug string) (*Document, error) {
	slugPieces := strings.Split(slug, ".")
	if len(slugPieces) != 2 {
		return nil, errors.New("invalid slug")
	}

	var document Document
	success, err := x.Where("slug = ?", slugPieces[1]).Get(&document)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, nil
	}

	return &document, nil
}

func GetDocumentByFilePath(path string) (*Document, error) {
	var document Document
	success, err := x.Where("path = ?", path).Get(&document)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, nil
	}

	return &document, nil
}

func GetDocumentsForSubcategory(subcategoryId int64) ([]Document, error) {
	var documents []Document
	err := x.Where("subcategory_id = ?", subcategoryId).Find(&documents)

	if err != nil {
		return nil, err
	}

	return documents, nil
}

func SearchForDocuments(query string) ([]Document, error) {
	var documents []Document
	err := x.Where("content @@ to_tsquery(?)", query).Find(&documents)

	if err != nil {
		return nil, err
	}

	return documents, nil
}
