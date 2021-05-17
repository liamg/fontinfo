package fontinfo

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Font struct {
	Family string
	Style  string
	Path   string
}

var fontDirs = []string{
	"~/.fonts",
	"~/.local/share/fonts",
	"/usr/local/share/fonts",
	"/usr/share/fonts",
	filepath.Join(os.Getenv("XDG_DATA_HOME"), "fonts"),
	filepath.Join(os.Getenv("XDG_DATA_DIRS"), "fonts"),
}

var validExtensions = []string{
	".ttf",
	".otf",
}

func List() ([]Font, error) {

	var fonts []Font
	meta := make(map[string]*fontMetadata)

	for _, dir := range fontDirs {

		if info, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		} else if !info.IsDir() {
			continue
		}

		if err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if _, ok := meta[path]; ok {
				return nil
			}
			ext := filepath.Ext(path)
			for _, valid := range validExtensions {
				if strings.EqualFold(ext, valid) {
					f, err := os.Open(path)
					if err != nil {
						return err
					}
					defer f.Close()
					metadata, err := readMetadata(f)
					if err != nil {
						return err
					}
					meta[path] = metadata
					return nil
				}
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	for path, metadata := range meta {
		fonts = append(fonts, Font{
			Family: metadata.FontFamily,
			Style:  metadata.FontStyle,
			Path:   path,
		})
	}

	return fonts, nil
}
