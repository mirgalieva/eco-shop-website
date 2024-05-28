package service

import (
	"golang.org/x/crypto/bcrypt"
	entity2 "mirgalievaal-project/backend/internal/entity"
)

type UserService struct {
	userRepo entity2.UserRepository
}

func (s *UserService) GetAll() (*[]entity2.User, error) {
	userDB, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		for _, user := range *userDB {
			user.OmitPassword()
		}
		return userDB, nil
	}
}

func NewUserService(userRepo entity2.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Get(id uint) (*entity2.User, error) {
	userDB, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		userDB.OmitPassword()
		return userDB, nil
	}
}
func (s *UserService) Update(user *entity2.User) error {
	userDB, err := s.userRepo.Get(user.ID)
	if err != nil {
		return err
	}
	var newUserPasswordHash string
	if comparePasswordWithHash(user.Password, userDB.Password) != nil {
		newUserPasswordHash, err = generatePasswordHash(user.Password)
		if err != nil {
			return err
		}
	}
	user.Password = newUserPasswordHash
	err = s.userRepo.Update(user)
	return err
}

func comparePasswordWithHash(passwdFromInput, passwdHashFromDB string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwdHashFromDB), []byte(passwdFromInput))
	return err
}

func generatePasswordHash(passwd string) (string, error) {
	passwdHash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(passwdHash), err
}

func (s *UserService) Delete(user *entity2.User) error {
	err := s.userRepo.Delete(user.ID)
	return err
}

func (s *UserService) Register(userReg *entity2.UserRegister) error {
	if userReg.Email == "" {
		return entity2.ErrInvalidEmail
	}
	if len(userReg.Password) < 8 {
		return entity2.ErrInvalidPassword
	}
	passwdHash, err := generatePasswordHash(userReg.Password)
	if err != nil {
		return err
	}
	userReg.Password = passwdHash

	user := entity2.User{UserRegister: *userReg}

	err = s.userRepo.Create(&user)
	return err
}
func (s *UserService) Login(userLogin *entity2.UserLogin) error {
	userDB, err := s.userRepo.GetByEmail(userLogin.Email)
	if err != nil {
		return err
	}
	err = comparePasswordWithHash(userLogin.Password, userDB.Password)
	if err != nil {
		return err
	}
	return nil
}
