package errs

import (
	"embed"
	"fmt"
	"net/http"
	"sync"

	"gopkg.in/yaml.v3"
)

var rawMessages embed.FS

// Структура для описания ошибки из YAML (messages.yaml)
type template struct {
	Code    int    `yaml:"code"`
	Status  int    `yaml:"status"`
	Message string `yaml:"message"`
}

// Интерфейс ошибки приложения
type AppError interface {
	error
	StatusCode() int
	Code() int
	Message() string
}

type appError struct {
	code   int
	status int
	msg    string
}

func (e appError) Error() string   { return e.msg }
func (e appError) StatusCode() int { return e.status }
func (e appError) Code() int       { return e.code }
func (e appError) Message() string { return e.msg }

var (
	once      sync.Once
	templates map[string]template
)

// Возвращает ошибку по ключу из messages.yaml и можно дополнить текст ошибки
func New(key string, formatArgs ...any) AppError {
	loadTemplates()

	tmpl, ok := templates[key]
	if !ok {

		tmpl = templates["INTERNAL"]
	}

	msg := tmpl.Message
	if len(formatArgs) > 0 {
		msg = fmt.Sprintf(msg, formatArgs...)
	}

	return appError{code: tmpl.Code, status: tmpl.Status, msg: msg}
}

// Инициализирует map с шаблонами ошибок
func loadTemplates() {
	once.Do(func() {
		data, _ := rawMessages.ReadFile("messages.yaml")
		templates = make(map[string]template)

		var raw map[string]template
		if err := yaml.Unmarshal(data, &raw); err != nil {

			templates["INTERNAL"] = template{Code: 1002, Status: http.StatusInternalServerError, Message: "Ошибка на сервере"}
			return
		}
		templates = raw
	})
}
