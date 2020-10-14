package restapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaindrag/goelasticsearch"
)

type restapi struct {
	repository goelasticsearch.Repository
}

func (r restapi) fetch(c echo.Context) error {
	documents, err := r.repository.Fetch(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, documents)
}

func (r restapi) store(c echo.Context) error {
	document := make(map[string]interface{})

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := json.Unmarshal(body, &document); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := r.repository.Store(c.Request().Context(), document); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, document)
}

func (r restapi) update(c echo.Context) error {
	document := make(map[string]interface{})

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := json.Unmarshal(body, &document); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := r.repository.Update(c.Request().Context(), c.Param("id"), document); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, document)
}

func (r restapi) getByID(c echo.Context) error {
	document, err := r.repository.GetByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, document)
}

func (r restapi) delete(c echo.Context) error {
	if err := r.repository.Delete(c.Request().Context(), c.Param("id")); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func NewAPI(e *echo.Echo, repository goelasticsearch.Repository) {
	r := restapi{repository: repository}

	e.GET("/document", r.fetch)
	e.POST("/document", r.store)
	e.PUT("/document/:id", r.update)
	e.DELETE("/document/:id", r.delete)
	e.GET("/document/:id", r.getByID)
}
