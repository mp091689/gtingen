package controllers

type iController interface {
	Index(args map[string]string) error
}
