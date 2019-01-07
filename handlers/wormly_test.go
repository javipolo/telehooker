package handlers

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var path = "/wormly"
var jsonValues0 = `{
"isrecovery":1,
"hostid":123,
"name":"juanito.com",
"downtime":10,
"alertlevel":3,
"alertlevel_name":"high",
"failedsensors":[
{"sensorid":33,"message":"http a la mierda","type":"HTTP"},
{"sensorid":69,"message":"otra cosa rota","type":"VAYAKK"}
]}`

var resultMessage0 = `RECOVERY on host juanito.com
Sensor HTTP - http a la mierda
Sensor VAYAKK - otra cosa rota
Duration 10 secs`

func buildRequest(j string) http.Request {
	d := url.Values{}
	d.Set(postParameter, j)
	r, _ := http.NewRequest("POST", path, strings.NewReader(d.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return *r
}

func TestParseData(t *testing.T) {
	r := buildRequest(jsonValues0)
	_, alert := parseData(&r)

	if alert.Name != "juanito.com" {
		t.Fatalf("alert.Name does not match")
	}
	if alert.IsRecovery != 1 {
		t.Fatalf("alert.IsRecovery does not match")
	}
	if alert.Hostid != 123 {
		t.Fatalf("alert.Hostid does not match")
	}
	if alert.AlertLevel != 3 {
		t.Fatalf("alert.AlertLevel does not match")
	}
	if alert.AlertLevelName != "high" {
		t.Fatalf("alert.AlertLevelName does not match: %v", alert)
	}
}

func TestCreateMessage(t *testing.T) {
	r := buildRequest(jsonValues0)
	_, alert := parseData(&r)
	msg := createMessage(alert)

	if msg != resultMessage0 {
		t.Fatalf("Message does not match:\n%s", msg)
	}
}
