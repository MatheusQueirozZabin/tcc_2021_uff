package usersvc

import (
	"context"
	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
)

func (s UserService) GetUserInfo(
	ctx context.Context,
	userID uservo.UserID,
	addressID uservo.AddressID,
) (userInfo userent.RegisteredUser, err error) {
	return s.repo.GetUserInfo(ctx, userID)
}
