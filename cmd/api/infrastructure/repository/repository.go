package repository

import (
	database "crave/shared/database"
)

type Repository struct {
	mysql *database.MysqlWrapper
}

func NewRepository(mysql *database.MysqlWrapper) *Repository {
	return &Repository{mysql: mysql}
}

// if err := r.mysql.Driver.Transaction(func(tx *gorm.DB) error {
// 	// Select the last item
// 	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
// 		Table(r.GetTable().Statement.Table).
// 		Where("done = ?", false).
// 		Order("id DESC").
// 		First(&target).Error; err != nil {
// 		return err
// 	}
// 	target.Done = true
// 	// Delete the selected item
// 	if err := tx.Save(&target).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }); err != nil {
// 	return nil, err
// }

// return &target, nil
