package products

import (
	productsService "CaliYa/core/domain/ports/products"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Flavors interface {
	SearchFlavorsBy(c echo.Context) error
}

type flavors struct {
	svc productsService.FlavorApp
}

func NewFlavorsHandlers(svc productsService.FlavorApp) Flavors {
	return &flavors{svc: svc}
}

// SearchFlavorsBy godoc
// @Summary      Get flavors by product ID
// @Description  Returns the list of flavors associated with a given product
// @Tags         Flavors
// @Accept       json
// @Produce      json
// @Param        product_id  query     string  true  "Product ID (UUID)"
// @Success      200  {array}   items.Flavor
// @Failure      400  {string}  "Invalid product_id or missing parameter"
// @Failure      404  {string}  "No flavors found for the given product"
// @Failure      500  {string}  "Unexpected internal server error"
// @Router       /products/flavors [get]
func (f *flavors) SearchFlavorsBy(c echo.Context) error {
	ctx := c.Request().Context()

	productIDParam := c.QueryParam("product_id")
	if productIDParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "the 'product_id' parameter is required")
	}

	productID, err := uuid.Parse(productIDParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "'product_id' must be a valid UUID")
	}

	results, err := f.svc.SearchFlavorsByProductID(ctx, productID)
	if err != nil {
		switch err.Error() {
		case "not found":
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
		}
	}

	return c.JSON(http.StatusOK, results)
}
