// Package logins provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package logins

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CreateLoginRequest defines model for CreateLoginRequest.
type CreateLoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Login defines model for Login.
type Login struct {
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	ID               *string    `json:"id,omitempty"`
	PasswordVersion  *int       `json:"password_version,omitempty"`
	TwoFactorEnabled *bool      `json:"two_factor_enabled,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	Username         *string    `json:"username,omitempty"`
}

// TwoFactorMethod defines model for TwoFactorMethod.
type TwoFactorMethod struct {
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`
}

// UpdateLoginRequest defines model for UpdateLoginRequest.
type UpdateLoginRequest struct {
	Username string `json:"username"`
}

// UpdatePasswordRequest defines model for UpdatePasswordRequest.
type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

// GetParams defines parameters for Get.
type GetParams struct {
	// Maximum number of logins to return
	Limit *int `json:"limit,omitempty"`

	// Offset for pagination
	Offset *int `json:"offset,omitempty"`

	// Search term for filtering logins
	Search *string `json:"search,omitempty"`
}

// PostJSONBody defines parameters for Post.
type PostJSONBody CreateLoginRequest

// PutIDJSONBody defines parameters for PutID.
type PutIDJSONBody UpdateLoginRequest

// PutIDPasswordJSONBody defines parameters for PutIDPassword.
type PutIDPasswordJSONBody UpdatePasswordRequest

// PostJSONRequestBody defines body for Post for application/json ContentType.
type PostJSONRequestBody PostJSONBody

