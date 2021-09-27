package chat

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"simple_websocket/internal/handlers"
)

//Конфигурация логгирования
func (chat *Chat) configLogger() error {
	log_level, err := logrus.ParseLevel(chat.config.LoggerLevel)
	if err != nil {
		chat.logger.SetLevel(log_level)
		return nil
	}
	return err
}

// configRoutes конфигурируем роутинг
func (chat *Chat) configRoutes() {
	chat.routes.Get("/", http.HandlerFunc(handlers.Home))
	chat.routes.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
}