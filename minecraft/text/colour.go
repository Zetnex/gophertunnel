package text

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

// ANSI converts all Minecraft text formatting codes in the values passed to ANSI formatting codes, so that
// it may be displayed properly in the terminal.
func ANSI(a ...interface{}) string {
	str := make([]string, len(a))
	for i, v := range a {
		str[i] = minecraftReplacer.Replace(fmt.Sprint(v))
	}
	return strings.Join(str, " ")
}

// Colourf colours the format string using HTML tags after first escaping all parameters passed and
// substituting them in the format string. The following colours and formatting may be used:
// 	black, dark-blue, dark-green, dark-aqua, dark-red, dark-purple, gold, grey, dark-grey, blue, green, aqua,
// 	red, purple, yellow, white, dark-yellow, obfuscated, bold, strikethrough, underline, and italic.
// These HTML tags may also be nested, like so:
// `<red>Hello <bold>World</bold>!</red>`
func Colourf(format string, a ...interface{}) string {
	params := make([]interface{}, len(a))
	for i := 0; i < len(a); i++ {
		params[i] = html.EscapeString(fmt.Sprintf("%+v", a[i]))
	}
	str := fmt.Sprintf(format, params...)

	e := &enc{w: &strings.Builder{}}
	t := html.NewTokenizer(strings.NewReader(str))
	for {
		if t.Next() == html.ErrorToken {
			break
		}
		e.process(t.Token())
	}
	return e.w.String()
}

// enc holds the state of a string to be processed for colour substitution.
type enc struct {
	w           *strings.Builder
	formatStack []string
}

// process handles a single html.Token and either writes the string to the strings.Builder, adds a colour to
// the stack or removes a colour from the stack.
func (e *enc) process(tok html.Token) {
	switch tok.Type {
	case html.TextToken:
		for _, s := range e.formatStack {
			e.w.WriteString(s)
		}
		e.w.WriteString(tok.Data)
		if len(e.formatStack) != 0 {
			e.w.WriteString(reset)
		}
	case html.StartTagToken:
		if format, ok := strMap[tok.Data]; ok {
			e.formatStack = append(e.formatStack, format)
		}
	case html.EndTagToken:
		for i, format := range e.formatStack {
			if f, ok := strMap[tok.Data]; ok && f == format {
				e.formatStack = append(e.formatStack[:i], e.formatStack[i+1:]...)
				break
			}
		}
	}
}
