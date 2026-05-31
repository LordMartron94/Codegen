package gocode

import (
	"io"

	"codegen"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

/*
FileStreamConfig controls when buffered elements are rendered and written to disk.
*/
type FileStreamConfig struct {
	FlushEveryElements int
	FlushEveryBytes    int
}

/*
FileStream renders Go file elements incrementally to an io.Writer with periodic flushes.
*/
type FileStream struct {
	writer  io.Writer
	eng     *engine.Engine
	options goast.RenderOptions
	config  FileStreamConfig
	emitter emit.Emitter
	pending []codegen.FileElement
}

/*
FileStreamCreate constructs a stream that writes rendered Go source to writer.
*/
func FileStreamCreate(writer io.Writer, options goast.RenderOptions, config FileStreamConfig) *FileStream {
	return &FileStream{
		writer:  writer,
		eng:     GoEngineCreate(),
		options: options,
		config:  config,
		emitter: GoEmitterCreate(),
		pending: make([]codegen.FileElement, 0, 64),
	}
}

/*
FileStreamWriteElements appends top-level file elements and flushes when thresholds are met.
*/
func FileStreamWriteElements(stream *FileStream, elements ...codegen.FileElement) error {
	if stream == nil {
		return nil
	}
	stream.pending = append(stream.pending, elements...)
	return fileStreamFlushIfNeeded(stream)
}

/*
FileStreamFlush renders pending elements to the writer and clears the emitter buffer.
*/
func FileStreamFlush(stream *FileStream) error {
	if stream == nil || len(stream.pending) == 0 {
		return nil
	}

	ctx := &engine.RenderContext{
		Engine:  stream.eng,
		Emitter: &stream.emitter,
		Data:    stream.options,
	}
	for _, element := range stream.pending {
		if err := engine.EngineRenderNode(ctx, element.Node); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(stream.writer, emit.EmitterRender(&stream.emitter)); err != nil {
		return err
	}
	emit.EmitterReset(&stream.emitter)
	stream.pending = stream.pending[:0]
	return nil
}

/*
FileStreamClose flushes any remaining pending elements.
*/
func FileStreamClose(stream *FileStream) error {
	return FileStreamFlush(stream)
}

func fileStreamFlushIfNeeded(stream *FileStream) error {
	if stream.config.FlushEveryElements > 0 && len(stream.pending) >= stream.config.FlushEveryElements {
		return FileStreamFlush(stream)
	}
	if stream.config.FlushEveryBytes > 0 && emit.EmitterBufferedLen(&stream.emitter) >= stream.config.FlushEveryBytes {
		return FileStreamFlush(stream)
	}
	return nil
}
