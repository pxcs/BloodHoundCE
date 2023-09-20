package database

import (
	"github.com/gofrs/uuid"
	"github.com/specterops/bloodhound/src/model"
)

func (s *BloodhoundDB) ListSavedQueries(userID uuid.UUID, order string, filter model.SQLFilter, skip, limit int) (model.SavedQueries, int, error) {
	var (
		queries model.SavedQueries
		count   int64
	)

	cursor := s.Scope(Paginate(skip, limit)).Where("user_id = ?", userID)

	if filter.SQLString != "" {
		cursor = cursor.Where(filter.SQLString, filter.Params)
	}

	if order != "" {
		cursor = cursor.Order(order)
	}

	result := s.db.Find(&queries).Count(&count)
	if result.Error != nil {
		return queries, 0, result.Error
	}

	result = cursor.Find(&queries)
	return queries, int(count), CheckError(result)
}
