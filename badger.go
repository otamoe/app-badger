package appbadger

import (
	"context"

	"github.com/dgraph-io/badger/v3"
	"go.uber.org/fx"
)

func NewFX(db *badger.DB) (fxOption fx.Option) {
	return fx.Provide(func(lc fx.Lifecycle) *badger.DB {
		ctx, cancel := context.WithCancel(context.Background())
		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				go GC(ctx, db)
				return nil
			},
			OnStop: func(c context.Context) error {
				cancel()
				return nil
			},
		})
		return db
	})
}
