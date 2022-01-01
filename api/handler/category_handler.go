package handler

import (
	"net/http"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) CategoryHandler {
	return CategoryHandler{
		categoryUsecase: categoryUsecase,
	}
}

func (handler *CategoryHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var category model.Category

		if err := c.Bind(&category); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusUnprocessableEntity, category)
		}
		if err := handler.categoryUsecase.Create(&category); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func (handler *CategoryHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := handler.categoryUsecase.List()
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, categories)
	}
}
