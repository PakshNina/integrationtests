package ourservice

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "integrationtests/internal/ourservice/mock"
)

const (
	phone         = "+1 234567"
	userID uint32 = 1
)

func TestOurService_GetBirthdayByPhone(t *testing.T) {
	ctrl := gomock.NewController(t)

	birthdaysRepoMock := mock.NewMockbirthdaysRepo(ctrl)
	phonebookMock := mock.NewMockphonebook(ctrl)

	ourService := NewService(birthdaysRepoMock, phonebookMock)

	var birthday = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name     string
		mockFunc func()
		err      error
		birthday time.Time
	}{
		{
			"No phone in phonebook",
			func() {
				phonebookMock.EXPECT().GetUserIDByPhone(phone).Return(uint32(0), errNoPhone).Times(1)
			},
			errNoPhone,
			time.Time{},
		},
		{
			"No phone in phonebook",
			func() {
				phonebookMock.EXPECT().GetUserIDByPhone(phone).Return(userID, nil).Times(1)
				birthdaysRepoMock.EXPECT().GetBirthdayByID(userID).Return(time.Time{}, errNoBirthday).Times(1)
			},
			errNoBirthday,
			time.Time{},
		},
		{
			"Success",
			func() {
				phonebookMock.EXPECT().GetUserIDByPhone(phone).Return(userID, nil).MaxTimes(1)
				birthdaysRepoMock.EXPECT().GetBirthdayByID(userID).Return(birthday, nil).MaxTimes(1)
			},
			nil,
			birthday,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			bd, err := ourService.GetBirthdayByPhone(phone)
			assert.Equal(t, tc.birthday, bd)
			assert.Equal(t, tc.err, err)
		})
	}
}
