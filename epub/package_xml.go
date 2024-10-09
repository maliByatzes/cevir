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
	Identifier   DCElement `xml:"dc:identifier"`
	Title        DCElement `xml:"dc:title"`
	Language     DCElement `xml:"dc:language"`
	Date         string    `xml:"dc:date"`
	Creator      DCElement `xml:"dc:creator"`
	Contributors []string  `xml:"dc:contributor"`
	Publisher    string    `xml:"dc:publisher"`
	Rights       string    `xml:"rights"`
	Metas        []Meta    `xml:"meta"`
	Links        []Link    `xml:"link,omitempty"`
}

type Package struct {
	XMLName          xml.Name `xml:"package"`
	Xmlns            string   `xml:"xmlns,attr"`
	XmlnsDC          string   `xml:"xmlns:dcterms,attr"`
	Version          string   `xml:"version,attr"`
	Lang             string   `xml:"xml:lang,attr"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`
	Metadata         Metadata `xml:"metadata"`
	Manifest         Manifest `xml:"manifest"`
	Spine            Spine    `xml:"spine"`
}

func DecodePackageXML(fileName string) (*Package, error) {
  data, err := os.ReadFile(fileName)
  if err != nil {
    return nil, fmt.Errorf("error reading file: %w\n", err);
  }

  var p Package

  err = xml.Unmarshal(data, &p)
  if err != nil {
    return nil, fmt.Errorf("error unmarshalling xml: %w\n", err);
  }

  return &p, nil
}
