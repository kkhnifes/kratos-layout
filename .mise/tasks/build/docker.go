package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if _, err := exec.LookPath("docker"); err != nil {
		log.Fatalf("docker not found in PATH: %v", err)
	}

	appName, ok := os.LookupEnv("APP_NAME")
	if !ok || appName == "" {
		log.Fatal("APP_NAME is not set")
	}
	registry := os.Getenv("REGISTRY")
	pkgPath := os.Getenv("PKG_PATH")
	if registry == "" || pkgPath == "" {
		log.Fatal("REGISTRY and PKG_PATH must be set")
	}

	version, err := gitOut("describe", "--tags", "--always")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if version == "" {
		log.Fatal("git describe: empty output")
	}
	vcsRef, err := gitOut("rev-parse", "--short", "HEAD")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if vcsRef == "" {
		log.Fatal("git rev-parse: empty output")
	}

	tag := fmt.Sprintf("%s/%s/%s:%s", registry, pkgPath, appName, version)
	log.Printf("building %s", tag)

	cmd := exec.Command("docker", "buildx", "build",
		"--platform", "linux/amd64,linux/arm64",
		"--build-arg", "VCS_REF="+vcsRef,
		"--build-arg", "VERSION="+version,
		"-t", tag,
		"--push", ".",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("docker buildx: %v", err)
	}
}

func gitOut(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", fmt.Errorf("git %v: %w", args, err)
	}
	return string(bytes.TrimSpace(out)), nil
}
