package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// this
	oldModule, err := readGoModule("go.mod")
	if err != nil {
		log.Fatalf("read go.mod: %v", err)
	}
	log.Printf("current module: %s", oldModule)

	fmt.Print("New path name: ")
	newPath, err := readLine(os.Stdin)
	if err != nil {
		log.Fatalf("read input: %v", err)
	}
	newPath = strings.TrimSpace(newPath)
	if newPath == "" {
		log.Fatal("path cannot be empty")
	}
	if err := miseSet("env.PKG_PATH", newPath); err != nil {
		log.Fatalf("set PKG_PATH: %v", err)
	}

	app, err := miseGet("env.APP_NAME")
	if err != nil {
		log.Fatalf("read APP_NAME: %v", err)
	}
	provider, err := miseGet("env.GIT_PROVIDER")
	if err != nil {
		log.Fatalf("read GIT_PROVIDER: %v", err)
	}
	if app == "" || provider == "" {
		log.Fatal("APP_NAME and GIT_PROVIDER must be set before running setup:gopkg")
	}

	newModule := fmt.Sprintf("%s/%s/%s", provider, newPath, app)
	log.Printf("new module: %s", newModule)

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

func miseGet(key string) (string, error) {
	out, err := exec.Command("mise", "config", "get", key).Output()
	if err != nil {
		return "", fmt.Errorf("mise config get %s: %w", key, err)
	}
	return strings.TrimSpace(string(out)), nil
}

func miseSet(key, value string) error {
	arg := fmt.Sprintf("%s=%s", key, value)
	cmd := exec.Command("mise", "config", "set", arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mise config set %s: %w", arg, err)
	}
	return nil
}

func readLine(f *os.File) (string, error) {
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	if err != nil && line == "" {
		return "", err
	}
	return strings.TrimRight(line, "\r\n"), nil
}
