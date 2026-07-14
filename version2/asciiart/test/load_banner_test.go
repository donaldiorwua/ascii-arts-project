package asciiart

import (
    "os"
    "testing"
)

// TestLoadBanner_ReturnsEnoughLines checks that the banner files contain at least the minimum expected lines.
func TestLoadBanner_ReturnsEnoughLines(t *testing.T) {
    lines, err := LoadBanner("standard.txt")
    if err != nil {
        t.Fatalf("unexpected error loading standard.txt: %v", err)
    }
    if len(lines) < 856 {
        t.Errorf("expected at least 856 lines in banner, got %d", len(lines))
    }
}

// TestLoadBanner_ShadowAndThinkertoy verifies the other banner files also load with sufficient lines.
func TestLoadBanner_ShadowAndThinkertoy(t *testing.T) {
    for _, filename := range []string{"shadow.txt", "thinkertoy.txt"} {
        lines, err := LoadBanner(filename)
        if err != nil {
            t.Errorf("%s: unexpected error: %v", filename, err)
            continue
        }
        if len(lines) < 856 {
            t.Errorf("%s: expected at least 856 lines, got %d", filename, len(lines))
        }
    }
}

// TestLoadBanner_FileNotFound expects a non-nil error when the file does not exist.
func TestLoadBanner_FileNotFound(t *testing.T) {
    _, err := LoadBanner("notfound.txt")
    if err == nil {
        t.Error("expected a non-nil error for missing file, got nil")
    }
}

// TestLoadBanner_EmptyFile expects an error when the file exists but is empty.
func TestLoadBanner_EmptyFile(t *testing.T) {
    f, err := os.CreateTemp("", "empty_banner_*.txt")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(f.Name())
    f.Close()

    _, err = LoadBanner(f.Name())
    if err == nil {
        t.Error("expected error for empty file, got nil")
    }
}

// TestLoadBanner_IncompleteFile expects an error for files with too few lines.
func TestLoadBanner_IncompleteFile(t *testing.T) {
    f, err := os.CreateTemp("", "incomplete_banner_*.txt")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(f.Name())

    _, _ = f.WriteString("only a couple of lines\nnot nearly enough\n")
    f.Close()

    _, err = LoadBanner(f.Name())
    if err == nil {
        t.Error("expected error for incomplete banner content, got nil")
    }
}

// TestLoadBanner_NilFilename expects an error when an empty filename is provided.
func TestLoadBanner_NilFilename(t *testing.T) {
    _, err := LoadBanner("")
    if err == nil {
        t.Error("expected error for empty filename, got nil")
    }
}