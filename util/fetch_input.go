package util

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const sessionIDEnvVar = "AOC_SESSION_ID"

// FetchInput retrieves the puzzle input for a given day and year.
//
// It includes a local filesystem caching layer to avoid unnecessary requests to
// the AoC server. The cache is stored in the root of the current working
// directory's git repo.
func FetchInput(ctx context.Context, year, day int) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--show-toplevel")
	var stdout strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get git root directory: %w", err)
	}

	rootDir := strings.TrimSpace(stdout.String())
	cacheDir := filepath.Join(rootDir, ".cache", strconv.Itoa(year))

	cacheFile := filepath.Join(cacheDir, fmt.Sprintf("day%02d.txt", day))
	b, err := os.ReadFile(cacheFile)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return "", fmt.Errorf("failed to check if cache file exists: %w", err)
	} else if err == nil {
		return string(b), nil
	}

	input, err := doFetch(ctx, year, day)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(cacheDir, 0o700); err != nil {
		return "", err
	}

	if err := os.WriteFile(cacheFile, input, 0o600); err != nil {
		return "", err
	}

	return string(input), err
}

func doFetch(ctx context.Context, year, day int) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	sessionID, ok := os.LookupEnv(sessionIDEnvVar)
	if !ok {
		return nil, fmt.Errorf("no %q env var", sessionIDEnvVar)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionID,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
