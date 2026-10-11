package loader

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed workspace/**
var workspace embed.FS

func RunDockerImage() error {
	return nil
}

func RestartDocker() error {
	return nil
}

func StopDocker() error {
	return nil
}

func CloseDocker() error {
	return nil
}

func extractWorkspace() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	target := filepath.Join(cwd, "workspace")

	if err := os.MkdirAll(target, 0755); err != nil {
		return "", err
	}

	err = fs.WalkDir(workspace, "workspace", func(
		name string,
		entry fs.DirEntry,
		err error,
	) error {
		if err != nil {
			return err
		}

		if name == "workspace" {
			return nil
		}

		rel := strings.TrimPrefix(name, "workspace/")
		dest := filepath.Join(target, filepath.FromSlash(rel))

		if entry.IsDir() {
			return os.MkdirAll(dest, 0755)
		}

		data, err := workspace.ReadFile(name)
		if err != nil {
			return err
		}

		// ! не перезаписываем существующие файлы
		if _, err := os.Stat(dest); err == nil {
			return nil
		} else if !os.IsNotExist(err) {
			return err
		}

		return os.WriteFile(dest, data, 0644)
	})
	if err != nil {
		return "", err
	}

	return target, nil
}
