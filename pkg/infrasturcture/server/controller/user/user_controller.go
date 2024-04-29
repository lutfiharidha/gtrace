package controller

import (
	"net/http"

	"github.com/lutfiharidha/google-trace/pkg/shared/tracing"
	userUc "github.com/lutfiharidha/google-trace/pkg/usecase/user"
	"go.opentelemetry.io/otel/attribute"
)

type UserController struct {
	userUc userUc.UserInterface
}

func NewUserController(userUc userUc.UserInterface) *UserController {
	return &UserController{
		userUc: userUc,
	}
}
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query().Get("id")
	newCtx, sp := tracing.CreateSpan(r.Context(), r.URL.Path)
	defer sp.End()

	sp.SetAttributes(attribute.String("Request_Controller", req))

	c.userUc.GetUser(newCtx, req)
	w.Write([]byte(req))
}
