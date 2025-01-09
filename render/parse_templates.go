package render

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed views/*
var views embed.FS

func parseTemplates() (*template.Template, error) {
	// List all files in the embedded views directory.
	// You can manually get the list of HTML files from the views directory
	var filenames []string
	err := filepath.WalkDir("views", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Only add .html files
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".html") {
			filenames = append(filenames, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Read and parse each embedded file
	var tpl *template.Template
	for _, filename := range filenames {
		data, err := views.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("could not read embedded file %s: %v", filename, err)
		}

		// If this is the first file, initialize the template, otherwise add it
		if tpl == nil {
			tpl, err = template.New(filepath.Base(filename)).Parse(string(data))
		} else {
			tpl, err = tpl.New(filepath.Base(filename)).Parse(string(data))
		}
		if err != nil {
			return nil, fmt.Errorf("could not parse template %s: %v", filename, err)
		}
	}

	return tpl, nil
}
