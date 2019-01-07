package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/javipolo/telehooker/notify"
)

var service = "wormly"
var postParameter = "wormlyalert"

type failedSensor struct {
	Type     string `json:"type"`
	SensorID int    `json:"sensorid"`
	Message  string `json:"message"`
}

type wormlyAlert struct {
	Hostid         int            `json:"hostid"`
	Name           string         `json:"name"`
	IsRecovery     int            `json:"isrecovery"`
	Downtime       int            `json:"downtime"`
	AlertLevel     int            `json:"alertlevel"`
	AlertLevelName string         `json:"alertlevel_name"`
	FailedSensors  []failedSensor `json:"failedsensors"`
}

func parseData(r *http.Request) (string, wormlyAlert) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	data := r.PostForm.Get(postParameter)
	if data == "" {
		panic(fmt.Sprintf("No data on POST parameter %s", postParameter))
	}
	alert := wormlyAlert{}

	err = json.Unmarshal([]byte(data), &alert)
	if err != nil {
		panic(err)
	}

	compactData := bytes.Buffer{}
	err = json.Compact(&compactData, []byte(data))
	if err == nil {
		data = compactData.String()
	}

	return data, alert
}

func createMessage(a wormlyAlert) string {
	s := bytes.Buffer{}
	tpl := `{{ if eq .IsRecovery 0 }}ERROR{{ else }}RECOVERY{{ end }} on host {{ .Name }}
{{ range .FailedSensors }}Sensor {{ .Type }} - {{ .Message }}
{{ end }}
{{- if eq .IsRecovery 1 }}Duration {{ .Downtime }} secs{{- end }}`

	tmpl, err := template.New("message").Parse(tpl)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&s, a)
	if err != nil {
		panic(err)
	}
	return s.String()
}

// Temporary function to test stuff
func debugStuff(w http.ResponseWriter, d string, a wormlyAlert, m string) {
	fmt.Fprintf(w, "Original: %s\n", d)
	fmt.Fprintf(w, "Alert: %v\n", a)
	fmt.Fprintf(w, "Name: %s\n", a.Name)
	fmt.Fprintf(w, "Name: %s\n", a.Name)
	fmt.Fprintf(w, "Final text:\n%s\n", m)
}

// Wormly HTTP Handler
func Wormly(w http.ResponseWriter, r *http.Request) {
	data, alert := parseData(r)
	log.Printf("EVENT %s: %s", service, data)
	msg := createMessage(alert)
	notify.Send(service, msg)
	debugStuff(w, data, alert, msg)
}
