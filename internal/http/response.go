package http

import (
	"math"
	"net/http"

	errorsPkg "github.com/Confialink/wallet-pkg-errors"
	list_params "github.com/Confialink/wallet-pkg-list_params"
	"github.com/gin-gonic/gin"
)

const (
	TargetCommon = "common"

	// all error codes
	Forbidden                     = "FORBIDDEN"
	NotFound                      = "NOT_FOUND"
	CanNotRetrieveCollection      = "CANNOT_RETRIEVE_COLLECTION"
	CanNotUpdateCollection        = "CANNOT_UPDATE_COLLECTION"
	BadTestSMTPParams             = "BAD_TEST_SMTP_PARAMS"
	CanNotRetrieveProviderDetails = "CANNOT_RETRIEVE_PROVIDER_DETAILS"
	UnprocessableEntity           = "UNPROCESSABLE_ENTITY"
	CodeInvalidQueryParameters    = "INVALID_QUERY_PARAMETERS"
)

var statusCodes = map[string]int{
	Forbidden:                     http.StatusForbidden,
	NotFound:                      http.StatusNotFound,
	CanNotRetrieveCollection:      http.StatusBadRequest,
	CanNotUpdateCollection:        http.StatusInternalServerError,
	BadTestSMTPParams:             http.StatusBadRequest,
	CanNotRetrieveProviderDetails: http.StatusInternalServerError,
}

// ResponseService is an empty structure
type ResponseService struct{}

// Response is the abstract response model
type Response struct {
	Status     int         `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Errors     []*Error    `json:"errors,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Error is the abstract error model
type Error struct {
	Title   string      `json:"title"`
	Details string      `json:"details"`
	Code    string      `json:"code"`
	Source  string      `json:"source,omitempty"`
	Target  string      `json:"target"`
	Meta    interface{} `json:"meta,omitempty"`
}

type Pagination struct {
	TotalRecord uint64 `json:"totalRecord"`
	TotalPage   uint64 `json:"totalPage"`
	Limit       uint64 `json:"limit"`
	CurrentPage uint64 `json:"currentPage"`
}

// NewResponseService creates new router
func NewResponseService() *ResponseService {
	return new(ResponseService)
}

// NewResponse creates a new Response instance.
func NewResponse() *Response {
	return new(Response)
}

// SetStatus sets the status for the response object.
func (r *Response) SetStatus(status int) *Response {
	r.Status = status
	return r
}

// SetStatusByCode sets the status for the response object by code.
func (r *Response) SetStatusByCode(code string) *Response {
	if status, ok := statusCodes[code]; ok {
		r.Status = status
	}
	return r
}

// SetData sets the data for the response object.
func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetPagination(pagination *Pagination) *Response {
	r.Pagination = pagination
	return r
}

// AddError adds error into the response object.
func (r *Response) AddError(err *Error) *Response {
	r.Errors = append(r.Errors, err)
	return r
}

// NewError creates a new Error instance.
func NewError() *Error {
	return new(Error)
}

// SetTitle sets the given title for the error object.
func (e *Error) SetTitle(title string) *Error {
	e.Title = title
	return e
}

// SetTitleByStatus sets the title for the error object by status code.
func (e *Error) SetTitleByStatus(status int) *Error {
	e.Title = http.StatusText(status)
	return e
}

// SetTitleByCode sets the title for the error object by code.
func (e *Error) SetTitleByCode(code string) *Error {
	if status, ok := statusCodes[code]; ok {
		e.Title = http.StatusText(status)
	}
	return e
}

// SetDetails sets the given details for the error object.
func (e *Error) SetDetails(details string) *Error {
	e.Details = details
	return e
}

// SetCode sets the code for the error object.
func (e *Error) SetCode(code string) *Error {
	e.Code = code
	return e
}

// SetSourсe sets the sourсe for the error object.
func (e *Error) SetSource(source string) *Error {
	e.Source = source
	return e
}

// SetTarget sets the sourсe for the error object.
func (e *Error) SetTarget(target string) *Error {
	e.Target = target
	return e
}

// SuccessResponse returns a success response
func (res *ResponseService) SuccessResponse(ctx *gin.Context, status int, data interface{}) {
	r := NewResponse().SetStatus(status).SetData(data)
	ctx.JSON(status, r)
}

// ErrorResponse returns a error response
func (res *ResponseService) ErrorResponse(ctx *gin.Context, code, details string) {
	e := NewError().SetCode(code).SetTitleByCode(code).SetTarget(TargetCommon).SetDetails(details)
	r := NewResponse().SetStatusByCode(code).AddError(e)
	ctx.AbortWithStatusJSON(r.Status, r)
}

// ValidatorErrorResponse returns a error response
func (res *ResponseService) ValidatorErrorResponse(ctx *gin.Context, code string, err error) {
	errorsPkg.AddShouldBindError(ctx, err)
}

// OkResponse returns a "200 StatusOK" response
func (res *ResponseService) OkResponse(ctx *gin.Context, data interface{}) {
	status := http.StatusOK
	r := NewResponse().SetStatus(status).SetData(data)
	ctx.JSON(status, r)
}

func (res *ResponseService) OkResponseWithPagination(
	ctx *gin.Context,
	data interface{},
	listParams *list_params.ListParams,
	total uint64,
) {
	status := http.StatusOK
	pagination := paginate(
		(uint64)(listParams.Pagination.PageNumber),
		(uint64)(listParams.Pagination.PageSize),
		total,
	)

	r := NewResponse().SetStatus(status).SetData(data).SetPagination(pagination)
	ctx.JSON(status, r)
}

// ForbiddenResponse returns a "403 StatusForbidden" response
func (res *ResponseService) ForbiddenResponse(ctx *gin.Context) {
	e := NewError().SetCode(Forbidden).SetTitleByCode(Forbidden).SetTarget(TargetCommon).SetDetails("User is not logged in.")
	r := NewResponse().SetStatusByCode(Forbidden).AddError(e)
	ctx.AbortWithStatusJSON(r.Status, r)
}

// NotFoundResponse returns a "404 StatusNotFound" response
func (res *ResponseService) NotFoundResponse(ctx *gin.Context) {
	e := NewError().SetCode(NotFound).SetTitleByCode(NotFound).SetTarget(TargetCommon).SetDetails("Not Found.")
	r := NewResponse().SetStatusByCode(NotFound).AddError(e)
	ctx.AbortWithStatusJSON(r.Status, r)
}

// paginate makes pagination
func paginate(number, size, total uint64) *Pagination {
	totalPage := uint64(math.Ceil(float64(total) / float64(size)))
	return &Pagination{
		Limit:       size,
		CurrentPage: number,
		TotalRecord: total,
		TotalPage:   totalPage,
	}
}
