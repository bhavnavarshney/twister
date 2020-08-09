package csvlog

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spf13/afero"
)

type Logger struct {
	Dir string // Output directory for logs
	FS  afero.Fs
}

func New(fs afero.Fs, dir string) *Logger {
	if dir == "" {
		dir = "log"
	}
	return &Logger{
		FS:  fs,
		Dir: dir,
	}
}

// LogRecord represents a single line in the CSV log file
type LogRecord struct {
	Time             time.Time
	Type             string
	ID               string
	CurrentOffset    uint16
	CalibratedOffset uint16
	ProfileData      []string
}

func (lr *LogRecord) buildFileName() string {
	return fmt.Sprintf("%s_%d%02d%02d.csv", lr.Type, lr.Time.Year(), lr.Time.Month(), lr.Time.Day())
}

func (lr *LogRecord) buildRow() []string {
	date := fmt.Sprintf("%d/%02d/%02d", lr.Time.Year(), lr.Time.Month(), lr.Time.Day())
	time := fmt.Sprintf("%02d:%02d:%02d", lr.Time.Hour(), lr.Time.Minute(), lr.Time.Second())
	drillID := fmt.Sprintf("0x%s", lr.ID)
	currentOffset := fmt.Sprintf("%d", lr.CurrentOffset)
	calibratedOffset := fmt.Sprintf("%d", lr.CalibratedOffset)
	record := []string{date, time, drillID, calibratedOffset, currentOffset}
	return append(record, lr.ProfileData...)
}

func (l *Logger) Log() {

}

// write writes a log record to the specified file
// If the file doesn't exist, it creates it with the appropriate headers
func (l *Logger) Write(lr *LogRecord) error {
	filePath := path.Join(l.Dir, lr.buildFileName())
	fileExists, err := afero.Exists(l.FS, filePath)
	if err != nil {
		return err
	}
	// Create the dir in case it doesn't exist
	err = l.FS.MkdirAll(l.Dir, os.ModePerm)
	if err != nil {
		return err
	}
	logFile, err := l.FS.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	w := csv.NewWriter(logFile)
	if !fileExists {
		err = w.Write(buildHeader())
		if err != nil {
			return err
		}
	}
	err = w.Write(lr.buildRow())
	w.Flush()
	if err != nil {
		return err
	}
	return nil
}

// buildHeader returns a string slice with the header for the CSV log file
func buildHeader() []string {
	headerFields := []string{"Date(YYYY/MM/DD)", "Time", "Tool_ID (Hex)", "Calibrated Offset(Dec)", "Current Offset(Dec)"}
	for i := 0; i < 12; i++ {
		headerFields = append(headerFields, fmt.Sprintf("CW Torq[%d] (Dec)", i+1))
		headerFields = append(headerFields, fmt.Sprintf("CW AD[%d] (Dec)", i+1))
	}
	for i := 0; i < 12; i++ {
		headerFields = append(headerFields, fmt.Sprintf("CCW Torq[%d] (Dec)", i+1))
		headerFields = append(headerFields, fmt.Sprintf("CCW AD[%d] (Dec)", i+1))
	}
	return headerFields
}
