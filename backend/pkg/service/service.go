package service

import "net/http"

type UserService interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Edit() http.HandlerFunc
}
