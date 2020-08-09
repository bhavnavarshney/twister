// +build unit

package csvlog

import (
	"bytes"
	"encoding/csv"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestBuildHeader(t *testing.T) {
	output := buildHeader()
	expected := []string{"Date(YYYY/MM/DD)", "Time",
		"Tool_ID (Hex)", "Calibrated Offset(Dec)", "Current Offset(Dec)",
		"CW Toq[1] (Dec)", "CW AD[1] (Dec)",
		"CW Toq[2] (Dec)", "CW AD[2] (Dec)",
		"CW Toq[3] (Dec)", "CW AD[3] (Dec)",
		"CW Toq[4] (Dec)", "CW AD[4] (Dec)",
		"CW Toq[5] (Dec)", "CW AD[5] (Dec)",
		"CW Toq[6] (Dec)", "CW AD[6] (Dec)",
		"CW Toq[7] (Dec)", "CW AD[7] (Dec)",
		"CW Toq[8] (Dec)", "CW AD[8] (Dec)",
		"CW Toq[9] (Dec)", "CW AD[9] (Dec)",
		"CW Toq[10] (Dec)", "CW AD[10] (Dec)",
		"CW Toq[11] (Dec)", "CW AD[11] (Dec)",
		"CW Toq[12] (Dec)", "CW AD[12] (Dec)",
		"CCW Toq[1] (Dec)", "CCW AD[1] (Dec)",
		"CCW Toq[2] (Dec)", "CCW AD[2] (Dec)",
		"CCW Toq[3] (Dec)", "CCW AD[3] (Dec)",
		"CCW Toq[4] (Dec)", "CCW AD[4] (Dec)",
		"CCW Toq[5] (Dec)", "CCW AD[5] (Dec)",
		"CCW Toq[6] (Dec)", "CCW AD[6] (Dec)",
		"CCW Toq[7] (Dec)", "CCW AD[7] (Dec)",
		"CCW Toq[8] (Dec)", "CCW AD[8] (Dec)",
		"CCW Toq[9] (Dec)", "CCW AD[9] (Dec)",
		"CCW Toq[10] (Dec)", "CCW AD[10] (Dec)",
		"CCW Toq[11] (Dec)", "CCW AD[11] (Dec)",
		"CCW Toq[12] (Dec)", "CCW AD[12] (Dec)",
	}
	assert.Equal(t, expected, output)
}

func TestBuildRow(t *testing.T) {
	lr := LogRecord{Time: time.Now()}
	res := lr.buildRow()
	t.Log(res)
}

func Test_LogRecord_BuildFileName(t *testing.T) {
	logTime, err := time.Parse(time.ANSIC, "Mon Jan 02 15:04:05 2006")
	assert.NoError(t, err)
	lr := &LogRecord{
		Time: logTime,
		Type: "NPT12",
	}
	fileName := lr.buildFileName()
	assert.Equal(t, "NPT12_20060102.csv", fileName)
}

func Test_Logger_Write_CreatesFile(t *testing.T) {
	lr := &LogRecord{Time: time.Now()}
	fileName := lr.buildFileName()
	fs := afero.NewMemMapFs()
	l := Logger{
		FS: fs,
	}
	err := l.Write(lr)
	assert.NoError(t, err)
	exists, err := afero.Exists(fs, fileName)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func Test_Logger_Write_CreatesFileInCustomDirectory(t *testing.T) {
	lr := &LogRecord{Time: time.Now(), Type: "NPT12"}
	filePath := path.Join("logs", lr.buildFileName())
	fs := afero.NewMemMapFs()
	l := Logger{
		FS:  fs,
		Dir: "logs",
	}
	err := l.Write(lr)
	assert.NoError(t, err)
	exists, err := afero.Exists(fs, filePath)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func Test_Logger_Write_CreatesNestedDir(t *testing.T) {
	lr := &LogRecord{Time: time.Now()}
	filePath := path.Join("logs/file/tmax/" + lr.buildFileName())
	fs := afero.NewMemMapFs()
	l := Logger{
		FS:  fs,
		Dir: "logs/file/tmax/",
	}
	err := l.Write(lr)
	assert.NoError(t, err)
	exists, err := afero.Exists(fs, filePath)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func Test_Logger_Write_WritesHeader(t *testing.T) {
	lr := &LogRecord{Time: time.Now()}
	fileName := lr.buildFileName()
	fs := afero.NewMemMapFs()
	l := Logger{
		FS: fs,
	}
	err := l.Write(lr)
	assert.NoError(t, err)
	exists, err := afero.Exists(fs, fileName)
	assert.NoError(t, err)
	assert.True(t, exists)
	file, err := afero.ReadFile(fs, fileName)
	assert.NoError(t, err)
	t.Log(string(file))
}

func Test_Logger_Write_AppendsToExistingFile(t *testing.T) {
	lr := &LogRecord{Time: time.Now()}
	fileName := lr.buildFileName()
	fs := afero.NewMemMapFs()
	l := Logger{
		FS: fs,
	}

	// Write to the file twice
	err := l.Write(lr)
	assert.NoError(t, err)
	err = l.Write(lr)
	assert.NoError(t, err)

	file, err := afero.ReadFile(fs, fileName)
	assert.NoError(t, err)
	timesWritten := strings.Count(string(file), convertToCSV(lr.buildRow()))
	t.Log(string(file))
	assert.Equal(t, 2, timesWritten)
}

// Takes an array of fields and returns a comma separated string
func convertToCSV(input []string) string {
	csvOut := bytes.NewBuffer([]byte{})
	w := csv.NewWriter(csvOut)
	err := w.Write(input)
	if err != nil {
		panic(err)
	}
	w.Flush()
	return csvOut.String()
}
