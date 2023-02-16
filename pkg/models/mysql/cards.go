package mysql

import (
	"database/sql"

	"github.com/udaycruise2903/cardbox/pkg/models"
)

type CardModel struct {
	DB *sql.DB
}

// To insert a new card into the database
func (m *CardModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// to return a specific card based on its id
func (m *CardModel) Get(id int) (*models.Card, error) {
	return nil, nil
}

// to return the 10 most recently created cards
func (m *CardModel) Latest() ([]*models.Card, error) {
	return nil, nil
}
