package entity

import "url_shortener/lib/base62"

type URL struct {
	ID   uint   `json:"id,omitempty"`
	Link string `json:"link" binding:"required,url"`
	Hash string `json:"hash,omitempty"`
}

func NewUrl(sizeHash int) URL {
	url := URL{
		Hash: base62.Encode(sizeHash),
	}

	return url
}