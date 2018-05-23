package getafix_philter

import (
	"fmt"
	"time"
)

type Eumenes struct {
	app     *string
	project *string
	server  *string
	version *string
}

func NewEumenes(app *string, project *string, server *string, version *string) Eumenes {
	return Eumenes{app: app, project: project, server: server, version: version}
}

func (g *Eumenes) string(level string, message *string, logTime time.Time) string {
	formatTime := "2018-05-07T15:38:01.841442892+07:00"
	return fmt.Sprintf(`{"app": "%s", "level": "%s", "log_time": "%s", "msg": "%s", "project": "%s", "server": "%s", "version": "%s"}`,
		*g.app, level, logTime.Format(formatTime), *message, *g.project, *g.server, *g.version)
}

func (g *Eumenes) Info(message *string, time time.Time) {
	fmt.Println(g.string("info", message, time))
}

func (g *Eumenes) Warning(message *string, time time.Time) {
	fmt.Println(g.string("warning", message, time))
}

func (g *Eumenes) Error(message *string, time time.Time) {
	fmt.Println(g.string("error", message, time))
}
