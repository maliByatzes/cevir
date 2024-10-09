package epub

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Rootfile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

type Rootfiles struct {
	Rootfile []Rootfile `xml:"rootfile"`
}

type Container struct {
	XMLName   xml.Name  `xml:"container"`
	Rootfiles Rootfiles `xml:"rootfiles"`
	Version   string    `xml:"version,attr"`
	Xmlns     string    `xml:"type,attr"`
}

func DecodeContainerXML(dir string) (*Container, error) {
	containerFilepath := fmt.Sprintf("%s/META-INF/container.xml", dir)

	data, err := os.ReadFile(containerFilepath)
	if err != nil {
		return nil, fmt.Errorf("error readig file: %w\n", err)
	}

	var container Container

	err = xml.Unmarshal(data, &container)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling xml: %w", err)
	}

	return &container, nil
}
