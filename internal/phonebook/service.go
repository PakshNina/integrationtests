package phonebook

type APIPhonebook interface {
	GetUserIDByPhone(phone string) (uint32, error)
	SavePhone(phone string, userID uint32) (bool, error)
	DeletePhone(phone string, userID uint32) (bool, error)
}

type apiPhonebook struct {
	// Dependencies
}

func NewPhonebook() APIPhonebook {
	return &apiPhonebook{}
}

func (a apiPhonebook) GetUserIDByPhone(phone string) (uint32, error) {
	// Implementation
	return 0, nil
}

func (a apiPhonebook) SavePhone(phone string, userID uint32) (bool, error) {
	// Implementation
	return false, nil
}

func (a apiPhonebook) DeletePhone(phone string, userID uint32) (bool, error) {
	// Implementation
	return false, nil
}
