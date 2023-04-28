package componenthandler

import (
	"github.com/NovikovAndrew/computer-club-admin-service.git/internal/adapters/api"
	"github.com/NovikovAndrew/computer-club-admin-service.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	componentsURL = "/components"
	componentURL  = "/component"
)

type componentHandler struct {
	logger           *logging.Logger
	componentService ComponentService
}

func NewComponentHandler(logger *logging.Logger, componentService ComponentService) api.Handler {
	return &componentHandler{
		logger:           logger,
		componentService: componentService,
	}
}

func (ch *componentHandler) Register(router *httprouter.Router) {

}
