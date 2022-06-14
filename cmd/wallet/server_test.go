package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet/internal/postgres"
	"wallet/internal/rabbitmq"
	"wallet/internal/request"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleDeposit(t *testing.T) {
	db, err := postgres.NewTestConn()
	assert.NoError(t, err)
	defer db.Close(context.Background())

	amqp, err := rabbitmq.NewConn()
	assert.NoError(t, err)
	defer amqp.Close()

	s := newServer()
	s.echo.POST("/deposit", s.handleDeposit)
	s.db = db
	s.amqp = amqp

	var (
		r = request.Deposit{
			Receiver: 1,
			Amount:   100,
		}
		b bytes.Buffer
	)
	err = json.NewEncoder(&b).Encode(r)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/deposit", &b)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	if assert.NoError(t, s.handleDeposit(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success", rec.Body.String())
	}
}

func TestHandleTransfer(t *testing.T) {
	db, err := postgres.NewTestConn()
	assert.NoError(t, err)
	defer db.Close(context.Background())

	amqp, err := rabbitmq.NewConn()
	assert.NoError(t, err)
	defer amqp.Close()

	s := newServer()
	s.echo.POST("/transfer", s.handleTransfer)
	s.db = db
	s.amqp = amqp

	var (
		r = request.Transfer{
			Sender:   1,
			Receiver: 3,
			Amount:   100,
		}
		b bytes.Buffer
	)
	err = json.NewEncoder(&b).Encode(r)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/transfer", &b)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	if assert.NoError(t, s.handleTransfer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success", rec.Body.String())
	}
}
