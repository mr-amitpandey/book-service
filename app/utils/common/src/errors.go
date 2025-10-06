package common

import "errors"

var (
    ErrSessionNotFound = errors.New("refresh session not found")

    ErrTokenRevoked = errors.New("refresh token revoked")

    ErrTokenExpired = errors.New("refresh token expired")
)
