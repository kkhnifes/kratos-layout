package mise

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Get(key string) (string, error) {
	out, err := exec.Command("mise", "config", "get", key).Output()
	if err != nil {
		return "", fmt.Errorf("mise config get %s: %w", key, err)
	}
	return strings.TrimSpace(string(out)), nil
}

func Set(key, value string) error {
	arg := fmt.Sprintf("%s=%s", key, value)
	cmd := exec.Command("mise", "config", "set", arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mise config set %s: %w", arg, err)
	}
	return nil
}

func Prompt(label, current string) (value string, changed bool, err error) {
	return promptFrom(label, current, os.Stdin, os.Stdout)
}

func promptFrom(label, current string, r io.Reader, w io.Writer) (string, bool, error) {
	if current != "" {
		fmt.Fprintf(w, "%s [current: %s]: ", label, current)
	} else {
		fmt.Fprintf(w, "%s: ", label)
	}
	line, err := bufio.NewReader(r).ReadString('\n')
	if err != nil && line == "" {
		return "", false, fmt.Errorf("read input: %w", err)
	}
	input := strings.TrimSpace(line)
	if input == "" {
		return current, false, nil
	}
	return input, true, nil
}
