package constants

type Status string

const (
	Inactive Status = "INACTIVE"
	Active   Status = "ACTIVE"
	Blocked  Status = "BLOCKED"
)

type UserType string

const (
	Employee UserType = "EMPLOYEE"
	Company  UserType = "COMPANY"
	Partner  UserType = "PARTNER"
	Agent    UserType = "AGENT"
)

type ResponseType string

const (
	ResponseOK           ResponseType = "OK"
	ResponseBad          ResponseType = "BAD"
	ResponseError        ResponseType = "ERROR"
	ResponseConflict     ResponseType = "CONFLICT"
	ResponseForbidden    ResponseType = "FORBIDDEN"
	ResponseUnauthorized ResponseType = "UNAUTHORIZED"
	ResponseNotFound     ResponseType = "NOT_FOUND"
)
