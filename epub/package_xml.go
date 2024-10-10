package epub

import (
	"encoding/xml"
	"fmt"
	"os"
)

type DCElement struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Meta struct {
	Property string `xml:"property,attr"`
	Value    string `xml:",chardata"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type ManifestItem struct {
	ID         string `xml:"id,attr"`
	Href       string `xml:"href,attr"`
	MediaType  string `xml:"media-type,attr"`
	Properties string `xml:"properties,attr,omitempty"`
}

type Manifest struct {
	Items []ManifestItem `xml:"item"`
}

type SpineItemRef struct {
	IDRef  string `xml:"idref,attr"`
	Linear string `xml:"linear,attr,omitempty"`
}

type Spine struct {
	ItemRefs []SpineItemRef `xml:"itemref"`
}

type Metadata struct {
	Identifier   DCElement `xml:"identifier"`
	Title        DCElement `xml:"title"`
	Language     DCElement `xml:"language"`
	Date         string    `xml:"date"`
	Creator      DCElement `xml:"creator"`
	Contributors []string  `xml:"contributor"`
	Publisher    string    `xml:"publisher"`
	Rights       string    `xml:"rights"`
	Metas        []Meta    `xml:"meta"`
	Links        []Link    `xml:"link,omitempty"`
}

type Package struct {
	XMLName          xml.Name `xml:"package"`
	Xmlns            string   `xml:"xmlns,attr"`
	XmlnsDC          string   `xml:"dc,attr"`
	XmlnsDcTerms     string   `xml:"dcterms,attr"`
	Version          string   `xml:"version,attr"`
	Lang             string   `xml:"lang,attr"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`
	Metadata         Metadata `xml:"metadata"`
	Manifest         Manifest `xml:"manifest"`
	Spine            Spine    `xml:"spine"`
}

func DecodePackageXML(fileName string) (*Package, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w\n", err)
	}

	var p Package

	err = xml.Unmarshal(data, &p)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling xml: %w\n", err)
	}

	return &p, nil
}
