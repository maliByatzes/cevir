package epub

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Epub struct {
	FilePath    string
	ExtractDir  string
	OpfFilepath string
}

func NewEpub() *Epub {
	return &Epub{}
}

func (e *Epub) ExtractEpubFile(fileName string) error {
	e.FilePath = fileName

	if err := confirmEpubFileExtension(fileName); err != nil {
		return err
	}

	dir, err := extractEpubZip(fileName)
	if err != nil {
		return err
	}
	fmt.Println("extract dir:", dir)
	e.ExtractDir = dir // NOTE: remember to clean this temp dir

	if err := confirmEpubFileType(dir); err != nil {
		return err
	}

	container, err := DecodeContainerXML(dir)
	if err != nil {
		return err
	}
	e.OpfFilepath = fmt.Sprintf("%s/%s", dir, container.Rootfiles.Rootfile[0].FullPath)

	return nil
}

func (e *Epub) ValidatePackageDocument() error {
  p, err := DecodePackageXML(e.OpfFilepath)
  if err != nil {
    return err
  }

  if err := validateMetadata(p); err != nil {
    return err
  }

  return nil
}

func confirmEpubFileExtension(fileName string) error {
	ext := strings.Split(fileName, ".")

	if ext[len(ext)-1] != "epub" {
		return errors.New("Invalid file extension")
	}

	return nil
}

func extractEpubZip(fileName string) (string, error) {
	s := strings.Split(fileName, "/")

	pattern := fmt.Sprintf("*%s", s[len(s)-1])
	dir, err := os.MkdirTemp("", pattern)
	if err != nil {
		return "", err
	}

	unzipCmd := exec.Command("unzip", "-q", "-d", dir, fileName)
	unzipCmdOutput, err := unzipCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("unzip command failed: %s\n", string(unzipCmdOutput))
	}

	return dir, nil
}

func confirmEpubFileType(dir string) error {
	mimetypeFilename := fmt.Sprintf("%s/mimetype", dir)

	content, err := extractFileContents(mimetypeFilename)
	if err != nil {
		return err
	}

	if content != "application/epub+zip" {
		return errors.New("invalid epub file type.")
	}

	return nil
}

func extractFileContents(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("read file failed: %s\n", err)
	}
	return string(content), nil
}

func validateMetadata(p *Package) error {
  if p.Metadata.Identifier.Value == "" {
    return errors.New("invalid metadata: identifier missing.")
  }

  if p.Metadata.Title.Value == "" {
    return errors.New("invalid metdata: title missing.")
  }

  if p.Metadata.Language.Value == "" {
    return errors.New("invalid metadata: langauge missing")
  }

  // Compulsory for EPUB 2
  /*
  if p.Metadata.Creator.Value == "" {
    return errors.New("invalid metadata: creator missing")
  }*/

  return nil
}
