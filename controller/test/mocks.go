package test

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/model/entities"

	"github.com/stretchr/testify/mock"
)

type NoteServiceMock struct {
	mock.Mock
}

func (m *NoteServiceMock) Create(note request.CreateNoteRequest) {
	m.Called(note)
}

func (m *NoteServiceMock) Update(note request.UpdateNoteRequest) {
	m.Called(note)
}

func (m *NoteServiceMock) Delete(noteId int) {
	m.Called(noteId)
}

func (m *NoteServiceMock) FindById(noteId int) response.NoteResponse {
	args := m.Called(noteId)
	return args.Get(0).(response.NoteResponse)
}

func (m *NoteServiceMock) FindAll() []response.NoteResponse {
	args := m.Called()
	return args.Get(0).([]response.NoteResponse)
}

type UserServiceMock struct {
	mock.Mock
}

// AuthenticateUser implements services.UserService.
func (m *UserServiceMock) AuthenticateUser(email string, password string) (bool, *entities.User, error) {
	panic("unimplemented")
}

// CreateNFA implements services.UserService.
func (m *UserServiceMock) CreateNFA(nfa *entities.NFA) error {
	panic("unimplemented")
}

// Delete implements services.UserService.
func (m *UserServiceMock) Delete(userId int) {
	panic("unimplemented")
}

// FindAll implements services.UserService.
func (m *UserServiceMock) FindAll() []response.UserResponse {
	panic("unimplemented")
}

// FindById implements services.UserService.
func (m *UserServiceMock) FindById(userId uint) response.UserResponse {
	panic("unimplemented")
}

// FindNFA implements services.UserService.
func (m *UserServiceMock) FindNFA(nfa *uint) (*entities.NFA, error) {
	panic("unimplemented")
}

// FindUser implements services.UserService.
func (m *UserServiceMock) FindUser(userId int) *entities.User {
	panic("unimplemented")
}

// Update implements services.UserService.
func (m *UserServiceMock) Update(user request.UpdateUserRequest) {
	panic("unimplemented")
}

func (m *UserServiceMock) Create(user request.CreateUserRequest) {
	m.Called(user)
}
