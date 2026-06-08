# ASCII ARTS PROJECT

## Overview
This is a small Go utility that renders ASCII art text from banner files. It reads a banner file (a set of multi-line character glyphs) and maps printable ASCII characters to 8-line art blocks, then uses that map to generate rendered ASCII output for input text.

## Features
- Load a banner file and map printable ASCII runes to 8-line glyphs using `LoadBanner`function  in loadfile.go.
- Render input strings (supports explicit newlines) with `GenerateArt` function in generate.go file.
- A validator for banner maps is available as `Validate` function in `validate.go` file.
- A helper to split provided input is implemented as `SplitInput` function in splitInput.go file.
- Example banner files included: [standard.txt], [shadow.txt], [thinkertoy.txt]


## Banner file format
- Each printable ASCII character (from rune 32 to 126 inclusive) must have an 8-line glyph block.
- The banner file is expected to be organized as 95 blocks, each block occupying 8 lines plus a separator/newline (9 lines per character). The loader expects at least 855–856 lines (95 * 9).
- Use the included banner files as templates when creating custom banners.

## Usage
Build and run with a single quoted argument. The program uses `standard.txt` by default:
The program accepts multi-line input by including explicit newline characters inside the quoted string:

### Run without build:
    

#### Render a single word:
 - bash
go run . "ASCII"


#### Render text with explicit newline:
 - bash
go run . "Line1\nLine2"

#### Note:
Use provided banner files (default is the standard banner shipped with the repo). Swap banners by replacing the default banner file with another from the banners directory if desired.

## Project layout (excluding any "function" folder)

ascii-arts-project/
- go.mod
- main.go
- loadfile.go
- generate.go
- validate.go
- splitInput.go
- banners:
    - standard.txt
    - shadow.txt
    - thinkertoy.txt
- README.md

## Development notes

- Keep banner files in sync with the 95-glyph expectation; an incomplete banner will cause validation to fail.
- The codebase separates responsibilities: banner loading, input splitting, validation, and art generation. This makes it straightforward to replace or extend any step.
- When adding new banners, follow the format precisely: 8 lines per glyph in the same sequence used by the loader.
