package request

import (
	"golang.org/x/sys/windows/registry"
	"net/http"
	"net/url"
)

func NewHttpClient() (*http.Client, error) {
	proxyURL, err := getSystemProxy()
	if err != nil {
		return nil, err
	}
	if proxyURL != nil {
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}, nil
	} else {
		return &http.Client{}, nil
	}

}

func getSystemProxy() (*url.URL, error) {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Internet Settings`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	proxyEnable, _, err := key.GetIntegerValue("ProxyEnable")
	if err != nil {
		return nil, err
	}

	if proxyEnable == 1 {
		proxyServer, _, err := key.GetStringValue("ProxyServer")
		if err != nil {
			return nil, err
		}
		if proxyServer != "" {
			proxyURL, err := url.Parse("http://" + proxyServer)
			if err != nil {
				return nil, err
			}
			return proxyURL, nil
		}
	}

	return nil, nil
}
