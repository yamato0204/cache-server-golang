package acontext

import (
	"context"
	"errors"

	"github.com/yamato0204/cache-server-golang/internal/handler/schema"
)

type key struct{}

func WithAPIResponse(ctx context.Context) context.Context {
	return context.WithValue(ctx, key{}, &schema.APIResponse{
		Common:   &schema.CommonResponse{},
		Original: &schema.OriginalResponse{},
	})
}

func ExtractAPIResponse(ctx context.Context) (*schema.APIResponse, error) {
	v, ok := ctx.Value(key{}).(*schema.APIResponse)
	if !ok {
		return nil, errors.New("api response not found")
	}

	return v, nil
}
