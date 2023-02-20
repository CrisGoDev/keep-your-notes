package note

import (
	"log"

	"github.com/CrisGoDev/keep-your-notes/domain"
	"gorm.io/gorm"
)

type Repository interface {
	// Create(user *Note) error
	// GetAll(filters Filters, ofsset int, limit int) ([]Note, error)
	Get() (*[]domain.Note, error)
	// Delete(id string) error
	// Udpate(id string, FirstName *string, LastName *string, Email *string, Phone *string) error
	// Count(filters Filters) (int, error)
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

// func (repo *repo) Create(user *Note) error {

// 	if err := repo.db.Create(user).Error; err != nil {
// 		repo.log.Println(err)
// 		return err
// 	}

// 	repo.log.Println("user creates with the id:", user.Id)

// 	return nil
// }

// func (repo *repo) GetAll(filters Filters, ofsset int, limit int) ([]Note, error) {
// 	var users []Note

// 	db := repo.db.Model(&users)
// 	db = applyFilters(db, filters)
// 	db = db.Limit(limit).Offset(ofsset)

// 	result := db.Order("created_at desc").Find(&users)

// 	if result.Error != nil {
// 		repo.log.Println(result.Error)
// 		return nil, result.Error
// 	}

// 	return users, nil
// }

// func (repo *repo) Get(id string) (*note.Note, error) {
// 	user := note.Note{Id: id}

// 	result := repo.db.First(&user)

// 	if result.Error != nil {
// 		repo.log.Println(result.Error)
// 		return nil, result.Error
// 	}

// 	return &user, nil
// }

func (repo *repo) Get() (*[]domain.Note, error) {
	var notes []domain.Note

	result := repo.db.Model(&notes).Order("created_at desc").Find(&notes)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return nil, result.Error
	}

	return &notes, nil
}

// func (repo *repo) Delete(id string) error {
// 	user := Note{Id: id}

// 	result := repo.db.Delete(&user)

// 	if result.Error != nil {
// 		repo.log.Println(result.Error)
// 		return result.Error
// 	}

// 	return nil
// }

// func (repo *repo) Udpate(id string, FirstName *string, LastName *string, Email *string, Phone *string) error {
// 	values := make(map[string]interface{})

// 	if FirstName != nil {
// 		values["first_name"] = *FirstName
// 	}

// 	if LastName != nil {
// 		values["last_name"] = *LastName
// 	}

// 	if Email != nil {
// 		values["email"] = *Email
// 	}

// 	if Phone != nil {
// 		values["phone"] = *Phone
// 	}

// 	if err := repo.db.Model(&Note{}).Where("id = ", id).Updates(values).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

// 	if filters.FirstName != "" {
// 		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
// 		tx = tx.Where("lower(first_name) like ?", filters.FirstName)

// 	}

// 	if filters.LastName != "" {
// 		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
// 		tx = tx.Where("lower(last_name) Like ?", filters.LastName)

// 	}

// 	return tx
// }

// func (repo *repo) Count(filters Filters) (int, error) {
// 	var count int64

// 	var users []Note

// 	db := repo.db.Model(&users)
// 	db = applyFilters(db, filters)

// 	if err := db.Count(&count).Error; err != nil {
// 		return 0, err
// 	}

// 	repo.log.Println("Numero de todos los registros", count)
// 	return int(count), nil
// }
