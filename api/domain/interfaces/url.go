package interfaces

import "url_shortener/domain/entity"

type UrlRepository interface {
	Save(entity.URL) (entity.URL, error)
	Find(hash string) (entity.URL, error)
}
