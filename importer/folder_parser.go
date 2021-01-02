package importer

import (
	"github.com/sirupsen/logrus"
	"hextechdocs-be/logger"
	"hextechdocs-be/model"
	"io/ioutil"
	"os"
	"strings"
)

// ParseFolders loops through each category and subcategory and parse every .md file in relevant folders
func ParseFolders() {
	categories, err := model.GetCategories()
	if err != nil {
		logger.HexLogger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Unable to fetch a list of every available category!")
		return
	}

	for _, category := range categories {
		if _, err := os.Stat(category.Slug); os.IsNotExist(err) {
			logger.HexLogger.WithFields(logrus.Fields{
				"error": err,
				"slug":  category.Slug,
			}).Warn("There is no folder for a category!")
			continue
		}

		logger.HexLogger.WithFields(logrus.Fields{
			"category": category.Slug,
		}).Info("Parsing category!")
		subcategories := category.GetSubcategories()
		for _, subcategory := range subcategories {
			if _, err := os.Stat(category.Slug + "/" + subcategory.Slug); os.IsNotExist(err) {
				logger.HexLogger.WithFields(logrus.Fields{
					"category":    category.Slug,
					"subcategory": subcategory.Slug,
				}).Warn("There is no folder for a subcategory!")
				continue
			}
			logger.HexLogger.WithFields(logrus.Fields{
				"category":    category.Slug,
				"subcategory": subcategory.Slug,
			}).Info("Parsing category!")

			files, err := ioutil.ReadDir(category.Slug + "/" + subcategory.Slug)
			if err != nil {
				logger.HexLogger.WithFields(logrus.Fields{
					"category":    category.Slug,
					"subcategory": subcategory.Slug,
				}).Fatal("Unable to read directory metadata for subcategory!")
			}
			for _, f := range files {
				if !strings.HasSuffix(strings.ToLower(f.Name()), ".md") {
					continue
				}
				logger.HexLogger.WithFields(logrus.Fields{
					"category":    category.Slug,
					"subcategory": subcategory.Slug,
					"filename":    f.Name(),
				}).Info("Parsing document!")

				parseDocument(f, subcategory, category)
			}
		}
	}
}
