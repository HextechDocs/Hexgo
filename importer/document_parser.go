/*
Package importer is responsible for importing the documents into the system.

We ingest our documents from .md files located in predetermined folders.

Here's how it works:
1. Folder parser goes into every relevant folder and calls parseDocument on every .md file in a known folder
2. parseDocument separates the front matter from the actual document, parses it and either updates an existing document or inserts a new one
*/
package importer

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"hextechdocs-be/logger"
	"hextechdocs-be/model"
	"io"
	"os"
	"regexp"
	"strings"
)

// parseDocument takes care of loading the file in, parsing its front matter and inserting or updating the database
func parseDocument(file os.FileInfo, subcategory model.Subcategory, category model.Category) {
	path := category.Slug + "/" + subcategory.Slug + "/" + file.Name()
	f, err := os.Open(path)
	if err != nil {
		logger.HexLogger.WithFields(logrus.Fields{
			"error": err,
			"path":  path,
		}).Error("Unable to open document on disk! ")
		return
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		logger.HexLogger.WithFields(logrus.Fields{
			"error": err,
			"path":  path,
		}).Error("Unable to hash document! ")
		return
	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	_, _ = f.Seek(0, io.SeekStart)

	document, err := model.GetDocumentByFilePath(path)
	if err != nil {
		logger.HexLogger.WithFields(logrus.Fields{
			"error": err,
			"path":  path,
		}).Error("Unable to fetch document from database! ")
		return
	}

	frontmatter, body, err := parseFrontMatter(f)
	if err != nil {
		return
	}

	if strings.HasPrefix(body, "---") {
		body = strings.Replace(body, "---", "", 1)
	}

	if document != nil {
		if document.Hash == hash {
			return
		}

		document.Hash = hash
		document.Title = frontmatter.Title
		document.Tags = strings.Split(frontmatter.Tags, ",")
		document.Content = body
		document.MarkersIds = parseMarkers(frontmatter.Markers)

		_, _ = model.UpdateDocument(document)
		logger.HexLogger.WithFields(logrus.Fields{
			"path":  path,
			"title": frontmatter.Title,
		}).Info("Updated document: ")
	} else {
		// Slug processing (all lowercase, no special characters and a dash instead of spaces)
		slug := strings.Replace(strings.ToLower(file.Name()), ".md", "", -1)
		slug = strings.Replace(slug, " ", "-", -1)

		regex := regexp.MustCompile("[^A-Za-z0-9-_]")
		slug = regex.ReplaceAllString(slug, "")

		document := &model.Document{
			Uuid:          uuid.New().String(),
			Slug:          slug,
			SubcategoryId: subcategory.Id,
			Title:         frontmatter.Title,
			Tags:          strings.Split(frontmatter.Tags, ","),
			Content:       body,
			Path:          path,
			Hash:          hash,
			Hidden:        false,
			AuthorIds:     []int64{},
			MarkersIds:    parseMarkers(frontmatter.Markers),
		}

		if _, err := model.NewDocument(document); err != nil {
			return
		}

		logger.HexLogger.WithFields(logrus.Fields{
			"path":  path,
			"title": frontmatter.Title,
		}).Info("New document: ")
	}
}
