package microservice

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type msgError string

const (
	msgInvalidToken        msgError = "Token is invalid"
	msgInvalidCode         msgError = "Code is invalid"
	msgInvalidRefreshToken msgError = "Refresh token is invalid"
	msgUnauthorizedAccess  msgError = "Unauthorized access"
	msgJWTExpired          msgError = "JWT Token is expired"
	msgTokenEmpty          msgError = "token up for parsing was not passed through the context"
)

var (
	listFormatingMsgErrors []msgError = []msgError{
		msgInvalidToken,
		msgUnauthorizedAccess,
		msgInvalidRefreshToken,
		msgJWTExpired,
		msgTokenEmpty,
	}

	//ErrorInvalidToken is response for token is invalid
	ErrorInvalidToken = status.Error(codes.PermissionDenied, "Token is invalid")
	// ErrorInvalidCode is response for Code is invalid
	ErrorInvalidCode = status.Error(codes.PermissionDenied, "Code is invalid")
	// ErrorInvalidRefreshToken is response for Refresh token is invalid
	ErrorInvalidRefreshToken = status.Error(codes.PermissionDenied, "Refresh token is invalid")
	//ErrUnauthorized is error for unauthorized access
	ErrUnauthorized = status.Error(codes.PermissionDenied, "Unauthorized access")
	//ErrJWTExpired is error for token expired
	ErrJWTExpired = status.Error(codes.PermissionDenied, "JWT Token is expired")
	//ErrTokenEmpty is error for token is empty
	ErrTokenEmpty = status.Error(codes.PermissionDenied, "token up for parsing was not passed through the context")
)

// FilterType ..
type FilterType int32

// list const of FilterType
const (
	NotFilter      FilterType = 0
	FilterDriverID FilterType = 1
	FilterAreaID   FilterType = 2
	FilterPoolID   FilterType = 3
)

type contextKey string

var (
	//CtxACL is context key for acl
	CtxACL  = contextKey("acl")
	CtxACLS = contextKey("acls")
	//CtxUserID is context key for user id
	CtxUserID = contextKey("uid")
	//CtxUserUUID is context key for user uuid
	CtxUserUUID = contextKey("uuid")
	//CtxDomain is context key for domain
	CtxDomain = contextKey("domain")
	//CtxPhone is context key for phone number
	CtxPhone = contextKey("phone_number")
	//CtxEmail is context key for email
	CtxEmail = contextKey("email")
	//CtxDomainName is context key for domain name
	CtxDomainName = contextKey("domain_name")
	//CtxTopic is context key for topic name
	CtxTopic = contextKey("topic")
	//CtxAudit is context key for audit parent id
	CtxAudit = contextKey("audit")
	//CtxDomainID is context key for domain id
	CtxDomainID = contextKey("domain_id")
	//CtxDomainType is context key for domain type
	CtxDomainType = contextKey("domain_type")
	//CtxExp is context key for expired time
	CtxExp = contextKey("exp")
	//CtxGroupName is context key for group name of the user
	CtxGroupName = contextKey("group_name")
	//CtxRequestID is context key for requestID
	CtxRequestID = contextKey("request_id")
	// CtxRequestUUID is context key for user UUID which request
	CtxRequestUUID = contextKey("request_uuid")
	// CtxRequestName is context key for user_name which request
	CtxRequestName = contextKey("request_name")
	// CtxSubName is context key for get full user name
	CtxSubName = contextKey("sub")
	// CtxAuthTime is context key for get time when user auth
	CtxAuthTime = contextKey("auth_time")
	// CtxBBGBBDAccess is context key for list of access
	CtxBBGBBDAccess = contextKey("bbg.bbd.access")
	// CtxAccessAreas is context key for list of access areas
	CtxAccessAreas = contextKey("areas")
	// CtxAccessPools is context key for list of access pools
	CtxAccessPools = contextKey("pools")
	// CtxSecurityPolicy is context key security policy
	CtxSecurityPolicy = contextKey("security_policy")
)
