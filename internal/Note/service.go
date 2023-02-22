package note

import (
	"log"

	"github.com/CrisGoDev/keep-your-notes/domain"
)

const (
	Asc  = "asc"
	Desc = "desc"
)

type MaritalStatus string

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
	Widowed  MaritalStatus = "widowed"
)

type (
	Filters struct {
		Body      string
		Title     string
		CreatedAt string
		OrderBy   string
		Order     string
	}

	Service interface {
		GetAll(filters Filters, ofsset int, limit int) ([]domain.Note, error)
		Count(filters Filters) (int, error)
		Create(title, body string) (*domain.Note, error)
		Get(id string) (*domain.Note, error)
		Update(id string, title *string, body *string) error
		Delete(id string) error
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {

	return &service{
		repo: repo,
		log:  log,
	}
}

func (s service) Get(id string) (*domain.Note, error) {

	note, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s service) Create(title, body string) (*domain.Note, error) {

	user := domain.Note{
		Title: title,
		Body:  body,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s service) GetAll(filters Filters, offset int, limit int) ([]domain.Note, error) {

	if filters.OrderBy == "" {
		filters.OrderBy = "created_at"
	}

	if filters.Order == "" {
		filters.Order = "asc"
	}
	users, err := s.repo.GetAll(filters, offset, limit)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}

func (s service) Delete(id string) error {
	err := s.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s service) Update(id string, title *string, body *string) error {

	s.log.Println("id en servicio a actualizar", id)
	err := s.repo.Udpate(id, title, body)

	if err != nil {
		return err
	}

	return nil
}
