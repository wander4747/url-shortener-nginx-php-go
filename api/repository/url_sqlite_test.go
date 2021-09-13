package repository_test

import (
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"url_shortener/domain/entity"
	"url_shortener/domain/interfaces"
	"url_shortener/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewUrlDBSqlite(t *testing.T) {
	db, _, _ := sqlmock.New()
	repositoryMock := repository.NewUrlDBSqlite(db)

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *repository.URLRepositoryDBSqlite
	}{
		{
			"New url repository",
			args{db: db},
			repositoryMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewUrlDBSqlite(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrlDBSqlite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLRepositoryDBSqlite_Find(t *testing.T) {
	db, mock, _ := sqlmock.New()
	repositoryMock := repository.NewUrlDBSqlite(db)
	sql := "SELECT id, link, hash FROM urls WHERE hash = ?"

	tests := []struct {
		name    string
		i       interfaces.UrlRepository
		mock    func()
		hash    string
		want    entity.URL
		err     error
		wantErr bool
	}{
		{
			name: "Success",
			i:    repositoryMock,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "link", "hash"}).AddRow(1, "http://www.google.com", "1234abc")
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("1234abc").WillReturnRows(rows)
			},
			hash:    "1234abc",
			want:    entity.URL{ID: 1, Link: "http://www.google.com", Hash: "1234abc"},
			err:     nil,
			wantErr: false,
		},
		{
			name: "Not found",
			i:    repositoryMock,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "link", "hash"})
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("9876abc").WillReturnRows(rows)
			},
			hash:    "9876abc",
			want:    entity.URL{},
			err:     nil,
			wantErr: false,
		},
		{
			name: "Error query",
			i:    repositoryMock,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "link", "hash-wrong"})
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("1234abc").WillReturnRows(rows)
			},
			hash:    "9876abc",
			want:    entity.URL{},
			err:     errors.New("error query"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.i.Find(tt.hash)

			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLRepositoryDBSqlite_Save(t *testing.T) {
	db, mock, _ := sqlmock.New()
	repositoryMock := repository.NewUrlDBSqlite(db)
	sql := "INSERT INTO urls"

	tests := []struct {
		name    string
		i       interfaces.UrlRepository
		mock    func()
		data    entity.URL
		want    entity.URL
		wantErr bool
	}{
		{
			name: "Success",
			i:    repositoryMock,
			mock: func() {
				mock.ExpectPrepare(sql).ExpectExec().WithArgs("http://www.google.com.br", "it65dBfr").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			data:    entity.URL{ID: 1, Link: "http://www.google.com.br", Hash: "it65dBfr"},
			want:    entity.URL{ID: 1, Link: "http://www.google.com.br", Hash: "it65dBfr"},
			wantErr: false,
		},
		{
			name: "Empty link",
			i:    repositoryMock,
			mock: func() {
				mock.ExpectPrepare(sql).ExpectExec().WithArgs("http://www.google.com.br", "it65dBfr").WillReturnError(errors.New("empty link"))
			},
			data:    entity.URL{ID: 1, Link: "", Hash: "it65dBfr"},
			want:    entity.URL{ID: 1, Link: "http://www.google.com.br", Hash: "it65dBfr"},
			wantErr: true,
		},
		{
			name: "Invalid SQL query",
			i:    repositoryMock,
			mock: func() {
				mock.ExpectPrepare("INSERT INTO wrong_table").ExpectExec().WithArgs("http://www.google.com.br", "it65dBfr").WillReturnError(errors.New("invalid sql query"))
			},
			data:    entity.URL{ID: 1, Link: "", Hash: "it65dBfr"},
			want:    entity.URL{ID: 1, Link: "http://www.google.com.br", Hash: "it65dBfr"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.i.Save(tt.data)

			if tt.wantErr == false {
				assert.Equal(t, got.ID, tt.want.ID)
				assert.Equal(t, got.Link, tt.want.Link)
			} else {
				assert.Error(t, err)
			}

			//u := URLRepositoryDBSqlite{
			//	db: tt.fields.db,
			//}
			//got, err := tt.i.Save(tt.data)
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Save() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
