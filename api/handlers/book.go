package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookService models.BookService
	Env         *configs.Env
}

func (bh *BookHandler) Create(c echo.Context) error {
	validate := validator.New()
	var payload *models.BookInput
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	userID := c.Get("x-user-id")
	value := fmt.Sprint(userID)
	UID, _ := uuid.Parse(value)
	newBook := models.Book{
		UserID:        UID,
		Title:         payload.Title,
		Author:        payload.Author,
		Description:   payload.Description,
		CoverPhotoUrl: payload.CoverPhotoUrl,
	}

	err = bh.BookService.Create(&newBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := fmt.Sprintf("Inserted ID: %s", newBook.ID)
	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": response})
}

func (bh *BookHandler) FetchAllByUserID(c echo.Context) error {
	newUID := fmt.Sprint(c.Get("x-user-id"))

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	books, err := bh.BookService.FetchAllByUserID(newUID, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.BookResponse, len(books))

	for k, v := range books {
		response[k] = models.BookResponse{
			Title:         v.Title,
			Author:        v.Author,
			Description:   v.Description,
			CoverPhotoUrl: v.CoverPhotoUrl,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(books), "book": response}})
}

func (bh *BookHandler) FetchByID(c echo.Context) error {
	bookId := c.Param("bookId")
	if bookId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	book, err := bh.BookService.FetchByID(bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := models.BookResponse{
		Title:         book.Title,
		Author:        book.Author,
		Description:   book.Description,
		CoverPhotoUrl: book.CoverPhotoUrl,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"book": response}})
}

func (bh *BookHandler) UpdateByID(c echo.Context) error {
	validate := validator.New()
	bookId := c.Param("bookId")
	var payload *models.BookInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	if bookId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	cr := time.Now()
	updatedBook := models.Book{
		Title:         payload.Title,
		Author:        payload.Author,
		Description:   payload.Description,
		CoverPhotoUrl: payload.CoverPhotoUrl,
		UpdatedAt:     &cr,
	}

	err = bh.BookService.UpdateByID(&updatedBook, bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "updated_at": updatedBook.UpdatedAt})
}

func (bh *BookHandler) DeleteByID(c echo.Context) error {
	bookId := c.Param("bookId")
	if bookId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	err := bh.BookService.DeleteByID(bookId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "user successfully deleted"})
}
