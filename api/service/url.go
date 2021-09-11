package service

import (
	"url_shortener/domain/entity"
	"url_shortener/domain/interfaces"
)

type URLService struct {
	Repository interfaces.UrlRepository
}

func NewURLService(repository interfaces.UrlRepository) *URLService {
	return &URLService{Repository: repository}
}

func (urlService URLService) Find(hash string) (entity.URL, error) {
	url, err := urlService.Repository.Find(hash)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func (u URLService) Save(data entity.URL) (entity.URL, error) {
	url, err := u.Repository.Save(data)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}