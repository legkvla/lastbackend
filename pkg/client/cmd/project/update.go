package project

import (
	"errors"
	e "github.com/lastbackend/lastbackend/libs/errors"
	"github.com/lastbackend/lastbackend/libs/model"
	"github.com/lastbackend/lastbackend/pkg/client/context"
)

type updateS struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

func UpdateCmd(name, description string) {

	var ctx = context.Get()

	err := Update(name, description)
	if err != nil {
		ctx.Log.Error(err)
		return
	}

	ctx.Log.Info("Successful")
}

func Update(name, description string) error {

	var (
		err error
		ctx = context.Get()
		er  = new(e.Http)
		res = new(model.Project)
	)

	if len(name) == 0 {
		return e.BadParameter("name").Err()
	}

	_, _, err = ctx.HTTP.
		PUT("/project").
		AddHeader("Content-Type", "application/json").
		AddHeader("Authorization", "Bearer "+ctx.Token).
		BodyJSON(updateS{name, description}).
		Request(&res, er)
	if err != nil {
		return err
	}

	if er.Code == 401 {
		return errors.New("You are currently not logged in to the system, to get proper access create a new user or login with an existing user.")
	}

	if er.Code != 0 {
		return errors.New(e.Message(er.Status))
	}

	return nil
}
