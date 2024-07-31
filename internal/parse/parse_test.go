package parse

import (
	"bytes"
	"github.com/pymeta/go-diff/internal/diff"
	"golang.org/x/tools/txtar"
	"path/filepath"
	"testing"
)

func clean(text []byte) []byte {
	text = bytes.ReplaceAll(text, []byte("$\n"), []byte("\n"))
	text = bytes.TrimSuffix(text, []byte("^D\n"))
	return text
}

func Test(t *testing.T) {
	files, _ := filepath.Glob("testdata/*.txt")
	if len(files) == 0 {
		t.Fatalf("no testdata")
	}

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			a, err := txtar.ParseFile(file)
			if err != nil {
				t.Fatal(err)
			}
			if len(a.Files) != 3 || a.Files[2].Name != "diff" {
				t.Fatalf("%s: want three files, third named \"diff\"", file)
			}
			edits, err := ParseEdits(string(clean(a.Files[0].Data)), string(clean(a.Files[2].Data)))
			if err != nil {
				t.Fatal(err)
			}
			out, err := diff.ApplyBytes(clean(a.Files[0].Data), edits)
			if err != nil {
				t.Fatal(err)
			}
			want := clean(a.Files[1].Data)
			if !bytes.Equal(out, want) {
				t.Fatalf("%s, %s:\nhave:\n%s\nwant:\n%s", file, edits, out, want)
			}
		})
	}
}
