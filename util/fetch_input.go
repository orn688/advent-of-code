package util

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const sessionIDEnvVar = "AOC_SESSION_ID"

func FetchInput(ctx context.Context, year, day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	sessionID, ok := os.LookupEnv(sessionIDEnvVar)
	if !ok {
		return "", fmt.Errorf("no %q env var", sessionIDEnvVar)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionID,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(b), " \n"), nil
}
