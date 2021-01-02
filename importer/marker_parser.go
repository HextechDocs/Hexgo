package importer

import (
	"github.com/sirupsen/logrus"
	"hextechdocs-be/logger"
	"hextechdocs-be/model"
	"strings"
)

// parseMarkers converts the input string we get from the front matter into database friendly IDs
func parseMarkers(markersInput string) []int64 {
	markersString := strings.Split(strings.TrimSpace(markersInput), ",")

	var markersInt = make([]int64, 0)
	for i := 0; i < len(markersString); i++ {
		marker, err := model.GetMarkerBySlug(markersString[i])

		if err != nil {
			logger.HexLogger.WithFields(logrus.Fields{
				"error": err,
			}).Error("Error while fetching marker data!")
		} else if marker != nil {
			markersInt = append(markersInt, marker.Id)
		} else {
			logger.HexLogger.WithFields(logrus.Fields{
				"slug": markersString[i],
			}).Error("There is no marker with this slug!")
		}
	}

	return markersInt
}
