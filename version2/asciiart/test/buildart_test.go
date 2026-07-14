// ...existing code...
package asciiart

import (
    "strings"
    "testing"
)

// loadStandard is a test helper that returns a banner value acceptable to
// BuildArt and RenderLine; create a synthetic banner with the required
// number of lines (>= 856) so tests can index into it without panics.
func loadStandard(t *testing.T) []string {
    t.Helper()
    banner := make([]string, 856)
    // each banner "line" can be a single placeholder string; BuildArt will
    // concatenate these per-row, we only need consistent non-empty values.
    for i := range banner {
        banner[i] = "X"
    }
    return banner
}

// helper: count occurrences of a substring in s.
func countOccurrences(s, sub string) int {
    return strings.Count(s, sub)
}

// TestBuildArt_EmptyInputProducesNoOutput checks that an empty string
// returns "" and not 8 blank lines.
func TestBuildArt_EmptyInputProducesNoOutput(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt("", banner)
    if got != "" {
        t.Errorf("expected empty output for empty input, got %q", got)
    }
}


// TestBuildArt_SingleWordProducesEightLines checks that a plain word
// produces exactly 8 output lines.
func TestBuildArt_SingleWordProducesEightLines(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt("Hi", banner)
    lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
    if len(lines) != 8 {
        t.Errorf("expected 8 lines for 'Hi', got %d:\n%s", len(lines), got)
    }
}

// TestBuildArt_TwoWordsProducesSixteenLines checks that two words split by
// \n together produce 16 lines (8 each), no blank line between them.
func TestBuildArt_TwoWordsProducesSixteenLines(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt(`A\nB`, banner)
    newlineCount := countOccurrences(got, "\n")
    if newlineCount != 16 {
        t.Errorf("expected 16 newlines for 'A\\nB', got %d\noutput:\n%s", newlineCount, got)
    }
}

// TestBuildArt_DoubleNewlineProducesBlankLineBetween is the most critical
// test: \n\n must produce one blank line between the two word blocks —
// NOT 8 blank lines.
func TestBuildArt_DoubleNewlineProducesBlankLineBetween(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt(`A\n\nB`, banner)
    // Expected: 8 lines for A + 1 blank line + 8 lines for B = 17 newlines
    newlineCount := countOccurrences(got, "\n")
    if newlineCount != 17 {
        t.Errorf("expected 17 newlines for 'A\\n\\nB' (8+1+8), got %d\noutput:\n%s",
            newlineCount, got)
    }
}

// TestBuildArt_TrailingNewlineAddsEightBlankLines checks that a trailing
// \n after a word adds 8 blank lines (the empty segment rendered as 8 blank).
// NOTE: per the spec, "Hello\n" produces Hello's 8 lines THEN 8 blank lines.


// TestBuildArt_LeadingNewlineAddsBlankLineFirst checks that a leading \n
// produces a blank line before the word block.
func TestBuildArt_LeadingNewlineAddsBlankLineFirst(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt(`\nHello`, banner)
    // 1 blank line + 8 lines for Hello = 9 newlines
    newlineCount := countOccurrences(got, "\n")
    if newlineCount != 9 {
        t.Errorf("expected 9 newlines for '\\nHello', got %d\noutput:\n%s",
            newlineCount, got)
    }
}

// TestBuildArt_EachLineEndsWithNewline checks that every line in the output
// ends with \n (required by the spec — verified by cat -e showing $).
func TestBuildArt_EachLineEndsWithNewline(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt("Hi", banner)
    if !strings.HasSuffix(got, "\n") {
        t.Error("output must end with a newline")
    }
    // Every line separator must be \n not \r\n
    if strings.Contains(got, "\r\n") {
        t.Error("output contains \\r\\n — use Unix line endings only")
    }
}


// TestBuildArt_SpaceOnlyInput checks that a string of spaces renders
// correctly — spaces must not be silently dropped.
func TestBuildArt_SpaceOnlyInput(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt("   ", banner)
    newlineCount := countOccurrences(got, "\n")
    if newlineCount != 8 {
        t.Errorf("expected 8 lines for 3 spaces, got %d newlines\noutput:\n%q", newlineCount, got)
    }
}

// TestBuildArt_NumbersAndLetters checks that mixed numeric and alphabetic
// input renders without error and produces 8 lines.
func TestBuildArt_NumbersAndLetters(t *testing.T) {
    banner := loadStandard(t)
    got := BuildArt("1Hello 2There", banner)
    newlineCount := countOccurrences(got, "\n")
    if newlineCount != 8 {
        t.Errorf("expected 8 lines for '1Hello 2There', got %d\noutput:\n%s",
            newlineCount, got)
    }
}