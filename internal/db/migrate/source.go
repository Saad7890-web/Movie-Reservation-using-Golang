package migrate

import (
	"os"
	"path/filepath"
	"sort"
)

func LoadSchemaFiles(dir string) ([]string, error){
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".sql"{
			files = append(files, filepath.Join(dir, e.Name()))
		}

	
	}
	sort.Strings((files))
	return files, nil
}