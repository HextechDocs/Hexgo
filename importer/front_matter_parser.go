package importer

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

// FrontMatter is the struct that stores everything we support coming from the YAML front matter
type FrontMatter struct {
	Title       string // Title will serve as the title of the document on the site
	Authors     string // Authors is a comma separated list of every github user that contributed to the document and wants to be public
	Tags        string // Tags is a comma separated list of all the keywords related to the document
	Description string // Description is a short description describing what the document is about
	Markers     string // Markers is a comma separated list of each marker slug the document should appear in
}

// parseFrontMatter takes care of unmarshalling the YAML portion of each .md file
func parseFrontMatter(file *os.File) (*FrontMatter, string, error) {
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", err
	}
	documentString := string(bytes)

	pieces := strings.SplitN(documentString, "---", 3)
	if len(pieces) != 3 {
		return nil, "", errors.New("invalid document")
	}

	var frontMatter FrontMatter
	if err := yaml.Unmarshal([]byte(pieces[1]), &frontMatter); err != nil {
		return nil, "", err
	}

	return &frontMatter, pieces[2], nil
}
