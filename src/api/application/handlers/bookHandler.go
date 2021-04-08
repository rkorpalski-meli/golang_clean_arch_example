package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/application/presenter"
	"github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
	"net/http"
)

type BookHandler struct {
	bookUsecase	entity.BookUsecase
}

func NewBookHandler(bookUsecase entity.BookUsecase) BookHandler {
	return BookHandler{bookUsecase: bookUsecase}
}

func (h BookHandler) FindBook(c *gin.Context) {
	title := c.Param("title")
	book, err := h.bookUsecase.FindBook(title)
	if err != nil  {
		c.AbortWithError(500, err)
		return
	}

	bookJson, _ := json.Marshal(book)
	c.String(200, string(bookJson))
}

func (h BookHandler) InsertBook(c *gin.Context) {
	var insertBookDTO presenter.InsertBookDTO
	if err := c.ShouldBindJSON(&insertBookDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.bookUsecase.CreateBook(insertBookDTO.Title, insertBookDTO.Author)
	if err != nil  {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "")
}