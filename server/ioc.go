package server

import (
	"fmt"
	repo "prod/repository"
	"prod/service"

	"gorm.io/gorm"
)

func Inject(d *gorm.DB) {
	var userRepo = repo.NewUserRepo(d)
	var userSerice = service.NewUserService(userRepo)
	fmt.Println(userSerice)
}
