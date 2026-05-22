package goast

/*
CommentStyle selects how documentation text is rendered in Go output.
*/
type CommentStyle uint8

const (
	/* CommentStyleLine renders documentation as // line comments. */
	CommentStyleLine CommentStyle = iota
	/* CommentStyleBlock renders documentation as a block comment (slash-star form). */
	CommentStyleBlock
)

/*
RenderOptions configures Go backend rendering behavior.

[Context]
Attach to engine.RenderContext.Data when rendering. Unset fields use RenderOptionsDefault values.
*/
type RenderOptions struct {
	DocCommentStyle CommentStyle
}

/*
RenderOptionsDefault returns Go render defaults (line documentation comments).
*/
func RenderOptionsDefault() RenderOptions {
	return RenderOptions{DocCommentStyle: CommentStyleLine}
}

/*
RenderOptionsEnumBindings returns render settings suited to generated Vulkan enum binding files.
*/
func RenderOptionsEnumBindings() RenderOptions {
	return RenderOptions{DocCommentStyle: CommentStyleBlock}
}
