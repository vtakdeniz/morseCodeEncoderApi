package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_encodeToMorse(t *testing.T) {
	type args struct {
		message        string
		letterSplitter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Can encode to mors",
			args: args{
				message:        "hello world",
				letterSplitter: " ",
			},
			want: ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeToMorse(tt.args.message, tt.args.letterSplitter); got != tt.want {
				t.Errorf("encodeToMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serverEncodesCorrectly(t *testing.T) {
	type args struct {
		message        string
		letterSplitter string
	}
	tests := []struct {
		name  string
		args  string
		route string
		want  string
	}{
		{
			name:  "Server encodes and responds correctly",
			args:  "hello world",
			route: "/encode?text=",
			want:  ".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
		},
	}
	app := fiber.New()
	app.Get("/encode", func(ctx *fiber.Ctx) error {
		str := ctx.Query("text")
		if str == "" {
			return ctx.SendStatus(fiber.StatusOK)
		}
		return ctx.SendString(encodeToMorse(str, " "))
	})

	for _, tt := range tests {
		param := url.QueryEscape(tt.args)
		route := tt.route + param

		rec := httptest.NewRequest(http.MethodGet, route, nil)
		resp, _ := app.Test(rec, 1)
		respBody, _ := ioutil.ReadAll(resp.Body)
		respStr := string(respBody)

		assert.Equal(t, respStr, tt.want, tt.name)
	}
}
