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
