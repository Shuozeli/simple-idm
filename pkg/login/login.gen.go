// Package login provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package login

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

	openapi_types "github.com/discord-gophers/goapi-gen/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// EmailVerifyRequest defines model for EmailVerifyRequest.
type EmailVerifyRequest struct {
	Email string `json:"email"`
}

// FindUsernameRequest defines model for FindUsernameRequest.
type FindUsernameRequest struct {
	// Email address to find username for
	Email openapi_types.Email `json:"email"`
}

// Login defines model for Login.
type Login struct {
	// Token for 2FA verification if required
	LoginToken *string `json:"loginToken,omitempty"`
	Message    string  `json:"message"`

	// Whether 2FA verification is required
	Requires2fA *bool  `json:"requires2FA,omitempty"`
	Status      string `json:"status"`
	User        User   `json:"user"`

	// List of users associated with the login. Usually contains one user, but may contain multiple if same username is shared.
	Users []User `json:"users,omitempty"`
}

// PasswordReset defines model for PasswordReset.
type PasswordReset struct {
	Code     string `json:"code"`
	Password string `json:"password"`
}

// PasswordResetInit defines model for PasswordResetInit.
type PasswordResetInit struct {
	// Username of the account to reset password for
	Username string `json:"username"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Tokens defines model for Tokens.
type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// TwoFactorVerify defines model for TwoFactorVerify.
type TwoFactorVerify struct {
	// TOTP code
	Code string `json:"code"`

	// Token from initial login response
	LoginToken string `json:"loginToken"`
}

// User defines model for User.
type User struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`

	// Whether 2FA is enabled for this user
	TwoFactorEnabled bool `json:"twoFactorEnabled"`
}

// Post2faVerifyJSONBody defines parameters for Post2faVerify.
type Post2faVerifyJSONBody TwoFactorVerify

// PostEmailVerifyJSONBody defines parameters for PostEmailVerify.
type PostEmailVerifyJSONBody EmailVerifyRequest

// PostLoginJSONBody defines parameters for PostLogin.
type PostLoginJSONBody struct {
	// 2FA verification code if enabled
	Code     string `json:"code,omitempty"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// PostMobileLoginJSONBody defines parameters for PostMobileLogin.
type PostMobileLoginJSONBody struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// PostPasswordResetJSONBody defines parameters for PostPasswordReset.
type PostPasswordResetJSONBody PasswordReset

// PostPasswordResetInitJSONBody defines parameters for PostPasswordResetInit.
type PostPasswordResetInitJSONBody PasswordResetInit

// PostRegisterJSONBody defines parameters for PostRegister.
type PostRegisterJSONBody RegisterRequest

// PostUserSwitchJSONBody defines parameters for PostUserSwitch.
type PostUserSwitchJSONBody struct {
	// ID of the user to switch to
	UserID string `json:"user_id"`
}

// PostUsernameFindJSONBody defines parameters for PostUsernameFind.
type PostUsernameFindJSONBody FindUsernameRequest

// Post2faVerifyJSONRequestBody defines body for Post2faVerify for application/json ContentType.
type Post2faVerifyJSONRequestBody Post2faVerifyJSONBody

// Bind implements render.Binder.
func (Post2faVerifyJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostEmailVerifyJSONRequestBody defines body for PostEmailVerify for application/json ContentType.
type PostEmailVerifyJSONRequestBody PostEmailVerifyJSONBody

// Bind implements render.Binder.
func (PostEmailVerifyJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody PostLoginJSONBody

// Bind implements render.Binder.
func (PostLoginJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostMobileLoginJSONRequestBody defines body for PostMobileLogin for application/json ContentType.
type PostMobileLoginJSONRequestBody PostMobileLoginJSONBody

// Bind implements render.Binder.
func (PostMobileLoginJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostPasswordResetJSONRequestBody defines body for PostPasswordReset for application/json ContentType.
type PostPasswordResetJSONRequestBody PostPasswordResetJSONBody

// Bind implements render.Binder.
func (PostPasswordResetJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostPasswordResetInitJSONRequestBody defines body for PostPasswordResetInit for application/json ContentType.
type PostPasswordResetInitJSONRequestBody PostPasswordResetInitJSONBody

// Bind implements render.Binder.
func (PostPasswordResetInitJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody PostRegisterJSONBody

// Bind implements render.Binder.
func (PostRegisterJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostUserSwitchJSONRequestBody defines body for PostUserSwitch for application/json ContentType.
type PostUserSwitchJSONRequestBody PostUserSwitchJSONBody

// Bind implements render.Binder.
func (PostUserSwitchJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PostUsernameFindJSONRequestBody defines body for PostUsernameFind for application/json ContentType.
type PostUsernameFindJSONRequestBody PostUsernameFindJSONBody

// Bind implements render.Binder.
func (PostUsernameFindJSONRequestBody) Bind(*http.Request) error {
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

// Post2faVerifyJSON200Response is a constructor method for a Post2faVerify response.
// A *Response is returned with the configured status code and content type from the spec.
func Post2faVerifyJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostEmailVerifyJSON200Response is a constructor method for a PostEmailVerify response.
// A *Response is returned with the configured status code and content type from the spec.
func PostEmailVerifyJSON200Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostLoginJSON200Response is a constructor method for a PostLogin response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostLoginJSON202Response is a constructor method for a PostLogin response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginJSON202Response(body struct {
	N2faMethods []string `json:"2fa_methods,omitempty"`
	Message     *string  `json:"message,omitempty"`
	Status      *string  `json:"status,omitempty"`

	// Temporary token to use for 2FA verification
	TempToken *string `json:"temp_token,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        202,
		contentType: "application/json",
	}
}

