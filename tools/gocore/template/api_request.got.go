// Code generated by hero.
// source: /Users/SM0286/code/core/gocore/tools/gocore/template/api_request.got
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func FromApiRequest(request, params string, buffer *bytes.Buffer) {
	buffer.WriteString(`

type `)
	hero.EscapeHTML(request, buffer)
	buffer.WriteString(` struct {
    `)
	hero.EscapeHTML(params, buffer)
	buffer.WriteString(`
}`)

}