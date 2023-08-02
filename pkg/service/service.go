package service

type Service struct {
	Authorization
	Events
}

type Authorization interface {
}

type Events interface {
}
