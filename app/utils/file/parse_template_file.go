package fileutil

import (
	"bytes"
	"context"
	"html/template"
)

func ParseTemplateFile(ctx context.Context, dir string, templateData any) (string, error) {
	t, err := template.ParseFiles(dir)
	if err != nil {
		return "", err
	}

	writer := bytes.NewBuffer([]byte{})
	err = t.Execute(writer, templateData)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}
