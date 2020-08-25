package model

import "github.com/jinzhu/gorm"

// Services to DB wrappers
type Services struct {
	Gallery GalleryService
	User    UserService
}

// NewServices instatiates all the available services with one DB connection
func NewServices(connectionInfo string) (*Services, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Services{
		User:    NewUserService(db),
		Gallery: &galleryGorm{db},
	}, nil
}
