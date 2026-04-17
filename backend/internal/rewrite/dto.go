package rewrite

type RewriteRequest struct {
	Text  string `json:"text" binding:"required"`
	Style string `json:"style" binding:"required,oneof=official humorous friendly professional"`
}

type RewriteResponse struct {
	OriginalText  string `json:"original_text"`
	RewrittenText string `json:"rewritten_text"`
	AppliedStyle  string `json:"applied_style"`
}
