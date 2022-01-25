package actions

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/kyoto-framework/kyoto"
)

// Flush is a function to immediately update a component layout during action.
// It is useful for updating a component layout multiple times during long action.
func Flush(b *kyoto.Core) {
	// Extract context
	rw := b.Context.GetResponseWriter()
	rwf := rw.(http.Flusher)
	// Render
	buffer := bytes.NewBufferString("")
	var err error
	if b.Context.Get("internal:render:cm") != nil {
		renderer := b.Context.Get("internal:render:cm").(func(io.Writer) error)
		err = renderer(buffer)
	} else {
		tbuilder := b.Context.Get("internal:render:tb").(func() *template.Template)
		err = tbuilder().Execute(buffer, b.State.Export())
	}
	if err != nil {
		panic(err)
	}
	html := buffer.String()
	// Remove newlines (not supported by SSA)
	html = strings.ReplaceAll(html, "\n", "")
	// Write SSE
	_, err = fmt.Fprintf(rw, "data: %v\n\n", html)
	if err != nil {
		panic(err)
	}
	// Flush
	rwf.Flush()
}
