package main

import (
	"github.com/go-martini/martini"
	"github.com/handshakejs/handshakejserrors"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/scottmotte/cartelogic"
	"net/http"
	"os"
)

const (
	LOGIC_ERROR_CODE_UNKNOWN = "unknown"
)

var (
	DB_ENCRYPTION_SALT string
	REDIS_URL          string
)

func CrossDomain() martini.Handler {
	return func(res http.ResponseWriter) {
		res.Header().Add("Access-Control-Allow-Origin", "*")
	}
}

type Deck struct {
	Name   string `form:"name" json:"name"`
	Email  string `form:"email" json:"email"`
	ApiKey string `form:"api_key" json:"api_key"`
}

func main() {
	loadEnvs()

	cartelogic.Setup(REDIS_URL)

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(CrossDomain())

	m.Any("/api/v0/decks/create.json", binding.Bind(Deck{}), DecksCreate)

	m.Run()
}

func ErrorPayload(logic_error *handshakejserrors.LogicError) map[string]interface{} {
	error_object := map[string]interface{}{"code": logic_error.Code, "field": logic_error.Field, "message": logic_error.Message}
	errors := []interface{}{}
	errors = append(errors, error_object)
	payload := map[string]interface{}{"errors": errors}

	return payload
}

func DecksPayload(deck map[string]interface{}) map[string]interface{} {
	decks := []interface{}{}
	decks = append(decks, deck)
	payload := map[string]interface{}{"decks": decks}

	return payload
}

func DecksCreate(deck Deck, req *http.Request, r render.Render) {
	email := deck.Email
	name := deck.Name
	api_key := deck.ApiKey

	params := map[string]interface{}{"email": email, "name": name, "api_key": api_key}
	result, logic_error := cartelogic.DecksCreate(params)
	if logic_error != nil {
		payload := ErrorPayload(logic_error)
		statuscode := determineStatusCodeFromLogicError(logic_error)
		r.JSON(statuscode, payload)
	} else {
		payload := DecksPayload(result)
		r.JSON(200, payload)
	}
}

func determineStatusCodeFromLogicError(logic_error *handshakejserrors.LogicError) int {
	code := 400
	if logic_error.Code == LOGIC_ERROR_CODE_UNKNOWN {
		code = 500
	}

	return code
}

func loadEnvs() {
	godotenv.Load()

	DB_ENCRYPTION_SALT = os.Getenv("DB_ENCRYPTION_SALT")
	REDIS_URL = os.Getenv("REDIS_URL")
}
