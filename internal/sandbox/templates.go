package sandbox

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type config struct {
	BaseImage string
}

func renderFile(src, dst string, data any) error {
	tmpl, err := template.ParseFiles(src)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("render template: %w", err)
	}

	if err := os.WriteFile(dst, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("write rendered file: %w", err)
	}

	return nil
}
