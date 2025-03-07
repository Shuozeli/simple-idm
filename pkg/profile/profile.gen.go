// Package profile provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package profile

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	// Error code
	Code string `json:"code"`

	// Error message
	Message string `json:"message"`
}

// PasswordPolicyResponse defines model for PasswordPolicyResponse.
type PasswordPolicyResponse struct {
	// Whether common passwords are disallowed
	DisallowCommonPwds *bool `json:"disallow_common_pwds,omitempty"`

	// Number of days until password expires
	ExpirationDays *int `json:"expiration_days,omitempty"`

	// Number of previous passwords to check against
	HistoryCheckCount *int `json:"history_check_count,omitempty"`

	// Maximum number of repeated characters allowed
	MaxRepeatedChars *int `json:"max_repeated_chars,omitempty"`

	// Minimum length of the password
	MinLength *int `json:"min_length,omitempty"`

	// Whether the password requires a digit
	RequireDigit *bool `json:"require_digit,omitempty"`

	// Whether the password requires a lowercase letter
	RequireLowercase *bool `json:"require_lowercase,omitempty"`

	// Whether the password requires a special character
	RequireSpecialChar *bool `json:"require_special_char,omitempty"`

	// Whether the password requires an uppercase letter
	RequireUppercase *bool `json:"require_uppercase,omitempty"`
}

// TwoFactorDisable defines model for TwoFactorDisable.
type TwoFactorDisable struct {
	// Current TOTP code
	Code string `json:"code"`

	// Current account password
	CurrentPassword string `json:"currentPassword"`
}

// TwoFactorEnable defines model for TwoFactorEnable.
type TwoFactorEnable struct {
	// Current TOTP code
	Code string `json:"code"`

	// TOTP secret from setup
	Secret string `json:"secret"`
}

// TwoFactorSetup defines model for TwoFactorSetup.
type TwoFactorSetup struct {
	// otpauth:// URL for manual setup
	OtpauthURL *string `json:"otpauthUrl,omitempty"`

	// Data URI of QR code image
	QrCode *string `json:"qrCode,omitempty"`

	// TOTP secret key
	Secret *string `json:"secret,omitempty"`
}

// Post2faDisableJSONBody defines parameters for Post2faDisable.
type Post2faDisableJSONBody TwoFactorDisable

// Post2faEnableJSONBody defines parameters for Post2faEnable.
type Post2faEnableJSONBody TwoFactorEnable

// ChangePasswordJSONBody defines parameters for ChangePassword.
type ChangePasswordJSONBody struct {
	// User's current password
	CurrentPassword string `json:"current_password"`

	// User's new password
	NewPassword string `json:"new_password"`
}

// ChangeUsernameJSONBody defines parameters for ChangeUsername.
type ChangeUsernameJSONBody struct {
	// User's current password for verification
	CurrentPassword string `json:"currentPassword"`

	// New username to set
	NewUsername string `json:"newUsername"`
}

// Post2faDisableJSONRequestBody defines body for Post2faDisable for application/json ContentType.
type Post2faDisableJSONRequestBody Post2faDisableJSONBody

