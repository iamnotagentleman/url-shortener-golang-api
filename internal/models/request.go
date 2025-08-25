package models

type UrlShortCreateInput struct {
	Payload *UrlShortPayload `in:"body=json"`
}

type UrlShortPayload struct {
	Url string `json:"url"`
	TTL int    `json:"ttl"`
}

type UrlShortGetInput struct {
	Key string `in:"path=key"`
}
