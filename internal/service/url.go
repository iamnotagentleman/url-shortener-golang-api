package service

import (
	"awesome-url-shortener/internal/models"
	"context"
	"encoding/hex"
	"time"

	"github.com/twmb/murmur3"
)

func (s *service) ShortUrl(ctx context.Context, req models.UrlShortPayload) models.UrlShortCreateResponse {
	h := murmur3.New64()
	h.Write([]byte(req.Url))
	hashResult := h.Sum(nil)
	hashString := hex.EncodeToString(hashResult)

	err := s.store.SetUrl(ctx, hashString, req.Url, req.TTL)

	if err != nil {
		return models.UrlShortCreateResponse{}
	}
	shortenUrl := s.config.BASE_URL + "/go/" + hashString

	var expireAt string

	if req.TTL != 0 {
		expireAt = time.Now().Local().Add(time.Second * time.Duration(req.TTL)).Format(time.RFC3339)
	} else {
		expireAt = ""
	}

	return models.UrlShortCreateResponse{
		ShortenUrl: shortenUrl,
		ExpireAt:   &expireAt,
	}
}

func (s *service) ResolveShortUrl(ctx context.Context, req models.UrlShortGetInput) models.UrlShortGetResponse {
	val, err := s.store.GetUrl(ctx, req.Key)
	if err != nil {
		return models.UrlShortGetResponse{
			Error:    err,
			ShortUrl: "",
		}
	}

	return models.UrlShortGetResponse{
		Error:    nil,
		ShortUrl: val,
	}
}
