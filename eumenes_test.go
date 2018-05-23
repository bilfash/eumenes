package eumenes_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/bilfash/eumenes"
	"github.com/stretchr/testify/assert"
)

var (
	appDummy     = "logger_dummy"
	projectDummy = "project_dummy"
	serverDummy  = "server_dummy"
	version      = "version"
	msg          = "Log message"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestEumenes_Info(t *testing.T) {
	logger := eumenes.NewEumenes(&appDummy, &projectDummy, &serverDummy, &version)
	dummyTime := time.Now()
	formattedTime := dummyTime.Format("2018-05-07T15:38:01.841442892+07:00")
	result := captureOutput(func() {
		logger.Info(&msg, dummyTime)
	})
	expected := fmt.Sprintf(`{"app": "%s", "level": "info", "log_time": "%s", "msg": "Log message", "project": "%s", "server": "%s", "version": "%s"}
`,
		appDummy, formattedTime, projectDummy, serverDummy, version)
	assert.Equal(t, expected, result, "should be equal")
}

func TestEumenes_Warning(t *testing.T) {
	logger := eumenes.NewEumenes(&appDummy, &projectDummy, &serverDummy, &version)
	dummyTime := time.Now()
	formattedTime := dummyTime.Format("2018-05-07T15:38:01.841442892+07:00")
	result := captureOutput(func() {
		logger.Warning(&msg, dummyTime)
	})
	expected := fmt.Sprintf(`{"app": "%s", "level": "warning", "log_time": "%s", "msg": "Log message", "project": "%s", "server": "%s", "version": "%s"}
`,
		appDummy, formattedTime, projectDummy, serverDummy, version)
	assert.Equal(t, expected, result, "should be equal")
}

func TestEumenes_Error(t *testing.T) {
	logger := eumenes.NewEumenes(&appDummy, &projectDummy, &serverDummy, &version)
	dummyTime := time.Now()
	formattedTime := dummyTime.Format("2018-05-07T15:38:01.841442892+07:00")
	result := captureOutput(func() {
		logger.Error(&msg, dummyTime)
	})
	expected := fmt.Sprintf(`{"app": "%s", "level": "error", "log_time": "%s", "msg": "Log message", "project": "%s", "server": "%s", "version": "%s"}
`,
		appDummy, formattedTime, projectDummy, serverDummy, version)
	assert.Equal(t, expected, result, "should be equal")
}
