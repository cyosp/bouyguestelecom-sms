package bouyguestelecom

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type loginner interface {
	Login(login, pass string) error
}

type httpLoginner struct {
	client httpClient
}

func (l *httpLoginner) Login(lastname, login, pass string) error {
	tokens, err := l.getTokens()
	if err != nil {
		return err
	}

	return l.postLogin(lastname, login, pass, tokens)
}

type tokens struct {
	execution string
}

func (l *httpLoginner) getTokens() (*tokens, error) {
	body, err := l.client.Get("https://www.mon-compte.bouyguestelecom.fr/cas/login?&service=https://oauth2.bouyguestelecom.fr/callback/picasso/protocol/cas")
	if err != nil {
		return nil, err
	}

	execution, err := l.extractExecution(body)
	if err != nil {
		return nil, err
	}

	return &tokens{execution}, nil
}

func (l *httpLoginner) extractExecution(body string) (string, error) {
	regex := regexp.MustCompile("name=\"execution\" value=\"(.+?)\"")
	matches := regex.FindStringSubmatch(body)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("execution token not found")
}

func (l *httpLoginner) postLogin(lastname, login, pass string, tokens *tokens) error {
	loginURL := "https://www.mon-compte.bouyguestelecom.fr/cas/login?service=https%3A%2F%2Fwww.secure.bbox.bouyguestelecom.fr%2Fservices%2FSMSIHD%2FsendSMS.phtml"

	data := make(url.Values)
	data.Add("lastname", lastname)
	data.Add("username", login)
	data.Add("password", pass)
	data.Add("rememberMe", "true")
	data.Add("_rememberMe", "on")
	data.Add("geolocation", "")
	data.Add("execution", tokens.execution)
	data.Add("execution", "e1s1")
	data.Add("_eventId", "submit")

	body, err := l.client.PostForm(loginURL, data)
	if err != nil {
		return err
	}

	if strings.Contains(body, "Votre identifiant ou votre mot de passe est incorrect") {
		return errors.New("invalid credentials")
	}

	return nil
}
