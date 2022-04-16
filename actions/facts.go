package actions

import (
	"facts/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/pkg/errors"
)

//FactsIndex displays all facts
func FactList(c buffalo.Context) error {
	//TODO: pagination
	// // Paginate results. Params "page" and "per_page" control pagination.
	// // Default values are "page=1" and "per_page=20".
	// q := tx.PaginateFromParams(c.Params())
	facts := &[]models.Fact{}
	tx := c.Value("tx").(*pop.Connection)

	if err := tx.All(facts); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(http.StatusOK, r.JSON(facts))
}

// FactsCreate creates a new fact from a POST request
// mapped to /facts.
func FactCreate(c buffalo.Context) error {
	fact := &models.Fact{}
	if err := c.Bind(fact); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(fact)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		return errors.New(verrs.Error())
	}

	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Fact created!", "id": fact.ID.String()}))
}

func FactShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	id := c.Param("id")
	fact := models.Fact{}

	err := tx.Find(&fact, id)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(http.StatusOK, r.JSON(fact))
}

func FactDestroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errors.WithStack(err)
	}

	fact := models.Fact{ID: id}

	err = tx.Destroy(&fact)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Fact destroyed!"}))
}
