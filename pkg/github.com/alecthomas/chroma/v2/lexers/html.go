package lexers

import (
	"github.com/go-task/task/v3/pkg/github.com/alecthomas/chroma/v2"
)

// HTML lexer.
var HTML = chroma.MustNewXMLLexer(embedded, "embedded/html.xml")
