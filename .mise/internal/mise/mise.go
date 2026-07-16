package mise

import (
	"bufio"
	"fmt"
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
	if current != "" {
		fmt.Printf("%s [current: %s]: ", label, current)
	} else {
		fmt.Printf("%s: ", label)
	}
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && line == "" {
		return "", false, fmt.Errorf("read input: %w", err)
	}
	input := strings.TrimSpace(line)
	if input == "" {
		return current, false, nil
	}
	return input, true, nil
}
