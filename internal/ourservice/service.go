package ourservice

import (
	"errors"
	"time"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

var errNoPhone = errors.New("no phone")
var errNoBirthday = errors.New("no birthday")

type birthdaysRepo interface {
	GetBirthdayByID(userID uint32) (time.Time, error)
}

type phonebook interface {
	GetUserIDByPhone(phone string) (uint32, error)
}

type ourService struct {
	repo birthdaysRepo
	pb   phonebook
}

type OurService interface {
	GetBirthdayByPhone(phone string) (time.Time, error)
}

func NewService(repo birthdaysRepo, pb phonebook) OurService {
	return &ourService{
		repo: repo,
		pb:   pb,
	}
}

func (s ourService) GetBirthdayByPhone(phone string) (time.Time, error) {
	userID, err := s.pb.GetUserIDByPhone(phone)
	if err != nil {
		return time.Time{}, errNoPhone
	}
	birthday, err := s.repo.GetBirthdayByID(userID)
	if err != nil {
		return time.Time{}, errNoBirthday
	}
	return birthday, nil
}
