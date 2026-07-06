package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	current, err := miseGet("env.APP_NAME")
	if err != nil {
		log.Fatalf("read APP_NAME: %v", err)
	}
	if current != "" {
		log.Printf("APP_NAME already set: %s", current)
		return
	}

	fmt.Print("Enter the app name: ")
	name, err := readLine(os.Stdin)
	if err != nil {
		log.Fatalf("read input: %v", err)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		log.Fatal("app name cannot be empty")
	}
	if err := miseSet("env.APP_NAME", name); err != nil {
		log.Fatalf("set APP_NAME: %v", err)
	}
	log.Printf("APP_NAME set to: %s", name)
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
