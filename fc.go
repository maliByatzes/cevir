package cevir

import "github.com/maliByatzes/cevir/epub"

func ConvertEpubToPDF(fileName string) error {
	epub := epub.NewEpub()

	if err := epub.ExtractEpubFile(fileName); err != nil {
		return err
	}

	return nil
}

func ConvertPDFToEpub(fileName string) error {
	return nil
}
