package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"

	adminCtx   = "adminId"
	studentCtx = "studentId"
)

func (h *Handler) studentIdentity(ctx *gin.Context) {
	id, err := h.parseAuthHeader(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ctx.Set(studentCtx, id)
	ctx.Next()
}

func (h *Handler) parseAuthHeader(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}
	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}
