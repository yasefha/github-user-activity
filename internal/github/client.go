package github

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func getUrl(username string) string {
	base := "https://api.github.com/users/"
	activity := "/events/public"

	safeUsername := url.PathEscape(username)
	fullURL := base + safeUsername + activity

	return fullURL
}

func FetchActivity(username string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", getUrl(username), nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
