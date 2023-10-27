package repositories

import (
	db "athe-autho-go/internal/database"
	"athe-autho-go/internal/models"
	"errors"
	"time"
)

func FindUserByNameAndPassword(username string, password string) (*models.User, error) {
	users := make(chan *models.User, 1)
	go func() {
		user := models.User{}
		db.Init().Model(&models.User{}).Where("username = ? AND password = ?", username, password).First(&user)
		users <- &user
	}()
	select {
	case loaded := <-users:
		if loaded == nil || len(loaded.Username) < 1 {
			return nil, errors.New("invalid credentials")
		}
		return loaded, nil
	case <-time.After(3 * time.Second):
		return nil, errors.New("user loading timed out")
	}
}
