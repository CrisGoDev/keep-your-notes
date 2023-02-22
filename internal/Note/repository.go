package note

import (
	"fmt"
	"log"
	"strings"

	"github.com/CrisGoDev/keep-your-notes/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(note *domain.Note) error
	GetAll(filters Filters, ofsset int, limit int) ([]domain.Note, error)
	Get(id string) (*domain.Note, error)
	Delete(id string) error
	Udpate(id string, title *string, body *string) error
	Count(filters Filters) (int, error)
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {

	return &repo{
		db:  db,
		log: log,
	}
}

func (repo *repo) Create(note *domain.Note) error {

	if err := repo.db.Create(note).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("note created with the id:", note.Id)

	return nil
}

func (repo *repo) GetAll(filters Filters, ofsset int, limit int) ([]domain.Note, error) {
	var users []domain.Note

	db := repo.db.Model(&users)
	db = applyFilters(db, filters)
	db = db.Limit(limit).Offset(ofsset)

	result := db.Order(filters.OrderBy + " " + filters.Order).Find(&users)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (repo *repo) Get(id string) (*domain.Note, error) {
	note := domain.Note{Id: id}

	result := repo.db.First(&note)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return nil, result.Error
	}

	return &note, nil
}

func (repo *repo) Delete(id string) error {
	user := domain.Note{Id: id}

	result := repo.db.Delete(&user)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (repo *repo) Udpate(id string, title *string, body *string) error {

	note := domain.Note{Id: id}
	if err := repo.db.First(&note).Error; err != nil {
		return err
	}

	values := make(map[string]interface{})

	if title != nil {
		values["title"] = *title
	}

	if body != nil {
		values["body"] = *body
	}

	if err := repo.db.Model(&domain.Note{}).Where("id = ?", id).Updates(values).Error; err != nil {

		repo.log.Println("Error on repository", err.Error())
		return err
	}

	return nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.Body != "" {
		filters.Body = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Body))
		tx = tx.Where("lower(body) like ?", filters.Body)

	}

	if filters.Title != "" {
		filters.Title = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Title))
		tx = tx.Where("lower(title) Like ?", filters.Title)

	}

	return tx
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64

	var users []domain.Note

	db := repo.db.Model(&users)
	db = applyFilters(db, filters)

	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
