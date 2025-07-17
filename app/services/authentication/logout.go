package authentication

import "context"

func (i *impl) Logout(ctx context.Context, tokenString string) error {
	err := i.jwtService.RevokeToken(ctx, tokenString)
	if err != nil {
		return err
	}

	return nil
}
