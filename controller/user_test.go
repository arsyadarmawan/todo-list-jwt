package controller

import (
	"net/http"
	"task/service"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestUserController_Login(t *testing.T) {
	type fields struct {
		authService service.AuthService
	}
	type args struct {
		writer  http.ResponseWriter
		request *http.Request
		params  httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UserController{
				authService: tt.fields.authService,
			}
			a.Login(tt.args.writer, tt.args.request, tt.args.params)
		})
	}
}
