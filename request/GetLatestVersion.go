package request

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func GetLatestVersion() (string, error) {
	resp, err := http.Get("https://gfl2worker.mcc.wiki/ElmoBeacon/version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
