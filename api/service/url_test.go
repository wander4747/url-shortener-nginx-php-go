package service_test

import (
	"errors"
	"reflect"
	"testing"
	"url_shortener/domain/entity"
	"url_shortener/domain/interfaces"
	"url_shortener/service"

	"github.com/stretchr/testify/assert"

	mock "url_shortener/tests/mocks"

	"github.com/golang/mock/gomock"
)

func TestNewURLService(t *testing.T) {
	ctrl := gomock.NewController(t)
	newUrlRepository := mock.NewMockUrlRepository(ctrl)

	type args struct {
		repository interfaces.UrlRepository
	}

	tests := []struct {
		name string
		args args
		want *service.URLService
	}{
		{
			"New url service",
			args{repository: newUrlRepository},
			&service.URLService{Repository: newUrlRepository},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewURLService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURLService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	newUrlRepository := mock.NewMockUrlRepository(ctrl)

	url := entity.URL{
		ID:   1,
		Link: "http://www.google.com",
		Hash: "1234abc",
	}

	type fields struct {
		Repository interfaces.UrlRepository
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   entity.URL
		err    error
	}{
		{
			"Success",
			fields{Repository: newUrlRepository},
			args{hash: "1234abc"},
			url,
			nil,
		},
		{
			"Error",
			fields{Repository: newUrlRepository},
			args{hash: "1234abc"},
			entity.URL{},
			errors.New("not found url"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlService := service.URLService{
				Repository: tt.fields.Repository,
			}

			newUrlRepository.EXPECT().Find(tt.args.hash).Return(tt.want, tt.err)
			got, err := urlService.Find(tt.args.hash)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestURLService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	newUrlRepository := mock.NewMockUrlRepository(ctrl)
	url := entity.URL{
		ID:   1,
		Link: "http://www.google.com",
		Hash: "1234abc",
	}

	type fields struct {
		Repository interfaces.UrlRepository
	}
	type args struct {
		data entity.URL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   entity.URL
		err    error
	}{
		{
			"Success",
			fields{Repository: newUrlRepository},
			args{data: entity.URL{Link: "http://www.google.com.br"}},
			url,
			nil,
		},
		{
			"Error",
			fields{Repository: newUrlRepository},
			args{data: entity.URL{}},
			entity.URL{},
			errors.New("error create url"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := service.URLService{
				Repository: tt.fields.Repository,
			}

			newUrlRepository.EXPECT().Save(tt.args.data).Return(tt.want, tt.err)
			got, err := u.Save(tt.args.data)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, tt.err, err)
		})
	}
}