// PostLoginJSON400Response is a constructor method for a PostLogin response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLoginJSON400Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// PostLogoutJSON200Response is a constructor method for a PostLogout response.
// A *Response is returned with the configured status code and content type from the spec.
func PostLogoutJSON200Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostMobileLoginJSON200Response is a constructor method for a PostMobileLogin response.
// A *Response is returned with the configured status code and content type from the spec.
func PostMobileLoginJSON200Response(body struct {
	// JWT access token
	AccessToken string `json:"accessToken"`

	// JWT refresh token
	RefreshToken string `json:"refreshToken"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostPasswordResetJSON200Response is a constructor method for a PostPasswordReset response.
// A *Response is returned with the configured status code and content type from the spec.
func PostPasswordResetJSON200Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostPasswordResetInitJSON200Response is a constructor method for a PostPasswordResetInit response.
// A *Response is returned with the configured status code and content type from the spec.
func PostPasswordResetInitJSON200Response(body struct {
	Code *string `json:"code,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostRegisterJSON201Response is a constructor method for a PostRegister response.
// A *Response is returned with the configured status code and content type from the spec.
func PostRegisterJSON201Response(body struct {
	Email *openapi_types.Email `json:"email,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        201,
		contentType: "application/json",
	}
}

// PostTokenRefreshJSON200Response is a constructor method for a PostTokenRefresh response.
// A *Response is returned with the configured status code and content type from the spec.
func PostTokenRefreshJSON200Response(body Tokens) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostUserSwitchJSON200Response is a constructor method for a PostUserSwitch response.
// A *Response is returned with the configured status code and content type from the spec.
func PostUserSwitchJSON200Response(body Login) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostUserSwitchJSON400Response is a constructor method for a PostUserSwitch response.
// A *Response is returned with the configured status code and content type from the spec.
func PostUserSwitchJSON400Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// PostUserSwitchJSON403Response is a constructor method for a PostUserSwitch response.
// A *Response is returned with the configured status code and content type from the spec.
func PostUserSwitchJSON403Response(body struct {
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        403,
		contentType: "application/json",
	}
}

// PostUsernameFindJSON200Response is a constructor method for a PostUsernameFind response.
// A *Response is returned with the configured status code and content type from the spec.
func PostUsernameFindJSON200Response(body struct {
	// A message indicating the request was processed
	Message *string `json:"message,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Verify 2FA code during login
	// (POST /2fa/verify)
	Post2faVerify(w http.ResponseWriter, r *http.Request) *Response
	// Verify email address
	// (POST /email/verify)
	PostEmailVerify(w http.ResponseWriter, r *http.Request) *Response
	// Login a user
	// (POST /login)
	PostLogin(w http.ResponseWriter, r *http.Request) *Response
	// Logout user
	// (POST /logout)
	PostLogout(w http.ResponseWriter, r *http.Request) *Response
	// Mobile login endpoint
	// (POST /mobile/login)
	PostMobileLogin(w http.ResponseWriter, r *http.Request) *Response
	// Reset password
	// (POST /password/reset)
	PostPasswordReset(w http.ResponseWriter, r *http.Request) *Response
	// Initiate password reset using username
	// (POST /password/reset/init)
	PostPasswordResetInit(w http.ResponseWriter, r *http.Request) *Response
	// Register a new user
	// (POST /register)
	PostRegister(w http.ResponseWriter, r *http.Request) *Response
	// Refresh JWT tokens
	// (POST /token/refresh)
	PostTokenRefresh(w http.ResponseWriter, r *http.Request) *Response
	// Switch to a different user when multiple users are available for the same login
	// (POST /user/switch)
	PostUserSwitch(w http.ResponseWriter, r *http.Request) *Response
	// Send username to user's email address
	// (POST /username/find)
	PostUsernameFind(w http.ResponseWriter, r *http.Request) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// Post2faVerify operation middleware
func (siw *ServerInterfaceWrapper) Post2faVerify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Post2faVerify(w, r)
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

// PostEmailVerify operation middleware
func (siw *ServerInterfaceWrapper) PostEmailVerify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostEmailVerify(w, r)
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

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostLogin(w, r)
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

// PostLogout operation middleware
func (siw *ServerInterfaceWrapper) PostLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostLogout(w, r)
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

// PostMobileLogin operation middleware
func (siw *ServerInterfaceWrapper) PostMobileLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostMobileLogin(w, r)
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

// PostPasswordReset operation middleware
func (siw *ServerInterfaceWrapper) PostPasswordReset(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostPasswordReset(w, r)
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

// PostPasswordResetInit operation middleware
func (siw *ServerInterfaceWrapper) PostPasswordResetInit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostPasswordResetInit(w, r)
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

// PostRegister operation middleware
func (siw *ServerInterfaceWrapper) PostRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostRegister(w, r)
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

// PostTokenRefresh operation middleware
func (siw *ServerInterfaceWrapper) PostTokenRefresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostTokenRefresh(w, r)
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

// PostUserSwitch operation middleware
func (siw *ServerInterfaceWrapper) PostUserSwitch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostUserSwitch(w, r)
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

// PostUsernameFind operation middleware
func (siw *ServerInterfaceWrapper) PostUsernameFind(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostUsernameFind(w, r)
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
		r.Post("/2fa/verify", wrapper.Post2faVerify)
		r.Post("/email/verify", wrapper.PostEmailVerify)
		r.Post("/login", wrapper.PostLogin)
		r.Post("/logout", wrapper.PostLogout)
		r.Post("/mobile/login", wrapper.PostMobileLogin)
		r.Post("/password/reset", wrapper.PostPasswordReset)
		r.Post("/password/reset/init", wrapper.PostPasswordResetInit)
		r.Post("/register", wrapper.PostRegister)
		r.Post("/token/refresh", wrapper.PostTokenRefresh)
		r.Post("/user/switch", wrapper.PostUserSwitch)
		r.Post("/username/find", wrapper.PostUsernameFind)
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

	"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",
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
