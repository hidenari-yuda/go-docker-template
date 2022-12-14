package handler

import (
	"fmt"

	"github.com/hidenari-yuda/go-docker-template/domain/entity"
	"github.com/hidenari-yuda/go-docker-template/domain/entity/responses"
	"github.com/hidenari-yuda/go-docker-template/interfaces/presenter"

	"github.com/hidenari-yuda/go-docker-template/usecase/interactor"
)

type UserHandler interface {
	// Gest API
	SignUp(param *entity.SignUpParam) (presenter.Presenter, error)
	SignIn(param *entity.SignInParam) (presenter.Presenter, error)
}

type UserHandlerImpl struct {
	UserInteractor interactor.UserInteractor
}

func NewUserHandlerImpl(ui interactor.UserInteractor) UserHandler {
	return &UserHandlerImpl{
		UserInteractor: ui,
	}
}

func (h *UserHandlerImpl) SignUp(param *entity.SignUpParam) (presenter.Presenter, error) {

	output, err := h.UserInteractor.SignUp(interactor.SignUpInput{
		Param: param,
	})
	fmt.Println(output, err)

	if err != nil {
		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
		return nil, err
	}

	return presenter.NewOkJSONPresenter(responses.NewOK(output.Ok)), nil

}

func (h *UserHandlerImpl) SignIn(param *entity.SignInParam) (presenter.Presenter, error) {
	output, err := h.UserInteractor.SignIn(interactor.SignInInput{
		Param: param,
	})
	if err != nil {
		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
		return nil, err
	}

	return presenter.NewUserJSONPresenter(responses.NewUser(output.User)), nil

}
