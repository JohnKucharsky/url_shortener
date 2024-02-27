package dbstore

import (
	"errors"
	"github.com/JohnKucharsky/url_shortener/store"
	"log/slog"
)

type ShortURLStore struct {
	shortURLs []store.ShortURL
	logger    *slog.Logger
}

func NewShortURLStore(logger *slog.Logger) *ShortURLStore {
	var shortURLs []store.ShortURL
	return &ShortURLStore{shortURLs: shortURLs, logger: logger}
}

func (s *ShortURLStore) CreateShortURL(params store.CreateShortURLParams) (
	store.ShortURL,
	error,
) {

	shortURL := store.ShortURL{
		Destination: params.Destination,
		Slug:        params.Slug,
		ID:          len(s.shortURLs),
	}

	s.shortURLs = append(s.shortURLs, shortURL)

	s.logger.Info("short URL created", slog.Any("values", shortURL))

	return shortURL, nil
}

func (s *ShortURLStore) GetShortURLBySlug(slug string) (
	*store.ShortURL,
	error,
) {
	for _, shortURL := range s.shortURLs {
		if shortURL.Slug == slug {
			return &shortURL, nil
		}
	}

	return nil, errors.New("short URL not found")
}
