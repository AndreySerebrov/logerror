package client

import (
	"err_test/apperrors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Client struct {
	log *logrus.Logger
}

func NewClient(log *logrus.Logger) Client {
	return Client{log: log}
}

func (c Client) Get() error {
	// Обращаемся к внешнему сервису
	status := doExternalRequest()

	// Матчим ошибки сервиса в доменные ошибки приложения
	switch status {
	case http.StatusNotFound:
		c.log.Errorf("client.ExternalRequest.%d", status)
		return apperrors.ErrNotFound
	case http.StatusInternalServerError:
		c.log.Errorf("client.ExternalRequest.%d", status)
		return apperrors.ErrInternalError
	}

	return nil
}

func (с Client) Put() error {
	// Вызов завершился успешно
	return nil
}

func (с Client) Update() error {
	// Вызов завершился успешно
	return nil
}

func doExternalRequest() int {
	return http.StatusNotFound
}
