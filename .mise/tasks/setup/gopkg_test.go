package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadGoModule(t *testing.T) {
	dir := t.TempDir()
	mod := filepath.Join(dir, "go.mod")
	if err := os.WriteFile(mod, []byte("module github.com/old/x\n\ngo 1.26\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := readGoModule(mod)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if got != "github.com/old/x" {
		t.Fatalf("got=%q", got)
	}
}

func TestReadGoModuleMissing(t *testing.T) {
	dir := t.TempDir()
	_, err := readGoModule(filepath.Join(dir, "go.mod"))
	if err == nil {
		t.Fatal("want error when module line absent")
	}
}

func TestReplaceOnce(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "go.mod")
	if err := os.WriteFile(path, []byte("module github.com/old/x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := replaceOnce(path, "module github.com/old/x", "module github.com/new/y"); err != nil {
		t.Fatalf("err: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "module github.com/new/y\n" {
		t.Fatalf("got=%q", got)
	}
}

func TestReplaceOnceMissing(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "go.mod")
	if err := os.WriteFile(path, []byte("module x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := replaceOnce(path, "module z", "module y"); err == nil {
		t.Fatal("want error when old string absent")
	}
}

func TestReplaceInTree(t *testing.T) {
	root := t.TempDir()
	inner := filepath.Join(root, "pkg", "svc")
	if err := os.MkdirAll(inner, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "pkg", "svc", "a.go"), []byte("module github.com/old/x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "b.go"), []byte("module github.com/old/x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "c.go"), []byte("no match"), 0o644); err != nil {
		t.Fatal(err)
	}
	n, err := replaceInTree(root, "github.com/old/x", "github.com/new/y")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if n != 2 {
		t.Fatalf("n=%d want 2", n)
	}
	a, err := os.ReadFile(filepath.Join(root, "pkg", "svc", "a.go"))
	if err != nil {
		t.Fatal(err)
	}
	if string(a) != "module github.com/new/y" {
		t.Fatalf("a.go=%q", a)
	}
}
