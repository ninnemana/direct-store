package store

import (
	"context"
	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
)

type service struct{}

func (s service) Option() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			accept := ctx.Value(httptransport.ContextKeyRequestAccept)
			switch accept.(type) {
			case string:
			default:
				httptransport.EncodeJSONResponse(ctx, w, err)
				return
			}

			// accept = accept.(string)
			switch {
			case strings.Contains(accept.(string), "proto"):
				// marvin.EncodeProtoResponse(ctx, w, err)
			default:
				httptransport.EncodeJSONResponse(ctx, w, err)
			}
		}),
	}
}
