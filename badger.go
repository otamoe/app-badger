package appbadger

import (
	"context"

	"github.com/dgraph-io/badger/v3"
	"go.uber.org/fx"
)

func NewFX(db *badger.DB) (fxOption fx.Option) {
	return fx.Provide(func(ctx context.Context, lc fx.Lifecycle) *badger.DB {
		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				go GC(ctx, db)
				return nil
			},
		})
		return db
	})
}
