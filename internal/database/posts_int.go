package database

import "context"

func (q *Queries) GetPostsForUserInt(ctx context.Context, limit int) ([]Post, error) {
	return q.GetPostsForUser(ctx, int32(limit))
}
