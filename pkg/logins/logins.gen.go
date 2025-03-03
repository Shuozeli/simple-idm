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

// EnableTwoFactorRequest defines model for EnableTwoFactorRequest.
type EnableTwoFactorRequest struct {
	Secret           string `json:"secret"`
	VerificationCode string `json:"verification_code"`
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

// UpdateLoginRequest defines model for UpdateLoginRequest.
type UpdateLoginRequest struct {
	Username string `json:"username"`
}

// UpdatePasswordRequest defines model for UpdatePasswordRequest.
type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

// GetLoginsParams defines parameters for GetLogins.
type GetLoginsParams struct {
	// Maximum number of logins to return
	Limit *int `json:"limit,omitempty"`

	// Offset for pagination
	Offset *int `json:"offset,omitempty"`

	// Search term for filtering logins
	Search *string `json:"search,omitempty"`
}

// PostLoginsJSONBody defines parameters for PostLogins.
type PostLoginsJSONBody CreateLoginRequest

// PutLoginsIDJSONBody defines parameters for PutLoginsID.
type PutLoginsIDJSONBody UpdateLoginRequest

// PostLoginsID2faEnableJSONBody defines parameters for PostLoginsId2faEnable.
type PostLoginsID2faEnableJSONBody EnableTwoFactorRequest

// PutLoginsIDPasswordJSONBody defines parameters for PutLoginsIDPassword.
type PutLoginsIDPasswordJSONBody UpdatePasswordRequest

// PostLoginsJSONRequestBody defines body for PostLogins for application/json ContentType.
type PostLoginsJSONRequestBody PostLoginsJSONBody

// Bind implements render.Binder.
func (PostLoginsJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PutLoginsIDJSONRequestBody defines body for PutLoginsID for application/json ContentType.
type PutLoginsIDJSONRequestBody PutLoginsIDJSONBody

// Bind implements render.Binder.
func (PutLoginsIDJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostLoginsId2faEnableJSONRequestBody defines body for PostLoginsId2faEnable for application/json ContentType.
type PostLoginsId2faEnableJSONRequestBody PostLoginsID2faEnableJSONBody

// Bind implements render.Binder.
func (PostLoginsId2faEnableJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PutLoginsIDPasswordJSONRequestBody defines body for PutLoginsIDPassword for application/json ContentType.
type PutLoginsIDPasswordJSONRequestBody PutLoginsIDPasswordJSONBody

// Bind implements render.Binder.
func (PutLoginsIDPasswordJSONRequestBody) Bind(*http.Request) error {
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

// GetLoginsJSON200Response is a constructor method for a GetLogins response.
// A *Response is returned with the configured status code and content type from the spec.
func GetLoginsJSON200Response(body struct {
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

// PostLoginsJSON201Response is a constructor method for a PostLogins response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsJSON201Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        201,
		contentType: "application/json",
	}
}

// PostLoginsJSON400Response is a constructor method for a PostLogins response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsJSON400Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// DeleteLoginsIDJSON404Response is a constructor method for a DeleteLoginsID response.
// A *Response is returned with the configured status code and content type from the spec.
func DeleteLoginsIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// GetLoginsIDJSON200Response is a constructor method for a GetLoginsID response.
// A *Response is returned with the configured status code and content type from the spec.
func GetLoginsIDJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetLoginsIDJSON404Response is a constructor method for a GetLoginsID response.
// A *Response is returned with the configured status code and content type from the spec.
func GetLoginsIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PutLoginsIDJSON200Response is a constructor method for a PutLoginsID response.
// A *Response is returned with the configured status code and content type from the spec.
func PutLoginsIDJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PutLoginsIDJSON404Response is a constructor method for a PutLoginsID response.
// A *Response is returned with the configured status code and content type from the spec.
func PutLoginsIDJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PostLoginsId2faDisableJSON200Response is a constructor method for a PostLoginsId2faDisable response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsId2faDisableJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostLoginsId2faDisableJSON404Response is a constructor method for a PostLoginsId2faDisable response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsId2faDisableJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PostLoginsId2faEnableJSON200Response is a constructor method for a PostLoginsId2faEnable response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsId2faEnableJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostLoginsId2faEnableJSON404Response is a constructor method for a PostLoginsId2faEnable response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsId2faEnableJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PostLoginsIDBackupCodesJSON200Response is a constructor method for a PostLoginsIDBackupCodes response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsIDBackupCodesJSON200Response(body struct {
	BackupCodes []string `json:"backup_codes,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostLoginsIDBackupCodesJSON404Response is a constructor method for a PostLoginsIDBackupCodes response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginsIDBackupCodesJSON404Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        404,
		contentType: "application/json",
	}
}

// PutLoginsIDPasswordJSON200Response is a constructor method for a PutLoginsIDPassword response.
// A *Response is returned with the configured status code and content type from the spec.
func PutLoginsIDPasswordJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PutLoginsIDPasswordJSON404Response is a constructor method for a PutLoginsIDPassword response.
// A *Response is returned with the configured status code and content type from the spec.
func PutLoginsIDPasswordJSON404Response(body struct {
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
	// (GET /logins)
	GetLogins(w http.ResponseWriter, r *http.Request, params GetLoginsParams) *Response
	// Create a new login
	// (POST /logins)
	PostLogins(w http.ResponseWriter, r *http.Request) *Response
	// Delete login
	// (DELETE /logins/{id})
	DeleteLoginsID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Get login by ID
	// (GET /logins/{id})
	GetLoginsID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Update login
	// (PUT /logins/{id})
	PutLoginsID(w http.ResponseWriter, r *http.Request, id string) *Response
	// Disable two-factor authentication
	// (POST /logins/{id}/2fa/disable)
	PostLoginsId2faDisable(w http.ResponseWriter, r *http.Request, id string) *Response
	// Enable two-factor authentication
	// (POST /logins/{id}/2fa/enable)
	PostLoginsId2faEnable(w http.ResponseWriter, r *http.Request, id string) *Response
	// Generate new backup codes
	// (POST /logins/{id}/backup-codes)
	PostLoginsIDBackupCodes(w http.ResponseWriter, r *http.Request, id string) *Response
	// Update login password
	// (PUT /logins/{id}/password)
	PutLoginsIDPassword(w http.ResponseWriter, r *http.Request, id string) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// GetLogins operation middleware
func (siw *ServerInterfaceWrapper) GetLogins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLoginsParams

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
		resp := siw.Handler.GetLogins(w, r, params)
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

// PostLogins operation middleware
func (siw *ServerInterfaceWrapper) PostLogins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostLogins(w, r)
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

// DeleteLoginsID operation middleware
func (siw *ServerInterfaceWrapper) DeleteLoginsID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeleteLoginsID(w, r, id)
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

// GetLoginsID operation middleware
func (siw *ServerInterfaceWrapper) GetLoginsID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetLoginsID(w, r, id)
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

// PutLoginsID operation middleware
func (siw *ServerInterfaceWrapper) PutLoginsID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PutLoginsID(w, r, id)
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

// PostLoginsId2faDisable operation middleware
func (siw *ServerInterfaceWrapper) PostLoginsId2faDisable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostLoginsId2faDisable(w, r, id)
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

// PostLoginsId2faEnable operation middleware
func (siw *ServerInterfaceWrapper) PostLoginsId2faEnable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostLoginsId2faEnable(w, r, id)
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

// PostLoginsIDBackupCodes operation middleware
func (siw *ServerInterfaceWrapper) PostLoginsIDBackupCodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostLoginsIDBackupCodes(w, r, id)
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

// PutLoginsIDPassword operation middleware
func (siw *ServerInterfaceWrapper) PutLoginsIDPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "id" -------------
	var id string

	if err := runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "id"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PutLoginsIDPassword(w, r, id)
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
		r.Get("/logins", wrapper.GetLogins)
		r.Post("/logins", wrapper.PostLogins)
		r.Delete("/logins/{id}", wrapper.DeleteLoginsID)
		r.Get("/logins/{id}", wrapper.GetLoginsID)
		r.Put("/logins/{id}", wrapper.PutLoginsID)
		r.Post("/logins/{id}/2fa/disable", wrapper.PostLoginsId2faDisable)
		r.Post("/logins/{id}/2fa/enable", wrapper.PostLoginsId2faEnable)
		r.Post("/logins/{id}/backup-codes", wrapper.PostLoginsIDBackupCodes)
		r.Put("/logins/{id}/password", wrapper.PutLoginsIDPassword)
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

	"H4sIAAAAAAAC/9xYTW/jNhD9KwTbo7L2pjnp1l23gYFdNGizp0Vg0OLI4VYiFXIYxwj03wsO5Q9Zcuyk",
	"QTb2TbaG8/XePIp85JkpK6NBo+PpI3fZLZSCHj9bEAhfzEzpv+HOg8Pwb2VNBRYVkE0lnJsbK8MzLirg",
	"KXdolZ7xOuHegdWihJ6XdcIt3HllQfL0+9oyWXu8SZaLzPQHZBg8/qHFtIDruflTZGjszrQcZBawN6l7",
	"sCpXmUBl9CQz8oDsGm99a/uSpI51c8qonXIiKK/c2DI8cSkQzlBR7Z1klWzZeq9kn9myZ5N7sE4ZvVGS",
	"0ggzsMEK52aSU98mQH3cRG1qTAFCE2yVfHaiT0PdadE3CvE0tV7Anpudoa6aDr2EyFvRniBoMFU6N+RE",
	"YRHeUZGOfRVazKAEjez3q3HkUoSKf/ww/DAMqZoKtKgUT/lv9FcAFm8pvUFBbsLjLBI7JE9EHEue8kvA",
	"GIgWWVECgnU8/f7IJbjMqgpjsK/iQZW+ZNqXU7DM5Cx6ZmiYBfRW81ADT/mdB7vgCY8Y8EKVKsxAFIiQ",
	"gYRc+AJ5ej5MOnSrk+3If+W5A2S5sawSM6Up+R3BDNn2Rzso2D8gbHbLEGxJEXNVIAQ4m3J3xHW0rBV3",
	"mww3gQ2uMtpF4pwPhzTdRiNoQkZUVdGIxOCHi9O49tem3RpWhVDSw68Wcp7yXwZrYR40qjyI0rKeKGGt",
	"WNBvg6KIbdpsw3X4u4M17+tgH5fbzr4ohxtOgoHzZSnsYvlSFMXqbcIr43qYemXcmqo2juMnIxfPauJT",
	"PerZtur2DKP1UHdg/PhqGTQo9XQwvGDNRsCczzJwLvdFQRhe/C8mleCcmJFawoMoKxKfb40yMlFYEHLB",
	"4EE5dF0FPwT/sb4XhZLMrrq6SYDYdiaYhnkkARk0wjV4VLKO/CwAoUuLEf0fiTEedVWMpjXI4XpYaS9s",
	"o7o5uHt2zZ5BvugOUEQsJt2H2MVrIxYDahOE0mv5MqS2nbSRiq1eYpTs2VB+FhjDtxpHCShU4Y4Xz0vA",
	"CCabLth4RNLr+5TXvz2kry/vPZ+OB8n7m/Gp+Xw+EbGI7d4h6IPzXAykcuEwQSnv2fHHo/NcjBr7U1WV",
	"67k5i8csJjzegsYmAGs6dTL7SCyH4a6C+/kSz54H0yUe+Y9YsHbcWbwX0dpN1+aS4ETYGmF4DlmnIvvX",
	"V2eZkc0VwV66fqIVn2nBEchbG5hY7mRV7upA2rnsaR8+D4EmNoaRbzYDHRp4Msy6bOqhk890o9IupTZv",
	"mfZ9oy2vrI7+W2377u29KN8yrxP+YmMrwtV1Xf8XAAD//1skG0rsFwAA",
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
