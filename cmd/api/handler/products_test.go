package handler_test

import (
	"CaliYa/cmd/api/handler"
	"CaliYa/core/domain/dto"
	"CaliYa/core/domain/models"
	mocks "CaliYa/mocks/CaliYa/core/domain/ports"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"

	calierrors "CaliYa/core/errors"
)

var (
	ctx = context.Background()

	queryCategoryIsEmpty = dto.SearchProductsByCategory{
		Category: "",
	}

	queryCategoryIsIsLessThanThree = dto.SearchProductsByCategory{
		Category: "as",
	}

	queryCategoryIsValid = dto.SearchProductsByCategory{
		Category: "7bcceb44-fa9c-4451-8b5b-ab36f90fa95d",
	}
)

func TestProductsSuite(t *testing.T) {
	suite.Run(t, new(ProductsSuiteTest))
}

type ProductsSuiteTest struct {
	suite.Suite
	service   *mocks.ProductsApp
	underTest handler.Products
}

func (suite *ProductsSuiteTest) SetupTest() {
	suite.service = mocks.NewProductsApp(suite.T())
	suite.underTest = handler.NewProducts(suite.service)
}

func (suite *ProductsSuiteTest) TestGetProductsByCategory_WhenBindFail() {

	body, _ := json.Marshal("")
	controller := SetupControllerCase("/api/products/category?category=", http.MethodGet, bytes.NewBuffer(body))

	err := suite.underTest.GetProductsByCategory(controller.Ctx)
	suite.Contains(err.Error(), "400")
	suite.Error(err)

}

func (suite *ProductsSuiteTest) TestGetProductsByCategory_WhenTheParamIsEmpty() {

	body, _ := json.Marshal(queryCategoryIsEmpty)
	controller := SetupControllerCase("/api/products/category?category=", http.MethodGet, bytes.NewBuffer(body))

	err := suite.underTest.GetProductsByCategory(controller.Ctx)

	httpErr, ok := err.(*echo.HTTPError)

	suite.True(ok, "el error debería ser tipo *echo.HTTPError")

	suite.Equal(http.StatusBadRequest, httpErr.Code)
	suite.Equal("the value cannot be empty", httpErr.Message)

}

func (suite *ProductsSuiteTest) TestGetProductsByCategory_WhenTheParameterIsLessThanThree() {

	body, _ := json.Marshal(queryCategoryIsIsLessThanThree)
	controller := SetupControllerCase("/api/products/category?category=", http.MethodGet, bytes.NewBuffer(body))

	err := suite.underTest.GetProductsByCategory(controller.Ctx)

	httpErr, ok := err.(*echo.HTTPError)

	suite.True(ok, "el error debería ser tipo *echo.HTTPError")

	suite.Equal(http.StatusBadRequest, httpErr.Code)
	suite.Equal("the value must be at least 3 characters long", httpErr.Message)

}

func (suite *ProductsSuiteTest) TestGetProductsByCategory_WhenProductsNotFound() {

	body, _ := json.Marshal(queryCategoryIsValid)
	controller := SetupControllerCase("/api/products/category?category=", http.MethodGet, bytes.NewBuffer(body))

	suite.service.Mock.On("GetProductByCategory", ctx, queryCategoryIsValid.Category).
		Return([]models.Items{}, calierrors.ErrNotFound)

	err := suite.underTest.GetProductsByCategory(controller.Ctx)
	httpErr, ok := err.(*echo.HTTPError)

	suite.True(ok, "el error debería ser tipo *echo.HTTPError")

	suite.Equal(http.StatusNotFound, httpErr.Code)
	suite.Equal("not found.", httpErr.Message)

}

func (suite *ProductsSuiteTest) TestGetProductsByCategory_WhenInternalServerError() {

	body, _ := json.Marshal(queryCategoryIsValid)
	controller := SetupControllerCase("/api/products/category?category=", http.MethodGet, bytes.NewBuffer(body))

	suite.service.Mock.On("GetProductByCategory", ctx, queryCategoryIsValid.Category).
		Return([]models.Items{}, calierrors.ErrUnexpected)

	err := suite.underTest.GetProductsByCategory(controller.Ctx)
	httpErr, ok := err.(*echo.HTTPError)

	suite.True(ok, "el error debería ser tipo *echo.HTTPError")

	suite.Equal(http.StatusInternalServerError, httpErr.Code)
	suite.Equal("unexpected error.", httpErr.Message)

}

func (suite *ProductsSuiteTest) TestGetProductsByCategory_WhenSuccess() {

	body, _ := json.Marshal(queryCategoryIsValid)
	controller := SetupControllerCase("/api/products/category?category=", http.MethodGet, bytes.NewBuffer(body))

	suite.service.Mock.On("GetProductByCategory", ctx, queryCategoryIsValid.Category).
		Return([]models.Items{}, nil)

	err := suite.underTest.GetProductsByCategory(controller.Ctx)
	suite.NoError(err)

	suite.Equal(http.StatusOK, controller.Res.Code)
}
