package test

import (
	"Marcketplace/controller"
	"Marcketplace/data/request"
	"Marcketplace/data/response"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNoteController(t *testing.T) {
	app := fiber.New()

	noteServiceMock := new(NoteServiceMock)
	noteController := controller.NewNoteController(noteServiceMock)

	app.Post("/api/notes/create", noteController.Create)

	noteRequest := request.CreateNoteRequest{
		Content: "Test note",
	}

	noteServiceMock.On("Create", mock.AnythingOfType("request.CreateNoteRequest")).Return(nil)

	body, _ := json.Marshal(noteRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/notes/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var webResponse response.Response
	json.NewDecoder(resp.Body).Decode(&webResponse)
	assert.Equal(t, "ok", webResponse.Status)
	assert.Equal(t, "Successfully created notes data!", webResponse.Message)

	noteServiceMock.AssertExpectations(t)
}
