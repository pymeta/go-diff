package parse

import (
	"errors"
	"github.com/pymeta/go-diff/internal/diff"
	"regexp"
	"strconv"
	"strings"
)

func indexN(s string, char rune, n int) int {
	if n <= 0 {
		return -1
	}

	count := 0
	for i, c := range s {
		if c == char {
			count++
			if count == n {
				return i
			}
		}
	}
	return -1
}

const (
	deleteOp     = '-'
	insertOp     = '+'
	equalOp      = ' '
	patchBeginOp = '@'
	eofOp        = '\\'
)

// ParseEdits parses a textual representation of patches and returns a List of Edits
//
//goland:noinspection GoNameStartsWithPackageName
func ParseEdits(input string, patch string) ([]diff.Edit, error) {
	edits := make([]diff.Edit, 0)
	if len(patch) == 0 {
		return edits, nil
	}
	patch = strings.TrimSpace(patch)
	patchLines := strings.Split(patch, "\n")
	patchHeader := regexp.MustCompile("^@@ -(\\d+),?(\\d*) \\+(\\d+),?(\\d*) @@$")

	var linePointer = 0
	var patchStartPos = -1
	var prevSign = uint8(' ')

	for linePointer < len(patchLines) {
		if !patchHeader.MatchString(patchLines[linePointer]) {
			linePointer++
			continue
		}

		m := patchHeader.FindStringSubmatch(patchLines[linePointer])

		patchStartLine, _ := strconv.Atoi(m[1])
		if len(m[2]) == 0 {
			patchStartLine--
		} else if m[2] == "0" {
		} else {
			patchStartLine--
		}

		linePointer++
		patchStartPos = indexN(input, '\n', patchStartLine) + 1
		patchStopPos := patchStartPos
		patchText := ""
		var sign uint8

		for linePointer < len(patchLines) {
			if len(patchLines[linePointer]) > 0 {
				sign = patchLines[linePointer][0]
			} else {
				linePointer++
				patchStopPos++
				patchText += "\n"
				continue
			}
			line := patchLines[linePointer][1:]

			if sign == insertOp {
				patchText += line + "\n"
			} else if sign == eofOp {
				if prevSign == insertOp {
					patchStopPos++
					patchText = patchText[:len(patchText)-1]
				}
			} else if sign == patchBeginOp {
				break
			} else if sign == equalOp {
				patchText += line + "\n"
				patchStopPos += len(line) + 1
			} else if sign == deleteOp {
				patchStopPos += len(line) + 1
			} else {
				return nil, errors.New("undefined operator: %s" + string(sign))
			}
			linePointer++
			prevSign = sign
		}
		edits = append(edits, diff.Edit{
			Start: max(0, patchStartPos),
			End:   min(max(0, patchStartPos, patchStopPos), len(input)),
			New:   patchText,
		})
	}
	return edits, nil
}
