package domain


// Claims is not a separate entity but a type embedded in other domain entities
type Claims map[string]any

const (
	UserKindClaim = "user_kind"
)

func (c Claims) GetUserKind() string {
	kind, found := c[UserKindClaim]
	if !found {
		return ""
	}
	return kind.(string)
}
