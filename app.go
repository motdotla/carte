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

type Account struct {
	Email  string `form:"email" json:"email"`
	ApiKey string `form:"api_key" json:"api_key"`
}

type Card struct {
	ApiKey string `form:"api_key" json:"api_key"`
	Front  string `form:"front" json:"front"`
	Back   string `form:"back" json:"back"`
}

func main() {
	loadEnvs()

	cartelogic.Setup(REDIS_URL)

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(CrossDomain())

	m.Any("/api/v0/accounts/create.json", binding.Bind(Account{}), AccountsCreate)
	m.Any("/api/v0/cards/create.json", binding.Bind(Card{}), CardsCreate)

	m.Run()
}

func ErrorPayload(logic_error *handshakejserrors.LogicError) map[string]interface{} {
	error_object := map[string]interface{}{"code": logic_error.Code, "field": logic_error.Field, "message": logic_error.Message}
	errors := []interface{}{}
	errors = append(errors, error_object)
	payload := map[string]interface{}{"errors": errors}

	return payload
}

func AccountsPayload(account map[string]interface{}) map[string]interface{} {
	accounts := []interface{}{}
	accounts = append(accounts, account)
	payload := map[string]interface{}{"accounts": accounts}

	return payload
}

func AccountsCreate(account Account, req *http.Request, r render.Render) {
	email := account.Email
	api_key := account.ApiKey

	params := map[string]interface{}{"email": email, "api_key": api_key}
	result, logic_error := cartelogic.AccountsCreate(params)
	if logic_error != nil {
		payload := ErrorPayload(logic_error)
		statuscode := determineStatusCodeFromLogicError(logic_error)
		r.JSON(statuscode, payload)
	} else {
		payload := AccountsPayload(result)
		r.JSON(200, payload)
	}
}

func CardsPayload(card map[string]interface{}) map[string]interface{} {
	front := card["front"].(string)
	back := card["back"].(string)
	id := card["id"].(string)

	cards := []interface{}{}
	output_card := map[string]interface{}{"front": front, "back": back, "id": id}
	cards = append(cards, output_card)

	payload := map[string]interface{}{"cards": cards}

	return payload
}

func CardsCreate(card Card, req *http.Request, r render.Render) {
	front := card.Front
	back := card.Back
	api_key := card.ApiKey

	params := map[string]interface{}{"front": front, "back": back, "api_key": api_key}
	result, logic_error := cartelogic.CardsCreate(params)
	if logic_error != nil {
		payload := ErrorPayload(logic_error)
		statuscode := determineStatusCodeFromLogicError(logic_error)
		r.JSON(statuscode, payload)
	} else {
		payload := CardsPayload(result)
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

	REDIS_URL = os.Getenv("REDIS_URL")
}
