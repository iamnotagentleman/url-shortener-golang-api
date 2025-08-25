package models

type UrlShortCreateResponse struct {
	ShortenUrl string  `json:"shorten_url"`
	ExpireAt   *string `json:"expire_at"`
}

type UrlShortGetResponse struct {
	ShortUrl string `json:"short_url"`
	Error    error  `json:"error"`
}