// Bind implements render.Binder.
func (PostJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PutIDJSONRequestBody defines body for PutID for application/json ContentType.
type PutIDJSONRequestBody PutIDJSONBody

// Bind implements render.Binder.
func (PutIDJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PutIDPasswordJSONRequestBody defines body for PutIDPassword for application/json ContentType.
type PutIDPasswordJSONRequestBody PutIDPasswordJSONBody

// Bind implements render.Binder.
func (PutIDPasswordJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// GetJSON200Response is a constructor method for a Get response.
// A *Response is returned with the configured status code and content type from the spec.
func GetJSON200Response(body struct {
	Logins []Login `json:"logins,omitempty"`

	// Total number of logins
	Total *int `json:"total,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostJSON201Response is a constructor method for a Post response.
// A *Response is returned with the configured status code and content type from the spec.
func PostJSON201Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        201,
		contentType: "application/json",
	}
}

// PostJSON400Response is a constructor method for a Post response.
// A *Response is returned with the configured status code and content type from the spec.
func PostJSON400Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// DeleteIDJSON404Response is a constructor method for a DeleteID response.
// A *Response is returned with the configured status code and content type from the spec.
func DeleteIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// GetIDJSON200Response is a constructor method for a GetID response.
// A *Response is returned with the configured status code and content type from the spec.
func GetIDJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetIDJSON404Response is a constructor method for a GetID response.
// A *Response is returned with the configured status code and content type from the spec.
func GetIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PutIDJSON200Response is a constructor method for a PutID response.
// A *Response is returned with the configured status code and content type from the spec.
func PutIDJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PutIDJSON404Response is a constructor method for a PutID response.
// A *Response is returned with the configured status code and content type from the spec.
func PutIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// Get2faMethodsByLoginIDJSON200Response is a constructor method for a Get2faMethodsByLoginID response.
// A *Response is returned with the configured status code and content type from the spec.
func Get2faMethodsByLoginIDJSON200Response(body []TwoFactorMethod) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// Get2faMethodsByLoginIDJSON404Response is a constructor method for a Get2faMethodsByLoginID response.
// A *Response is returned with the configured status code and content type from the spec.
func Get2faMethodsByLoginIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PutIDPasswordJSON200Response is a constructor method for a PutIDPassword response.
// A *Response is returned with the configured status code and content type from the spec.
func PutIDPasswordJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PutIDPasswordJSON404Response is a constructor method for a PutIDPassword response.
// A *Response is returned with the configured status code and content type from the spec.
func PutIDPasswordJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all logins
	// (GET /)
	Get(w http.ResponseWriter, r *http.Request, params GetParams) *Response
	// Create a new login
	// (POST /)
	Post(w http.ResponseWriter, r *http.Request) *Response
	// Delete login
	// (DELETE /{id})
	DeleteID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Get login by ID
	// (GET /{id})
	GetID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Update login
	// (PUT /{id})
	PutID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Get login 2FA methods
	// (GET /{id}/2fa)
	Get2faMethodsByLoginID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Update login password
	// (PUT /{id}/password)
	PutIDPassword(w http.ResponseWriter, r *http.Request, id string) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// Get operation middleware
func (siw *ServerInterfaceWrapper) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetParams

	// ------------- Optional query parameter "limit" -------------

	if err := runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit); err != nil {
		err = fmt.Errorf("invalid format for parameter limit: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "limit"})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	if err := runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset); err != nil {
		err = fmt.Errorf("invalid format for parameter offset: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "offset"})
		return
	}

	// ------------- Optional query parameter "search" -------------

	if err := runtime.BindQueryParameter("form", true, false, "search", r.URL.Query(), &params.Search); err != nil {
		err = fmt.Errorf("invalid format for parameter search: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "search"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Get(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// Post operation middleware
func (siw *ServerInterfaceWrapper) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Post(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// DeleteID operation middleware
func (siw *ServerInterfaceWrapper) DeleteID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeleteID(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetID operation middleware
func (siw *ServerInterfaceWrapper) GetID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetID(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// PutID operation middleware
func (siw *ServerInterfaceWrapper) PutID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PutID(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// Get2faMethodsByLoginID operation middleware
func (siw *ServerInterfaceWrapper) Get2faMethodsByLoginID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Get2faMethodsByLoginID(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// PutIDPassword operation middleware
func (siw *ServerInterfaceWrapper) PutIDPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PutIDPassword(w, r, id)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:    "/",
		BaseRouter: chi.NewRouter(),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/", wrapper.Get)
		r.Post("/", wrapper.Post)
		r.Delete("/{id}", wrapper.DeleteID)
		r.Get("/{id}", wrapper.GetID)
		r.Put("/{id}", wrapper.PutID)
		r.Get("/{id}/2fa", wrapper.Get2faMethodsByLoginID)
		r.Put("/{id}/password", wrapper.PutIDPassword)
	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xXTW/bOBD9K8TsHpXa9eakW7tBCwMNNthtT0URjKWRw4IiFX4kMQz99wWH8qdkN80G",
	"3iaXQCHHM6P33jyKSyhM3RhN2jvIl+CKG6qRH/+0hJ4+mbnUf9NtIOfjamNNQ9ZL4pgGnbs3tozPftEQ",
	"5OC8lXoObQbBkdVY08Bmm4Gl2yAtlZB/3URmm4zfstWPzOw7FT5m5Gb6XRTcaXmN3GFlbB2foERPZ15y",
	"2l5zstyJDUGWQ2Grdq7vyDpp9Na7SO1pTjZG+XtzXWHhjb0mjTNF24DMjFGEmhFpyp9u9DiKPYg+35sP",
	"3Mkl+RtT9sE62mBa+RFdvJutM/WYyuDhbG7OTOOl0ajO7lAFgtzbQG0GXxiE47p6gnSG9JJKXXUcPkXF",
	"e9WOqDOGSl0ZTiK9inv8kk5cosY51aS9eHc1hQzWYoK3b8ZvxrFV05DGRkIOf/BSlJ6/4fZG8c+cuPHY",
	"NkZYpyXk8JE8B1qsyZN1kH9dQkmusJKxhxwu8UHWoRY61DOywlRCpaa8EZZ8sBpi35DDbSC7gAwS7qBk",
	"LWP25AixdkkVBuUhn4yz3hC02X7lv6rKkReVsaLBudTc9oFihmOHqz2q2D+EtrgRnmzNFSupPEUKu9c9",
	"UNfxz3bq7gvgW1SAa4x2SSyT8Zg9x2hPmjnBplGy4NcbfXfJIzb5dqXWdZMvQXqq+eF3SxXk8Nto48Sj",
	"zoZHyfA2c47W4oL/Nx5Vgmkbhs9xucc1DCE4pN/dZJ+k81tJYoALdY12sdpEpda7GTTGDWj0Kq6mKSLn",
	"35ty8VPwHUNn4IRqdye2s5w9At8+WwcdPwPYxQ3RHUzChaIg56qgFLN3/p80VJNzOGdvpAesG7aaL50P",
	"ClSWsFwIepDOu/6J8hjmp/oOlSyFXaO6TX2CXaDQdJ/o54DRUpZtkqQiT30lXPD69KJvWTya0e82k8nH",
	"8S6R21P6g4N7YGrP+9OSSErtDpF0/twkpYLaRFcMunwaOftJdslJIK9oyQ6eG/8XDeNTzV5JHqVyL5fJ",
	"j+QTjWK2ENMLdtgwZLDhlGQ+v4sPfA8+ysVPpqTuq/2VGESCe9+3R5MKj31mTipMFwr3fsH5X4Z/POoj",
	"a//K1Pvc6kM8+fBO1BydMD6dNiYVdpXd8+hkOOEhK9q8uNsSz/YF6rBHre5hL96r9i+Uv4pdrfp6xY4l",
	"1lJr27b9NwAA//9UH4KbvhIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
