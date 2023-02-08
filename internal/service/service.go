package service

import (
	"encoding/hex"
	"github.com/dwnGnL/validation/internal/config"
	"github.com/dwnGnL/validation/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

type repositoryIter interface {
	GetTestTable() (*repository.TestTable, error)
	Registration(users *repository.Users) error
	Login(users string) (*repository.Users, error)
	SaveToken(token repository.Tokens) error
}

type ServiceImpl struct {
	conf *config.Config
	repo repositoryIter
}

type Option func(*ServiceImpl)

func New(conf *config.Config, repo repositoryIter, opts ...Option) *ServiceImpl {
	s := ServiceImpl{
		conf: conf,
		repo: repo,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return &s
}

func (s ServiceImpl) TestService() string {
	return "it`s test"
}

func (s ServiceImpl) Registration(info *repository.Users) error {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return err
	}
	token := hex.EncodeToString(buf)

	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	info.AccessToken = token
	info.Password = string(hash)
	info.Active = true

	err = s.repo.Registration(info)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s ServiceImpl) Login(info *repository.Users) (string, error) {
	userData, err := s.repo.Login(info.Login)
	if err != nil {
		log.Println(err)
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(info.Password))
	if err != nil {
		log.Println("error in CompareHash")
		return "", err
	}
	log.Println(string(hash))

	//token := jwt.New(jwt.SigningMethodHS256)
	//
	//claims := token.Claims.(jwt.MapClaims)
	//claims["auth"] = true
	//claims["user"] = "test"
	//claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	//
	//tokenString, err := token.SignedString([]byte("JWTtoken"))
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}

	//==========================================================================
	//	openssl genrsa -out privatekey.pem 2048
	//	openssl rsa -in privatekey.pem -out publickey.pem -pubout -outform PEM

	prvKey, err := ioutil.ReadFile("privatekey.pem")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := ioutil.ReadFile("publickey.pem")
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := NewJWT(prvKey, pubKey)

	tokenString, err := jwtToken.Create(time.Hour, "test")
	if err != nil {
		log.Println(err)
		return "", err
	}

	var TokenData repository.Tokens
	//
	TokenData.Token = tokenString
	TokenData.UserID = userData.ID

	os := runtime.GOOS
	switch os {
	case "windows":
		TokenData.Platform = "Windows"
	case "darwin":
		TokenData.Platform = "MAC operating system"
	case "linux":
		TokenData.Platform = "Linux"
	default:
		TokenData.Platform = os
	}

	log.Println(TokenData)
	err = s.repo.SaveToken(TokenData)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
