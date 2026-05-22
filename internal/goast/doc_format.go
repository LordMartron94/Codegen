package goast

import (
	"strings"
	"unicode"
)

/*
GoDocFormatExported rewrites documentation so the first paragraph begins with the exported identifier.

[Parameters]
subject is the Go name being documented (type, const, or field). doc is raw documentation text from an external spec.

[Returns]
Formatted documentation suitable for GoDoc and gopls, or an empty string when doc is empty.

[Context]
Go requires that the comment on an exported declaration start with the name of the item. Multi-paragraph text only adjusts the first paragraph when it does not already start with subject.
*/
func GoDocFormatExported(subject string, doc string) string {
	subject = strings.TrimSpace(subject)
	doc = strings.TrimSpace(doc)
	if doc == "" {
		return ""
	}
	if subject == "" {
		return doc
	}
	if goDocStartsWithSubject(subject, doc) {
		return doc
	}

	paragraphs := strings.Split(doc, "\n\n")
	if len(paragraphs) == 0 {
		return subject + " " + doc
	}

	first := strings.TrimSpace(paragraphs[0])
	if !goDocStartsWithSubject(subject, first) {
		paragraphs[0] = subject + " " + first
	}

	return strings.Join(paragraphs, "\n\n")
}

func goDocStartsWithSubject(subject string, text string) bool {
	text = strings.TrimSpace(text)
	if text == "" {
		return false
	}
	if strings.HasPrefix(text, subject) {
		remainder := text[len(subject):]
		return remainder == "" || !unicode.IsLetter(rune(remainder[0]))
	}
	return false
}
