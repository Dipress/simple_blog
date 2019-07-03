package ability

import (
	"github.com/dipress/blog/internal/post"
)

// PostAbillity allows checking ability to view post.
type PostAbillity struct{}

// CanUpdate checks permission to update the post.
func (p PostAbillity) CanUpdate(userID int, post *post.Post) bool {
	return userID == post.UserID
}

// CanDelete checks permission to delete the post.
func (p PostAbillity) CanDelete(userID int, post *post.Post) bool {
	return userID == post.UserID
}
