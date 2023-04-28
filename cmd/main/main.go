package main

import (
	"github.com/NovikovAndrew/computer-club-admin-service.git/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Start application")

}
