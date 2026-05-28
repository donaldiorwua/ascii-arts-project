package asciiartsproject

import (
	"fmt"
	"strings"
)

func Start() {
	anim := NewAnimation("HELLO", 8)

	anim.GenerateSpinFrames()
	fmt.Println(anim.Play())
}

type Animation struct {
	text   string
	frames int
	data   []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text:   text,
		frames: frames,
		data:   make([]string, frames),
	}
}

func (a *Animation) GenerateSpinFrames() {
	spineers := []string{"|", "/", "-", "\\"}

	for i := 0; i < a.frames; i++ {
		symbol := spineers[i%len(spineers)]

		lines := make([]string, 10)

		for j := 0; j < 10; j++ {
			padding := strings.Repeat(" ", (i+j)%5)
			content := fmt.Sprintf("%s %s %s", symbol, a.text, symbol)

			lines[j] = padLine(padding + content)
		}

		a.data[i] = strings.Join(lines, "\n")
	}
}

func (a *Animation) GenerateWaveFrames() {
	for i := 0; i < a.frames; i++ {
		lines := make([]string, 10)

		for j := 0; j < 10; j++ {
			offset := (i + j) % 6
			padding := strings.Repeat(" ", offset)

			lines[j] = padLine(padding + a.text)
		}
		a.data[i] = strings.Join(lines, "\n")
	}
}


func (a *Animation) GenerateZoomFrames() {
	for i := 0; i < a.frames; i++ {
		lines := make([]string, 10)
		scale := i + 1

		if i > a.frames/2 {
			scale = a.frames - 1
		}
		zoomed := zoomText(a.text, scale)

		for j := 0; j < 10; j++ {
			padding := strings.Repeat(" ", j%3)
			lines[j] = padLine(padding + zoomed)
		}

		a.data[i] = strings.Join(lines, "\n")
	}
}

func (a *Animation) GetFrame(index int) string {
	if len(a.data) == 0 {
		return ""
	}

	index = index % len(a.data)

	return a.data[index]
}

func (a *Animation) Play() string {
	var b strings.Builder

	for i := 0; i < len(a.data); i++ {
		b.WriteString(fmt.Sprintf("=== Frame %d ===\n", i))
		b.WriteString(a.data[i])
		b.WriteString("\n\n")
	}

	return b.String()
}

func zoomText(text string, scale int) string {
	var b strings.Builder

	for _, r := range text {
		b.WriteString(strings.Repeat(string(r), scale))
	}

	return b.String()
}

func padLine(s string) string {
	const width = 40

	if len(s) < width {
		s += strings.Repeat(" ", width-len(s))
	}

	return s
}
