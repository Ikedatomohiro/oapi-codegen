// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add a new pet to the store
	// (POST /pet)
	AddPet(ctx echo.Context) error
	// Update an existing pet
	// (PUT /pet)
	UpdatePet(ctx echo.Context) error
	// Finds Pets by status
	// (GET /pet/findByStatus)
	FindPetsByStatus(ctx echo.Context, params FindPetsByStatusParams) error
	// Finds Pets by tags
	// (GET /pet/findByTags)
	FindPetsByTags(ctx echo.Context, params FindPetsByTagsParams) error
	// Deletes a pet
	// (DELETE /pet/{petId})
	DeletePet(ctx echo.Context, petId int64, params DeletePetParams) error
	// Find pet by ID
	// (GET /pet/{petId})
	GetPetById(ctx echo.Context, petId int64) error
	// Updates a pet in the store with form data
	// (POST /pet/{petId})
	UpdatePetWithForm(ctx echo.Context, petId int64, params UpdatePetWithFormParams) error
	// uploads an image
	// (POST /pet/{petId}/uploadImage)
	UploadFile(ctx echo.Context, petId int64, params UploadFileParams) error
	// Returns pet inventories by status
	// (GET /store/inventory)
	GetInventory(ctx echo.Context) error
	// Place an order for a pet
	// (POST /store/order)
	PlaceOrder(ctx echo.Context) error
	// Delete purchase order by ID
	// (DELETE /store/order/{orderId})
	DeleteOrder(ctx echo.Context, orderId int64) error
	// Find purchase order by ID
	// (GET /store/order/{orderId})
	GetOrderById(ctx echo.Context, orderId int64) error
	// Create user
	// (POST /user)
	CreateUser(ctx echo.Context) error
	// Creates list of users with given input array
	// (POST /user/createWithList)
	CreateUsersWithListInput(ctx echo.Context) error
	// Logs user into the system
	// (GET /user/login)
	LoginUser(ctx echo.Context, params LoginUserParams) error
	// Logs out current logged in user session
	// (GET /user/logout)
	LogoutUser(ctx echo.Context) error
	// Delete user
	// (DELETE /user/{username})
	DeleteUser(ctx echo.Context, username string) error
	// Get user by user name
	// (GET /user/{username})
	GetUserByName(ctx echo.Context, username string) error
	// Update user
	// (PUT /user/{username})
	UpdateUser(ctx echo.Context, username string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// UpdatePet converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePet(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdatePet(ctx)
	return err
}

// FindPetsByStatus converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetsByStatus(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByStatusParams
	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FindPetsByStatus(ctx, params)
	return err
}

// FindPetsByTags converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetsByTags(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByTagsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FindPetsByTags(ctx, params)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeletePetParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "api_key" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("api_key")]; found {
		var ApiKey string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for api_key, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "api_key", runtime.ParamLocationHeader, valueList[0], &ApiKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter api_key: %s", err))
		}

		params.ApiKey = &ApiKey
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePet(ctx, petId, params)
	return err
}

// GetPetById converts echo context to params.
func (w *ServerInterfaceWrapper) GetPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{})

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPetById(ctx, petId)
	return err
}

// UpdatePetWithForm converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePetWithForm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdatePetWithFormParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdatePetWithForm(ctx, petId, params)
	return err
}

// UploadFile converts echo context to params.
func (w *ServerInterfaceWrapper) UploadFile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UploadFileParams
	// ------------- Optional query parameter "additionalMetadata" -------------

	err = runtime.BindQueryParameter("form", true, false, "additionalMetadata", ctx.QueryParams(), &params.AdditionalMetadata)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter additionalMetadata: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UploadFile(ctx, petId, params)
	return err
}

// GetInventory converts echo context to params.
func (w *ServerInterfaceWrapper) GetInventory(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetInventory(ctx)
	return err
}

// PlaceOrder converts echo context to params.
func (w *ServerInterfaceWrapper) PlaceOrder(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PlaceOrder(ctx)
	return err
}

// DeleteOrder converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteOrder(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "orderId" -------------
	var orderId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "orderId", runtime.ParamLocationPath, ctx.Param("orderId"), &orderId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter orderId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteOrder(ctx, orderId)
	return err
}

// GetOrderById converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrderById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "orderId" -------------
	var orderId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "orderId", runtime.ParamLocationPath, ctx.Param("orderId"), &orderId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter orderId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOrderById(ctx, orderId)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// CreateUsersWithListInput converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUsersWithListInput(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUsersWithListInput(ctx)
	return err
}

// LoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) LoginUser(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginUserParams
	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// ------------- Optional query parameter "password" -------------

	err = runtime.BindQueryParameter("form", true, false, "password", ctx.QueryParams(), &params.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter password: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.LoginUser(ctx, params)
	return err
}

// LogoutUser converts echo context to params.
func (w *ServerInterfaceWrapper) LogoutUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.LogoutUser(ctx)
	return err
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUser(ctx, username)
	return err
}

// GetUserByName converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserByName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserByName(ctx, username)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUser(ctx, username)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/pet", wrapper.AddPet)
	router.PUT(baseURL+"/pet", wrapper.UpdatePet)
	router.GET(baseURL+"/pet/findByStatus", wrapper.FindPetsByStatus)
	router.GET(baseURL+"/pet/findByTags", wrapper.FindPetsByTags)
	router.DELETE(baseURL+"/pet/:petId", wrapper.DeletePet)
	router.GET(baseURL+"/pet/:petId", wrapper.GetPetById)
	router.POST(baseURL+"/pet/:petId", wrapper.UpdatePetWithForm)
	router.POST(baseURL+"/pet/:petId/uploadImage", wrapper.UploadFile)
	router.GET(baseURL+"/store/inventory", wrapper.GetInventory)
	router.POST(baseURL+"/store/order", wrapper.PlaceOrder)
	router.DELETE(baseURL+"/store/order/:orderId", wrapper.DeleteOrder)
	router.GET(baseURL+"/store/order/:orderId", wrapper.GetOrderById)
	router.POST(baseURL+"/user", wrapper.CreateUser)
	router.POST(baseURL+"/user/createWithList", wrapper.CreateUsersWithListInput)
	router.GET(baseURL+"/user/login", wrapper.LoginUser)
	router.GET(baseURL+"/user/logout", wrapper.LogoutUser)
	router.DELETE(baseURL+"/user/:username", wrapper.DeleteUser)
	router.GET(baseURL+"/user/:username", wrapper.GetUserByName)
	router.PUT(baseURL+"/user/:username", wrapper.UpdateUser)

}
