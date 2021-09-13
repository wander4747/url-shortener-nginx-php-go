//go:generate mockgen -destination=../../tests/mocks/url_mock.go -package=mock -source=url.go interfaces
package interfaces

import "url_shortener/domain/entity"

type UrlRepository interface {
	Save(entity.URL) (entity.URL, error)
	Find(hash string) (entity.URL, error)
}
