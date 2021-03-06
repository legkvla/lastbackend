package mock

import (
	e "github.com/lastbackend/lastbackend/libs/errors"
	"github.com/lastbackend/lastbackend/libs/interface/storage"
	"github.com/lastbackend/lastbackend/libs/model"
	r "gopkg.in/dancannon/gorethink.v2"
)

const HookTable string = "hooks"

// Service Build type for interface in interfaces folder
type HookMock struct {
	Mock *r.Mock
	storage.IHook
}

// Get hooks by image
func (s *HookMock) GetByToken(token string) (*model.Hook, *e.Err) {
	return nil, nil
}

// Get hooks by image
func (s *HookMock) GetByUser(id string) (*model.HookList, *e.Err) {
	return nil, nil
}

// Get hooks by image
func (s *HookMock) GetByImage(user, id string) (*model.HookList, *e.Err) {
	return nil, nil
}

// Get hooks by service
func (s *HookMock) GetByService(user, id string) (*model.HookList, *e.Err) {
	return nil, nil
}

// Insert new hook into storage
func (s *HookMock) Insert(hook *model.Hook) (*model.Hook, *e.Err) {
	return nil, nil
}

// Insert new hook into storage
func (s *HookMock) Delete(user, id string) *e.Err {
	return nil
}

func newHookMock(mock *r.Mock) *HookMock {
	s := new(HookMock)
	s.Mock = mock
	return s
}
