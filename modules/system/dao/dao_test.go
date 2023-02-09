package dao

import (
	"atom/container"
	"context"
	"testing"

	"atom/providers/config"
	_ "atom/providers/database"
	_ "atom/providers/http"
	_ "atom/providers/log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Fields struct {
	dig.In

	Conf *config.Config
	DB   *gorm.DB
}

func TestDaoImpl_Release(t *testing.T) {
	var ff Fields
	err := container.Container.Invoke(func(f Fields) {
		ff = f
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	type args struct {
		ctx context.Context
		a   int
		b   string
	}
	tests := []struct {
		name    string
		fields  Fields
		args    args
		wantErr bool
	}{
		{"1. ", ff, args{context.Background(), 10, "Rogee"}, false},
		{"2. ", ff, args{context.Background(), 20, "Rogee"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Dao{
				Conf: tt.fields.Conf,
				DB:   tt.fields.DB,
			}
			if err := c.Release(tt.args.ctx, tt.args.a, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("DaoImpl.Release() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
