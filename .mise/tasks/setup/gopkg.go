package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	mise "github.com/kkhnifes/kratos-layout/.mise/internal/mise"
)

func main() {
	oldModule, err := readGoModule("go.mod")
	if err != nil {
		log.Fatalf("read go.mod: %v", err)
	}

	pkgPath, _ := mise.Get("env.PKG_PATH")
	app, _ := mise.Get("env.APP_NAME")
	provider, _ := mise.Get("env.GIT_PROVIDER")
	if app == "" || provider == "" {
		log.Fatal("APP_NAME and GIT_PROVIDER must be set before running setup:gopkg")
	}

	newPath, changed, err := mise.Prompt("New path name", pkgPath)
	if err != nil {
		log.Fatalf("read input: %v", err)
	}
	if !changed {
		log.Printf("module path kept: %s", oldModule)
		return
	}
	if newPath == "" {
		log.Fatal("path cannot be empty")
	}
	if err := mise.Set("env.PKG_PATH", newPath); err != nil {
		log.Fatalf("set PKG_PATH: %v", err)
	}

	newModule := fmt.Sprintf("%s/%s/%s", provider, newPath, app)
	if err := replaceOnce("go.mod", "module "+oldModule, "module "+newModule); err != nil {
		log.Fatalf("update go.mod: %v", err)
	}
	n, err := replaceInTree(".", oldModule, newModule)
	if err != nil {
		log.Fatalf("rewrite imports: %v", err)
	}
	log.Printf("rewrote %d files", n)

	tidy := exec.Command("go", "mod", "tidy")
	tidy.Stdout = os.Stdout
	tidy.Stderr = os.Stderr
	if err := tidy.Run(); err != nil {
		log.Fatalf("go mod tidy: %v", err)
	}
	log.Println("Done.")
}

func readGoModule(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", fmt.Errorf("no module line in %s", path)
}

func replaceOnce(path, old, new string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Contains(data, []byte(old)) {
		return fmt.Errorf("%s: %q not found", path, old)
	}
	updated := bytes.Replace(data, []byte(old), []byte(new), 1)
	return os.WriteFile(path, updated, 0o644)
}

func replaceInTree(root, old, new string) (int, error) {
	count := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if !bytes.Contains(data, []byte(old)) {
			return nil
		}
		updated := bytes.ReplaceAll(data, []byte(old), []byte(new))
		if err := os.WriteFile(path, updated, info.Mode()); err != nil {
			return err
		}
		count++
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
