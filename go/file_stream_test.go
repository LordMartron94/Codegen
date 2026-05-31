package gocode_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"codegen"
	gocode "codegen/go"
)

func TestFileStreamMatchesBatchRender(t *testing.T) {
	const count = 120
	elements := make([]codegen.FileElement, 0, count+4)
	elements = append(elements, gocode.GoBuildConstraint("linux"))
	elements = append(elements, gocode.GoBlankLine())
	elements = append(elements, gocode.GoGeneratedFileHeader("test tool", time.Unix(0, 0))...)
	elements = append(elements, gocode.FileElementFrom(gocode.DeclPackage("bindings")))
	elements = append(elements, gocode.GoBlankLine())
	for i := 0; i < count; i++ {
		elements = append(elements, gocode.GoLineComment("line"))
	}

	batchFile := gocode.DeclFile(elements...)
	batch, err := gocode.GoFileRenderWithOptions(batchFile, gocode.RenderOptionsDefault())
	if err != nil {
		t.Fatalf("batch render: %v", err)
	}

	var streamed bytes.Buffer
	stream := gocode.FileStreamCreate(&streamed, gocode.RenderOptionsDefault(), gocode.FileStreamConfig{
		FlushEveryElements: 10,
	})
	if err := gocode.FileStreamWriteElements(stream, elements...); err != nil {
		t.Fatalf("stream write: %v", err)
	}
	if err := gocode.FileStreamClose(stream); err != nil {
		t.Fatalf("stream close: %v", err)
	}

	if strings.TrimSpace(batch) != strings.TrimSpace(streamed.String()) {
		t.Fatalf("stream output differs from batch render")
	}
}
