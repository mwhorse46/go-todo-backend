package todos

import (
	"context"

	"github.com/go-rel/rel"
)

type clear struct {
	repository rel.Repository
}

func (c clear) Clear(ctx context.Context) {
	c.repository.MustDeleteAll(ctx, rel.From("todos"))
}
