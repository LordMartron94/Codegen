package emit

import (
	"fmt"
	"strings"
)

/*
Emitter is a language-neutral buffered text sink with indentation support.

[Context]
Renderers write through Emitter free functions; engines inject an Emitter instance per render pass.
*/
type Emitter struct {
	buffer *strings.Builder

	indentWidth int
	indentLevel int
	useTabs     bool

	pendingIndent bool
}

/*
EmitterCreate constructs an emitter with the given indent width in spaces.

[Parameters]
indentWidth — number of spaces per indent level; Go output typically uses 4.

[Returns]
A zeroed Emitter ready for writes.
*/
func EmitterCreate(indentWidth int) Emitter {
	return Emitter{
		buffer:      &strings.Builder{},
		indentWidth: indentWidth,
	}
}

/*
EmitterCreateTabs constructs an emitter that indents with tab characters per level.

[Context]
Go output uses tab indentation per gofmt; one tab per indent level.
*/
func EmitterCreateTabs() Emitter {
	return Emitter{
		buffer:  &strings.Builder{},
		useTabs: true,
	}
}

func EmitterIndent(emitter *Emitter) {
	emitter.indentLevel++
}

func EmitterDedent(emitter *Emitter) {
	if emitter.indentLevel == 0 {
		return
	}
	emitter.indentLevel--
}

func EmitterLineBreak(emitter *Emitter) {
	emitter.buffer.WriteRune('\n')
	emitter.pendingIndent = true
}

func EmitterWrite(emitter *Emitter, content string) {
	emitterHandleIndent(emitter)
	emitter.buffer.WriteString(content)
}

func EmitterWriteFormatted(emitter *Emitter, format string, args ...any) {
	emitterHandleIndent(emitter)
	fmt.Fprintf(emitter.buffer, format, args...)
}

func EmitterWriteLine(emitter *Emitter, line string) {
	EmitterWrite(emitter, line)
	EmitterLineBreak(emitter)
}

func EmitterWriteLineFormatted(emitter *Emitter, format string, args ...any) {
	EmitterWriteFormatted(emitter, format, args...)
	EmitterLineBreak(emitter)
}

/*
EmitterRender returns the accumulated output string.

[Returns]
All content written to the emitter since creation.
*/
func EmitterRender(emitter *Emitter) string {
	return emitter.buffer.String()
}

/*
EmitterBufferedLen returns the number of bytes accumulated in the emitter buffer.
*/
func EmitterBufferedLen(emitter *Emitter) int {
	return emitter.buffer.Len()
}

/*
EmitterReset clears buffered output and indentation state for reuse after a flush.
*/
func EmitterReset(emitter *Emitter) {
	emitter.buffer.Reset()
	emitter.indentLevel = 0
	emitter.pendingIndent = false
}

func emitterHandleIndent(emitter *Emitter) {
	if !emitter.pendingIndent {
		return
	}

	if emitter.useTabs {
		for i := 0; i < emitter.indentLevel; i++ {
			emitter.buffer.WriteRune('\t')
		}
	} else {
		numberOfSpaces := emitter.indentLevel * emitter.indentWidth
		for i := 0; i < numberOfSpaces; i++ {
			emitter.buffer.WriteRune(' ')
		}
	}

	emitter.pendingIndent = false
}
