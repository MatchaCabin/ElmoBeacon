package handler

import (
	"ElmoBeacon/request"
	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/exec"
)

const Version = ""

func (a *App) GetVersion() string {
	return Version
}

func (a *App) GetLatestVersion() (string, error) {
	return request.GetLatestVersion()
}

func (a *App) UpdateSelf() error {
	resp, err := http.Get("https://gfl2bucket.mcc.wiki/ElmoBeacon.exe")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}

	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	cmd := exec.Command(execPath)
	err = cmd.Start()
	if err != nil {
		return err
	}
	os.Exit(0)

	return nil
}
