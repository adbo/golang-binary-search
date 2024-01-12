package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"main/data"
	"net/http"
	"strconv"
	"strings"
)

type NumberIndexHandler struct {
	store  *data.NumberStore
	logger *logrus.Logger
}

func NewNumberIndexHandler(store *data.NumberStore, logger *logrus.Logger) *NumberIndexHandler {
	return &NumberIndexHandler{
		store:  store,
		logger: logger,
	}
}

func (h *NumberIndexHandler) HandleNumberIndexRequest(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
  if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		h.logger.Warn("Received an invalid request")
		return
	}
	valueStr := pathParts[2]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		http.Error(w, "Value must be an integer", http.StatusBadRequest)
		h.logger.Warnf("Invalid value received: %v", valueStr)
		return
	}

	index, found := h.store.FindIndex(value)
	if !found {
		index, found = h.store.FindClosestIndex(value, 0.10)
		if !found {
			msg := fmt.Sprintf("%d", value)
			http.Error(w, msg, http.StatusNotFound)
			h.logger.Info(msg)
			return
		}
	}

	response := fmt.Sprintf("%d", index)
	fmt.Fprintln(w, response)
	h.logger.Infof("Handled request successfully: %s", response)
}
