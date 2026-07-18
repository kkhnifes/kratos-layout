package mise

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestPromptKeep(t *testing.T) {
	var out bytes.Buffer
	v, changed, err := promptFrom("Enter app name", "kratos-layout", strings.NewReader("\n"), &out)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if changed {
		t.Fatal("empty input must keep current, not change")
	}
	if v != "kratos-layout" {
		t.Fatalf("v=%q want kratos-layout", v)
	}
	if got := out.String(); got != "Enter app name [current: kratos-layout]: " {
		t.Fatalf("prompt line=%q", got)
	}
}

func TestPromptOverride(t *testing.T) {
	var out bytes.Buffer
	v, changed, err := promptFrom("New path name", "kkhnifes", strings.NewReader("acme\n"), &out)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !changed {
		t.Fatal("non-empty input must set changed=true")
	}
	if v != "acme" {
		t.Fatalf("v=%q want acme", v)
	}
	if got := out.String(); got != "New path name [current: kkhnifes]: " {
		t.Fatalf("prompt line=%q", got)
	}
}

func TestPromptTrim(t *testing.T) {
	var out bytes.Buffer
	v, changed, err := promptFrom("Enter", "", strings.NewReader("  spaced  \r\n"), &out)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !changed || v != "spaced" {
		t.Fatalf("v=%q changed=%v want spaced/true", v, changed)
	}
	if got := out.String(); got != "Enter: " {
		t.Fatalf("no-current prompt=%q want %q", got, "Enter: ")
	}
}

func TestPromptEOF(t *testing.T) {
	var out bytes.Buffer
	_, _, err := promptFrom("Enter app name", "", errReader{}, &out)
	if err == nil {
		t.Fatal("want read error on immediate EOF")
	}
	if !errors.Is(err, errSentinel) {
		t.Fatalf("want wrapped sentinel, got %v", err)
	}
}

type errReader struct{}

var errSentinel = errors.New("boom")

func (errReader) Read(p []byte) (int, error) { return 0, errSentinel }
