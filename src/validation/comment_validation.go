package validation

type QueryComment struct {
	Page   int    `validate:"omitempty,number,max=50"`
	Limit  int    `validate:"omitempty,number,max=50"`
	Search string `validate:"omitempty,max=50"`
}
type UpdateComment struct {
	Content     *string `json:"content,omitempty" validate:"omitempty,max=500"`
	IsAnonymous *bool   `json:"is_anonymous,omitempty" validate:"omitempty"`
	IsDeleted   *bool   `json:"is_deleted,omitempty" validate:"omitempty"`
}
