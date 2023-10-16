package app

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/The-Gleb/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string, url io.Reader) (*http.Response, string) {

	req, err := http.NewRequest(method, ts.URL+path, url)
	require.NoError(t, err)
	log.Println("make request")
	resp, err := ts.Client().Do(req)
	log.Println("request done")
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func Test_app_GetShortenedURL(t *testing.T) {
	s := storage.New()
	a := NewApp(s, "http://localhost:8080/")
	router := chi.NewRouter()
	router.Post("/", a.GetShortenedURL)
	router.Get("/{id}", a.GetFullURL)
	ts := httptest.NewServer(router)
	defer ts.Close()

	type want struct {
		id   string
		code int
	}
	tests := []struct {
		name    string
		a       *app
		address string
		url     string
		want    want
	}{
		{
			name:    "pos test #1",
			a:       a,
			address: "/",
			url:     "http://yandex.ru",
			want:    want{"EwHXdJfB", http.StatusCreated},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, body := testRequest(t, ts, "POST", tt.address, strings.NewReader(tt.url))
			assert.Equal(t, tt.want.code, resp.StatusCode)
			if resp.StatusCode != 201 {
				return
			}
			assert.Equal(t, a.baseURL+tt.want.id, body)
		})
	}
}

func Test_app_GetFullURL(t *testing.T) {
	s := storage.New()
	s.AddURL("id1", "https://practicum.yandex.ru/")
	s.AddURL("id2", "u")
	a := NewApp(s, "http://localhost:8080/")
	router := chi.NewRouter()
	router.Post("/", a.GetShortenedURL)
	router.Get("/{id}", a.GetFullURL)
	ts := httptest.NewServer(router)
	defer ts.Close()
	testClient := ts.Client()
	testClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	type want struct {
		url  string
		code int
	}
	tests := []struct {
		name    string
		a       *app
		address string
		id      string
		want    want
	}{
		{
			name:    "pos test #1",
			a:       a,
			address: "/id1",
			// id:      "id2",
			want: want{"https://practicum.yandex.ru/", http.StatusTemporaryRedirect},
		},
		{
			name:    "neg test #2",
			a:       a,
			address: "/id3",
			// id:      "id1",
			want: want{"", http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println("Start test")

			resp, _ := testRequest(t, ts, "GET", tt.address, nil)

			assert.Equal(t, tt.want.code, resp.StatusCode)
			if resp.StatusCode != 307 {
				return
			}

			assert.Equal(t, tt.want.url, resp.Header.Get("Location"))

		})
	}
}
