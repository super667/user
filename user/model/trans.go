package model

import "context"

type (
	TransModel interface {
	}
	customTransModel struct {
		*defaultTransModel
	}
	defaultTransModel struct {
	}
)

func (t *defaultTransModel) TransRecord(ctx context.Context) error {
	return nil
}
