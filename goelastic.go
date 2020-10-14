package goelasticsearch

import "context"

type Repository interface {
	Fetch(ctx context.Context) ([]map[string]interface{}, error)
	Store(ctx context.Context, document map[string]interface{}) error
	Update(ctx context.Context, id string, document map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (map[string]interface{}, error)
}
