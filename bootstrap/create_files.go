package bootstrap

import (
	"fmt"
	"log"
	"os"

	"github.com/paoloanzn/go-bootstrap/config"
	"github.com/paoloanzn/go-bootstrap/format"
	"github.com/paoloanzn/go-bootstrap/parsing"
)

func CreateDir(path string, abortIfFailed bool) error {
	if path == "" {
		if abortIfFailed {
			log.Fatalf("Fatal: empty path is invalid\n")
		} else {
			return fmt.Errorf("Empty path is not valid.")
		}
	}

	formattedPath := format.MatchWildCards(path)
	formattedPath = format.FormatPath(formattedPath)

	if _, err := os.Stat(formattedPath); !os.IsNotExist(err) {
		return nil
	}

	err := os.Mkdir(formattedPath, 0755)
	if err != nil {
		if abortIfFailed {
			log.Fatalf("Fatal: %v\n", err)
		} else {
			return err
		}
	}

	fmt.Printf("Created %s\n", formattedPath)
	return nil
}

func CreateFile(path string, abortIfFailed bool) error {
	if path == "" {
		if abortIfFailed {
			log.Fatalf("Fatal: empty path is invalid\n")
		} else {
			return fmt.Errorf("Empty path is not valid.")
		}
	}

	formattedPath := format.MatchWildCards(path)
	formattedPath = format.FormatPath(formattedPath)

	if _, err := os.Stat(formattedPath); !os.IsNotExist(err) {
		return nil
	}

	f, err := os.Create(formattedPath)
	if err != nil {
		if abortIfFailed {
			log.Fatalf("Fatal: %v\n", err)
		} else {
			return err
		}
	}
	defer f.Close()

	fmt.Printf("Created %s\n", formattedPath)
	return nil
}

func TraverseNode(pNode map[string]interface{}, prefixPath string) error {
	for name, value := range pNode {
		if value == "file" {
			fullPath := fmt.Sprintf("%s%s", prefixPath, name)
			err := CreateFile(fullPath, true)
			if err != nil {
				return err
			}

			continue
		}

		asserted, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Error traversing template config file: Invalid structure.")
		}

		fullPath := fmt.Sprintf("%s%s/", prefixPath, name)
		err := CreateDir(fullPath, true)
		if err != nil {
			return err
		}

		err = TraverseNode(asserted, fullPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func Bootstrap(pJsonTemplate *parsing.JSONTemplate) error {
	projectConfig := pJsonTemplate.Config

	projectFolderName, exists := projectConfig["name"]
	if !exists {
		return fmt.Errorf("Error parsing config.name from template config file.")
	}
	config.Cfg.ProjectName = projectFolderName.(string)

	err := CreateDir(projectFolderName.(string), true)
	if err != nil {
		return err
	}

	asserted, ok := pJsonTemplate.Project.(map[string]interface{})
	if !ok {
		log.Fatalf("Invalid project json configuration.")
	}

	rootPath := fmt.Sprintf("%s/", projectFolderName.(string))
	err = TraverseNode(asserted, rootPath)
	if err != nil {
		return err
	}

	return nil
}
