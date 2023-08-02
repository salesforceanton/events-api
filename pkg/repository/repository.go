package repository

type Repository struct {
	Authorization
	Events
}

type Authorization interface {
}

type Events interface {
}