// Bind implements render.Binder.
func (Post2faDisableJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Post2faEnableJSONRequestBody defines body for Post2faEnable for application/json ContentType.
type Post2faEnableJSONRequestBody Post2faEnableJSONBody

// Bind implements render.Binder.
func (Post2faEnableJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// ChangePasswordJSONRequestBody defines body for ChangePassword for application/json ContentType.
type ChangePasswordJSONRequestBody ChangePasswordJSONBody

// Bind implements render.Binder.
func (ChangePasswordJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// ChangeUsernameJSONRequestBody defines body for ChangeUsername for application/json ContentType.
type ChangeUsernameJSONRequestBody ChangeUsernameJSONBody

// Bind implements render.Binder.
func (ChangeUsernameJSONRequestBody) Bind(*http.Request) error {
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

// Post2faDisableJSON200Response is a constructor method for a Post2faDisable response.
// A *Response is returned with the configured status code and content type from the spec.
func Post2faDisableJSON200Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// Post2faEnableJSON200Response is a constructor method for a Post2faEnable response.
// A *Response is returned with the configured status code and content type from the spec.
func Post2faEnableJSON200Response(body struct {
	// One-time use backup codes
	BackupCodes []string `json:"backupCodes,omitempty"`
	Message     *string  `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// Post2faSetupJSON200Response is a constructor method for a Post2faSetup response.
// A *Response is returned with the configured status code and content type from the spec.
func Post2faSetupJSON200Response(body TwoFactorSetup) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// ChangePasswordJSON400Response is a constructor method for a ChangePassword response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangePasswordJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// ChangePasswordJSON401Response is a constructor method for a ChangePassword response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangePasswordJSON401Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        401,
		contentType: "application/json",
	}
}

// ChangePasswordJSON403Response is a constructor method for a ChangePassword response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangePasswordJSON403Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        403,
		contentType: "application/json",
	}
}

// ChangePasswordJSON500Response is a constructor method for a ChangePassword response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangePasswordJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// GetPasswordPolicyJSON200Response is a constructor method for a GetPasswordPolicy response.
// A *Response is returned with the configured status code and content type from the spec.
func GetPasswordPolicyJSON200Response(body PasswordPolicyResponse) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// ChangeUsernameJSON400Response is a constructor method for a ChangeUsername response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangeUsernameJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// ChangeUsernameJSON401Response is a constructor method for a ChangeUsername response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangeUsernameJSON401Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        401,
		contentType: "application/json",
	}
}

// ChangeUsernameJSON403Response is a constructor method for a ChangeUsername response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangeUsernameJSON403Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        403,
		contentType: "application/json",
	}
}

// ChangeUsernameJSON409Response is a constructor method for a ChangeUsername response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangeUsernameJSON409Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        409,
		contentType: "application/json",
	}
}

// ChangeUsernameJSON500Response is a constructor method for a ChangeUsername response.
// A *Response is returned with the configured status code and content type from the spec.
func ChangeUsernameJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Disable 2FA for the user
	// (POST /2fa/disable)
	Post2faDisable(w http.ResponseWriter, r *http.Request) *Response
	// Enable 2FA for the user
	// (POST /2fa/enable)
	Post2faEnable(w http.ResponseWriter, r *http.Request) *Response
	// Generate 2FA secret and QR code
	// (POST /2fa/setup)
	Post2faSetup(w http.ResponseWriter, r *http.Request) *Response
	// Change user password
	// (PUT /password)
	ChangePassword(w http.ResponseWriter, r *http.Request) *Response
	// Get password policy
	// (GET /password/policy)
	GetPasswordPolicy(w http.ResponseWriter, r *http.Request) *Response
	// Change username
	// (PUT /username)
	ChangeUsername(w http.ResponseWriter, r *http.Request) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// Post2faDisable operation middleware
func (siw *ServerInterfaceWrapper) Post2faDisable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Post2faDisable(w, r)
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

// Post2faEnable operation middleware
func (siw *ServerInterfaceWrapper) Post2faEnable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Post2faEnable(w, r)
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

// Post2faSetup operation middleware
func (siw *ServerInterfaceWrapper) Post2faSetup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Post2faSetup(w, r)
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

// ChangePassword operation middleware
func (siw *ServerInterfaceWrapper) ChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.ChangePassword(w, r)
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

// GetPasswordPolicy operation middleware
func (siw *ServerInterfaceWrapper) GetPasswordPolicy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetPasswordPolicy(w, r)
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

// ChangeUsername operation middleware
func (siw *ServerInterfaceWrapper) ChangeUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.ChangeUsername(w, r)
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
		r.Post("/2fa/disable", wrapper.Post2faDisable)
		r.Post("/2fa/enable", wrapper.Post2faEnable)
		r.Post("/2fa/setup", wrapper.Post2faSetup)
		r.Put("/password", wrapper.ChangePassword)
		r.Get("/password/policy", wrapper.GetPasswordPolicy)
		r.Put("/username", wrapper.ChangeUsername)
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

	"H4sIAAAAAAAC/+xYb2/bthP+KgR/BX4Z5sRpsr6o36VpU2RoVy+JUWBBZpyls8xGItkjFcct9N0HkpL/",
	"yJJrt/GwAXtVR+Q9d/fw7uGVX3mkMq0kSmt47ys30QQz8D/fEClyPzQpjWQF+s+RitH9G6OJSGgrlOS9",
	"sJn5tQ63M428x40lIRNedHiGxkDSalYtr1kWHU74OReEMe/d8hK+2n43369GnzCyzlMfjJkqivsqFdHs",
	"Co1W0uB6FrEwkKZqOoxUlik51NPYrIf3cYJ2gi4vt4npEtwwIGQVBMaLwEdKpQjSRYKPWhA4oGEMswbw",
	"3/JshMTUmLl1lksr0rkL5s3RLKCFtJggOeiJMFbRbBhNMLofRiqXdhO8JnwQKjdL8VvFvDGDBIQ0ttFN",
	"Bo9DQo1gMR5GE6CGJN7Do8jyjMm5t8qCOQuILJJhazwtexFymKJM7KQBXUiPHtYdup3gPI1GuLJehrFI",
	"hG0/0WUcVtoYBiyYNZ1nBewyoQgM7g4+N2UpWou00Y/RGAlIPfG7uyqtF4ew0Veu9XfmJNnctj2poqFR",
	"b6bqAiKr6LUwMEpxW6E5z4lQWnbz4abfqjdR2FSJQTsKRL57GkpqswTVHdxtyvCNfPIEDUaEDeXtjcIi",
	"G5PKmEGb622zKlE3JnPtAddyUVZDbicDSteDKtd63S4bXL1jY6f4IHNI26Lr8M903kjOa7DABleXTgp+",
	"v/L0MJE13h3bkXSPs0Z2agQEtJyEnV27GzIk/QqBkM7yIF0j/9eFogws7/FfP944Sv1u1wx+deFrYq3m",
	"hQMWcqycvRXWVQnvkxqLFNl7kJBg5mrhrH/JO/wByYQMnh8dHx27FJVGCVrwHj/1nzpcg5344LonY+jG",
	"S82ljOfCnZq/li5j50wZezKGqglDWaCxr1Q8CzUqLYbrBbROReRNu5+Mi6MaF9yvZ4Rj3uP/6y7miW45",
	"THTXer1YLUBLOfoP4bb24Z8cH+/kf7UclwYOfIRMe2JPLs5YSUjMTB5FaMw4T9PtKqDo1MqoHW65XHjv",
	"drVQbu+Kuw43eZYBzVxJBwjm4FxrOKHNjbvMik44RJRbnWEpM3s+wtLL3k9wBNF9rp0KNEwdHyQeWpF5",
	"oljY6bXAjUvCYuZN1vSg/ABEMKtNpatFEgh/qhppRNupRALnGyrEzEV5U4Fcl2r7Q6e0VZEETy18+GCZ",
	"Uz2nlEJJlqB04f4QR29LDBZceHEHGVe3ROBKLw0EOm+4GgY6dhgVw/9fzMwMxhaJPSCJ8UzIxO0RxMpB",
	"YHmAWOX+fAIywf5i+Xu7szY7BL9D3TriDEL8DQGuNYbE6beBJE53GJPq4dWc3DX2zlZ6shpcRSzL/cnV",
	"K6jDf3nC+g7/I24o60v5AKkIczEayw7wKDnqMFF+rp/BTyGw5/sPbCDd5KVIfMGYHUhlWaqSBGMmZBnE",
	"6f6DuFA0EnGMkh1sZOTF33NUFkn62ZMekBiWG2tSM9ogNaGlvUAsOmJFX7raPz+4IBNsEOS3aFcfKvap",
	"yi1PIg3cXM97JxCyJK6Lo2Jlbj5hx4GEDHcQ1MqkRVC/IaSDyuETC2l/Vx31N7EPvvTUIqyDJYZqDzQ4",
	"XZBhlbsXuX90eVe+hrw49q8j1Z+nfrh31ct7/M9bOPxydvjH8eHL4eHdz8+2VeT+iiDPg3syPa4Q/9F6",
	"PGc9jCD/yXHtYnq5/yDOlRynIrLsYKEHKSHEM2bhHuW/9ULw3VQURfFXAAAA//9vamUaWRcAAA==",
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
