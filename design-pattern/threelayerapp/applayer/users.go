package applayer

import (
	"context"

	"github.com/dtherhtun/Learning-go/design-pattern/threelayerapp/storelayer"
)

func (a *app) CreateUser(ctx context.Context, name, handle string) error {
	return a.store.CreateUser(ctx, name, handle)
}

func (a *app) GetAllUsers(ctx context.Context) ([]storelayer.User, error) {
	return a.store.GetAllUsers(ctx)
}
