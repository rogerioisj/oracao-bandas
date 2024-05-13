package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracao-bandas.com/src/database/repositories"
)

type SessionCheckInterface interface {
	CheckSessionMiddleware() gin.HandlerFunc
}

type SessionCheckMiddleware struct {
	sessionRepository repositories.SessionRepositoryInterface
}

func NewSessionCheck(sessionRepository repositories.SessionRepositoryInterface) *SessionCheckMiddleware {
	return &SessionCheckMiddleware{
		sessionRepository: sessionRepository,
	}
}

func (service *SessionCheckMiddleware) CheckSessionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessionToken, err := context.Cookie("session_token")
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		session, errSearch := service.sessionRepository.Search(sessionToken)

		if errSearch != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		if session.Login == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		context.Next()
	}
}
