package configurationhabdler

import (
	"github.com/NovikovAndrew/computer-club-admin-service.git/internal/adapters/api"
	"github.com/NovikovAndrew/computer-club-admin-service.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	configurationsURL = "/configurations"
	configurationURL  = "/configuration"
)

type configurationHandler struct {
	logger               *logging.Logger
	configurationService ConfigurationService
}

func NewComponentHandler(logger *logging.Logger, configurationService ConfigurationService) api.Handler {
	return &configurationHandler{
		logger:               logger,
		configurationService: configurationService,
	}
}

func (ch *configurationHandler) Register(router *httprouter.Router) {

}
