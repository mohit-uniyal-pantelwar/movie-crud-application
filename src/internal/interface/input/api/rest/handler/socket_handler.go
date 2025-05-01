package handler

import (
	"movie-crud-application/src/pkg"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type SocketHandler struct {
	sugar *zap.SugaredLogger
}

func NewSocketHandler(sugar *zap.SugaredLogger) SocketHandler {
	return SocketHandler{
		sugar: sugar,
	}
}

func (sh *SocketHandler) UpgradeConnctionHandler(w http.ResponseWriter, r *http.Request) {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          "Something went wrong",
		}
		response.Set()
		return
	}

	defer connection.Close()

	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			sh.sugar.Errorf("Error reading message %s", err)
			break
		}

		if err := connection.WriteMessage(websocket.TextMessage, message); err != nil {
			sh.sugar.Errorf("Error reading message %s", err)
			break
		}
	}
}
