package asciiart

import "testing"

// TestValidateInput_ValidStrings checks that fully valid inputs return the same string.
func TestValidateInput_ValidStrings(t *testing.T) {
    valid := []string{
        "Hello",
        "hello world",
        "Hello, World!",
        "1234567890",
        "~`!@#$%^&*()-_=+[]{}|;':\",./<>?",
        " ",
    }
    for _, s := range valid {
        got := ValidateInput(s)
        if got != s {
            t.Errorf("ValidateInput(%q): expected %q, got %q", s, s, got)
        }
    }
}

// TestValidateInput_InvalidCharacters checks that inputs containing invalid runes return empty string.
func TestValidateInput_InvalidCharacters(t *testing.T) {
    cases := []string{
        "héllo",
        "naïve",
        "café",
        "こんにちは",
        "\x01hello",
        "\x00",
        "\x1F",      // ASCII 31 — just below valid range
        "hello\x7F", // ASCII 127 — just above valid range
        "😀",
    }
    for _, input := range cases {
        got := ValidateInput(input)
        if got != "" {
            t.Errorf("ValidateInput(%q): expected empty string for invalid input, got %q", input, got)
        }
    }
}

// TestValidateInput_ReturnsEmptyOnMultipleInvalid ensures strings with multiple invalids also return empty.
func TestValidateInput_ReturnsEmptyOnMultipleInvalid(t *testing.T) {
    got := ValidateInput("héñ")
    if got != "" {
        t.Errorf("ValidateInput(%q): expected empty string, got %q", "héñ", got)
    }
}

// TestValidateInput_BoundaryASCII32And126 checks that ASCII codes 32 (space) and 126 (~) are accepted.
func TestValidateInput_BoundaryASCII32And126(t *testing.T) {
    for _, code := range []rune{32, 126} {
        s := string(code)
        got := ValidateInput(s)
        if got != s {
            t.Errorf("ASCII %d (%q) should be valid, got %q", code, code, got)
        }
    }
}

// TestValidateInput_BoundaryASCII31And127 checks that ASCII codes 31 and 127 are rejected.
func TestValidateInput_BoundaryASCII31And127(t *testing.T) {
    for _, code := range []rune{31, 127} {
        s := string(code)
        got := ValidateInput(s)
        if got != "" {
            t.Errorf("ASCII %d should be invalid, expected empty string, got %q", code, got)
        }
    }
}

// TestValidateInput_ValidPrefixBeforeInvalid checks that a valid prefix with invalid char at end returns empty.
func TestValidateInput_ValidPrefixBeforeInvalid(t *testing.T) {
    got := ValidateInput("Hello World é")
    if got != "" {
        t.Errorf("expected empty string for %q, got %q", "Hello World é", got)
    }
}

// TestValidateInput_EmptyString expects empty string for empty input.
func TestValidateInput_EmptyString(t *testing.T) {
    got := ValidateInput("")
    if got != "" {
        t.Errorf("ValidateInput(\"\"): expected empty string, got %q", got)
    }
}

// TestValidateInput_EscapeNewline checks the special-case input "\\n" returns empty string.
func TestValidateInput_EscapeNewline(t *testing.T) {
    got := ValidateInput("\\n")
    if got != "" {
        t.Errorf("ValidateInput(\"\\\\n\"): expected empty string, got %q", got)
    }
}
// ...existing code...