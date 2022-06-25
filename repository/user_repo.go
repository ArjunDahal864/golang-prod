package repo

import (
	"prod/db"

	"gorm.io/gorm"
)

func NewUserRepo(d *gorm.DB) *UserRepo {
	return &UserRepo{
		db: d,
	}
}

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) Create(user *db.User) (*db.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Get(id int) (*db.User, error) {
	var user db.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetByEmail(email string) (*db.User, error) {
	var user db.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetAllUsers() ([]db.User, error) {
	var users []db.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) Update(user *db.User) (*db.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
