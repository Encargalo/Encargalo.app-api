package products

import (
	"CaliYa/core/domain/dto"
	productsPort "CaliYa/core/domain/ports/products"
	calierrors "CaliYa/core/errors"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Products interface {
	RegisterProducts(c echo.Context) error
	GetProductsByCategory(c echo.Context) error
	GetAdicionesByCategory(c echo.Context) error
}

type products struct {
	app productsPort.ProductsApp
}

func NewProducts(app productsPort.ProductsApp) Products {
	return &products{app}
}

func (p *products) RegisterProducts(c echo.Context) error {

	ctx := c.Request().Context()

	err := p.app.RegisterProduct(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// GetProductByCategory godoc
// @Tags Products
// @Summary Se obtienen todos los productos de una misma categoria.
// @Description Se obtiene una lista de productos filtradas por el nombre de una categoria, tambien puede ser una similitud, ej:Si se busca la palabra hamb, obtendrá hamburguesas o otra categoría similar.
// @Produce json
// @Param category query string true "Este es el nombre de la categoria ej:/products/category/?category=hamburguesas"
// @Success 200 {object} []items.Items
// @Failure 400 {string} string "Se retorna cuando cuando el valor es vacio o el valor es menor a 3 digitos."
// @Failure 404 {string} string "Se retorna cuando no se encuentra una concidencia en la busqueda."
// @Failure 500 {string} string "Se retorna cuando ocurre un error inexperado en el servidor."
// @Router /products/category [get]
func (p *products) GetProductsByCategory(c echo.Context) error {

	ctx := c.Request().Context()

	category := dto.SearchProductsByCategory{}

	if err := c.Bind(&category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if err := category.IsValid(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	items, err := p.app.GetProductByCategory(ctx, category.Category)
	if err != nil {
		switch {
		case errors.Is(err, calierrors.ErrNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, items)
}

// GetAdicionesGyCategory godoc
// @Summary Se obtienen las adiciones por categoría.
// @Description Devuelve la lista de adiciones filtradas por el ID de la categoría a la que pertenece un producto.
// @Tags Products
// @Produce json
// @Param category_id query string true "ID de la categoría (UUID)"
// @Success 200 {object} []items.Items
// @Failure 400 {string} string "Se retorna cuando el param se envía vacío o no es un UUID Valido."
// @Failure 404 {string} string "Se retorna cuando no se encuentran ninguna adición para el category_id enviada."
// @Failure 500 {string} string "Se retorna cuando ocurre un error inesperado dentro del servidor."
// @Router /products/adiciones [get]
func (p *products) GetAdicionesByCategory(c echo.Context) error {

	ctx := c.Request().Context()

	categorys := dto.Category{}

	if err := c.Bind(&categorys); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	adiciones, err := p.app.GetAditionsByCategory(ctx, categorys.ID)
	if err != nil {
		switch {
		case errors.Is(err, calierrors.ErrNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, adiciones)
}
