package service

import (
	"time"

	"github.com/Vardan1995/filter_notifyer/config"
	"github.com/Vardan1995/filter_notifyer/entity"
	"github.com/Vardan1995/filter_notifyer/repository"
	"github.com/golang-jwt/jwt"
)

type user_service struct{}

var (
	userRepository = repository.NewUserRepository()
)

func NewUserService() UserService {
	return &user_service{}
}

func (*user_service) CreateUser(user *entity.User) (*entity.User, string, error) {

	if err := userRepository.SaveUser(user); err != nil {
		return nil, "", err
	}
	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, "", err
	}
	return user, token, err
}

func (*user_service) LoginUser(user *entity.User, username, email string) (string, error) {
	if err := userRepository.FindUserByUsernameAndEmail(user, username, email); err != nil {
		return "", err
	}
	t, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return t, err
}

func (*user_service) GetUser(user *entity.User, id uint) (*entity.User, error) {

	if err := userRepository.FindUser(user, id); err != nil {
		return nil, err
	}
	return user, nil
}
func (*user_service) GetUsers(users *[]entity.User) error {

	if err := userRepository.FindUsers(users); err != nil {
		return err
	}
	return nil
}
func generateJWT(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}
