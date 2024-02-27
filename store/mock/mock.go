package mock

import (
	"github.com/JohnKucharsky/url_shortener/store"
	"github.com/stretchr/testify/mock"
)

type ShortURLStoreMock struct {
	mock.Mock
}

func (m *ShortURLStoreMock) CreateShortURL(params store.CreateShortURLParams) (
	store.ShortURL,
	error,
) {
	args := m.Called(params)

	return args.Get(0).(store.ShortURL), args.Error(1)
}

func (m *ShortURLStoreMock) GetShortURLBySlug(slug string) (
	*store.ShortURL,
	error,
) {
	args := m.Called(slug)

	return args.Get(0).(*store.ShortURL), args.Error(1)
}
