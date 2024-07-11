package services

import (
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"Marcketplace/helper"
	"Marcketplace/model"
	"Marcketplace/repository"

	"github.com/go-playground/validator/v10"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		validate:       validate,
	}
}

// Create implements NoteService.
func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) {
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)
	noteModel := model.Note{
		Content: note.Content,
	}
	n.NoteRepository.Save(noteModel)
}

// Delete implements NoteService.
func (n *NoteServiceImpl) Delete(noteId int) {
	n.NoteRepository.Delete(noteId)
}

// FindAll implements NoteService.
func (n *NoteServiceImpl) FindAll() []response.NoteResponse {
	result := n.NoteRepository.FindAll()
	var notes []response.NoteResponse

	for _, value := range result {
		note := response.NoteResponse{
			Id:      value.Id,
			Content: value.Content,
		}
		notes = append(notes, note)
	}
	return notes
}

// FindById implements NoteService.
func (n *NoteServiceImpl) FindById(noteId int) response.NoteResponse {
	noteData, err := n.NoteRepository.FindById(noteId)
	helper.ErrorPanic(err)
	NoteResponse := response.NoteResponse{
		Id:      noteId,
		Content: noteData.Content,
	}
	return NoteResponse
}

// Update implements NoteService.
func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) {
	noteData, err := n.NoteRepository.FindById(note.Id)
	helper.ErrorPanic(err)
	noteData.Content = note.Content
	n.NoteRepository.Update(noteData)
}
