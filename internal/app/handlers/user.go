package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/osmodiar16/evalua-tu-profe/internal/app/models"
	"github.com/osmodiar16/evalua-tu-profe/internal/web/view/user"
)

type UserHandler struct{}

func (u UserHandler) HandleUserShow(c echo.Context) error {
	usuario := models.User{
		UserName: "osmodiar",
	}
	return render(c, user.Show(usuario))
}
