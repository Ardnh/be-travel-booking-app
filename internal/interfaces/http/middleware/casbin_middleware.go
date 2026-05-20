package middleware

import (
	http "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/responses"
	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type CasbinMiddleware struct {
	enforcer *casbin.Enforcer
	log      *logrus.Logger
}

func NewCasbinMiddleware(enforcer *casbin.Enforcer, log *logrus.Logger) *CasbinMiddleware {
	return &CasbinMiddleware{enforcer: enforcer, log: log}
}

func (m *CasbinMiddleware) Authorize(resource, action string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(string)
		if !ok || userID == "" {
			m.log.Warn("casbin: user_id not found in context")
			return http.NewErrorResponse(c, fiber.ErrUnauthorized.Code, fiber.ErrUnauthorized.Message, nil)
		}

		allowed, err := m.enforcer.Enforce(userID, resource, action)
		if err != nil {
			m.log.WithFields(logrus.Fields{
				"user_id":  userID,
				"resource": resource,
				"action":   action,
				"error":    err,
			}).Error("casbin: failed to enforce policy")
			return http.NewErrorResponse(c, fiber.ErrInternalServerError.Code, "Failed to check permission", nil)
		}

		if !allowed {
			m.log.WithFields(logrus.Fields{
				"user_id":  userID,
				"resource": resource,
				"action":   action,
				"path":     c.Path(),
				"method":   c.Method(),
			}).Warn("casbin: access denied")
			return http.NewErrorResponse(c, fiber.ErrForbidden.Code, fiber.ErrForbidden.Message, nil)
		}

		m.log.WithFields(logrus.Fields{
			"user_id":  userID,
			"resource": resource,
			"action":   action,
		}).Debug("casbin: access granted")

		return c.Next()
	}
}
