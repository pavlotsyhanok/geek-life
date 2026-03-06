package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/asdine/storm/v3"
)

// ConnectStorm Create database connection
func ConnectStorm(dbFilePath string) (*storm.DB, error) {
	var dbPath string

	if dbFilePath != "" {
		info, err := os.Stat(dbFilePath)
		if err == nil && info.IsDir() {
			return nil, errors.New("mentioned DB path is a directory; please specify a file path")
		}

		dbPath = dbFilePath
	} else {
		dbPath = GetEnvStr("DB_FILE", "")
	}

	if dbPath == "" {
		homeDir, err := os.UserHomeDir()
		if err == nil && homeDir != "" {
			dbPath = filepath.Join(homeDir, ".geek-life", "default.db")
		} else {
			tmpFile, tempErr := os.CreateTemp("", "geek-life-default-*.db")
			if tempErr != nil {
				return nil, fmt.Errorf("create temporary database file: %w", tempErr)
			}

			dbPath = tmpFile.Name()
			if closeErr := tmpFile.Close(); closeErr != nil {
				return nil, fmt.Errorf("close temporary database file: %w", closeErr)
			}
		}
	}

	if err := CreateDirIfNotExist(filepath.Dir(dbPath)); err != nil {
		return nil, err
	}

	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, fmt.Errorf("connect embedded database file: %w", err)
	}

	return db, nil
}

// CreateDirIfNotExist creates a directory if not found
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create directory %q: %w", dir, err)
		}
	}

	return nil
}

// UnixToTime create time.Time from string timestamp
func UnixToTime(timestamp string) time.Time {
	parts := strings.Split(timestamp, ".")
	i, err := strconv.ParseInt(parts[0], 10, 64)
	if LogIfError(err, fmt.Sprintf("Could not parse timestamp: %s (using current time instead)", timestamp)) {
		return time.Now()
	}

	return time.Unix(i, 0)
}

// LogIfError logs the error and returns true on Error. think as IfError
func LogIfError(err error, message string) bool {
	if err != nil {
		log.Printf("%s: %v", message, err)

		return true
	}

	return false
}

// FatalIfError logs the error and Exit program on Error
func FatalIfError(err error, message string) {
	if LogIfError(err, message) {
		log.Fatal("FATAL ERROR: Exiting program! - ", message)
	}
}
