package government

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	Scheme = "https" // http or https
	Host   = "info.midpass.ru"
)

type Application struct {
	Info          applicationInfo `json:"applicationInfo"`
	PassportReady bool            `json:"passportReady"`
	ResultCode    int             `json:"resultCode"`
	Country       string          `json:"country"`
	City          string          `json:"city"`
}

type applicationInfo struct {
	ID              string            `json:"uid"`
	Status          applicationStatus `json:"applicationStatus"`
	ApplicationDate interface{}       `json:"applicationDate"`
	BirthDate       interface{}       `json:"birthDate"`
}
type applicationStatus struct {
	ID             int            `json:"id"`
	Percent        int            `json:"percent"`
	Name           interface{}    `json:"name"`
	PassportStatus passportStatus `json:"passportStatus"`
}
type passportStatus struct {
	ID            int    `json:"id"`
	StatusDescEng string `json:"englishName"`
	StatusDescRus string `json:"russianName"`
	Terminal      bool   `json:"terminal"`
	Notifiable    bool   `json:"notifiable"`
	Color         string `json:"color"`
}

func GetApplication(id string) (Application, error) {
	var application Application
	resp, err := http.Get(getServerURL() + "/svc/pi/app/34102/" + id)
	if err != nil {
		return application, err
	}
	if resp.StatusCode != 200 {
		return application, errors.New("Failed to load application. Retuned status: " + resp.Status + ".")
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return application, err
	}

	err = json.Unmarshal(body, &application)
	if err != nil {
		return application, err
	}

	return application, nil
}

func getServerURL() string { return Scheme + "://" + Host }
