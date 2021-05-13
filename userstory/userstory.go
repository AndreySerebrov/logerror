package userstory

import "err_test/apperrors"

type GetPutUpdater interface {
	Get() error
	Put() error
	Update() error
}

type UserStort struct {
	clietn GetPutUpdater
}

func NewUserStory(clietn GetPutUpdater) UserStort {
	return UserStort{clietn: clietn}
}

func (s UserStort) Do() error {
	err := s.clietn.Put()
	if err != nil {
		//  Оборачиваем ошибки уровня клиента ошибками уровня приложения
		return apperrors.NewImportantError(err)
	}
	err = s.clietn.Update()
	if err != nil {
		return apperrors.NewImportantError(err)
	}
	err = s.clietn.Get()
	if err != nil {
		return apperrors.NewVeryImportantError(err)
	}
	return nil
}
