package diff

import (
	"github.com/pymeta/go-diff/internal/diff"
	"github.com/pymeta/go-diff/internal/diffp"
	"github.com/pymeta/go-diff/internal/parse"
)

type Edit = diff.Edit

var Apply = diff.Apply
var ApplyBytes = diff.ApplyBytes
var SortEdits = diff.SortEdits
var Strings = diff.Strings
var Bytes = diff.Bytes

const DefaultContextLines = diff.DefaultContextLines

var Unified = diff.Unified
var ToUnified = diff.ToUnified
var Diff = diffp.Diff

var ParseEdits = parse.ParseEdits
