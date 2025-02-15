package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {
	c := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "tonic",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager
	ses := c.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type
	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("For loop: ", rv.Kind(), rv.Type(), rv)
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("Invalid type or kind; kind;", "type:", rv.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("worng kind returned testing cookie session. Expected", reflect.ValueOf(sm).Kind(), "AND GOT", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("worng kind returned testing cookie session. Expected", reflect.ValueOf(sm).Type(), "AND GOT", sessType)
	}

}
