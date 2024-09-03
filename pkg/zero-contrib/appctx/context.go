package appctx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"oracleX-backend/pkg/zero-contrib/interceptorx"

	"google.golang.org/grpc/metadata"
)

type AppContext struct {
	Uid int64
}

func GetContext(ctx context.Context) *AppContext {
	uid := ctx.Value(interceptorx.UID).(int64)

	return &AppContext{
		Uid: uid,
	}
}

func GetMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", MetadataNotFound
	}
	List := md.Get(fmt.Sprintf("gateway-%s", key))

	if len(List) == 0 || List[0] == "" {
		return "", EmptyMetadata
	}
	return List[0], nil
}

func GetUid(ctx context.Context) int64 {
	return ctx.Value(interceptorx.UID).(int64)
}

func GetTelegramUid(ctx context.Context) (int64, error) {
	uidStr, err := GetMetadata(ctx, "tg-uid")
	if err != nil {
		return 0, errors.Wrap(err, "fail to get telegram uid")
	}

	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		return 0, errors.New("invalid telegram uid")
	}
	return uid, nil
}

func GetTelegramUsername(ctx context.Context) (string, error) {
	username, err := GetMetadata(ctx, "tg-username")
	if err != nil {
		return "", errors.Wrap(err, "fail to get telegram username")
	}
	return username, nil
}

func GetTelegramFirstName(ctx context.Context) (string, error) {
	firstName, err := GetMetadata(ctx, "tg-first-name")
	if err != nil {
		return "", errors.Wrap(err, "fail to get telegram first name")
	}
	return firstName, nil
}

func GetTelegramLastName(ctx context.Context) (string, error) {
	lastName, err := GetMetadata(ctx, "tg-last-name")
	if err != nil {
		return "", errors.Wrap(err, "fail to get telegram last name")
	}
	return lastName, nil
}
