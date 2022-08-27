package repository

import (
	"github.com/rohitchauraisa1997/service-repo-pattern/entity"
)

// seperated PostRepository from Firestore to enable ease of use for different dbs
// and prevent cluttered code.
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
}
