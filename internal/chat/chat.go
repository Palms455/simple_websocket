package chat

import (
	"github.com/bmizerany/pat"
	"github.com/sirupsen/logrus"
	"net/http"
	"simple_websocket/internal/handlers"
)

//Базовая струтура Чата
type Chat struct {
	config *Config
	logger *logrus.Logger
	routes *pat.PatternServeMux
}

//Chat конструктор
func New(config *Config) *Chat {
	return &Chat{
		config: config,
		logger: logrus.New(),
		routes: pat.New(),
	}
}

//Start Конфигурирует сервер, логгер, роутинг
func (chat *Chat) Start() error {
	//Конфигурация логгера
	if err := chat.configLogger(); err != nil {
		return err
	}

	chat.logger.Info("starting ws server at port:", chat.config.BindAddr)

	chat.configRoutes()
	chat.logger.Info("configure routes")
	chat.logger.Info("start ws Listen")
	go handlers.ListenToWsChan()
	return http.ListenAndServe(chat.config.BindAddr, chat.routes)
}