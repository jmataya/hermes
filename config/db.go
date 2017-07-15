package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	configFileName  = "db.yml"
	configFolder    = "config"
	errDirNotFound  = "Failed to find folder '%s' in '%s' with error %v"
	errFileNotFound = "Failed to find file '%s' in '%s' with error %v"
	projectRoot     = "hermes"
)

// DB is a configuration structure for database configuration options.
// It is generally loaded from db.yml.
type DB struct {
	Host     string `yaml:"host"`
	Database string `yaml:"name" binding:"required"`
	User     string `yaml:"user" binding:"required"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
}

// NewDB loads db.yml and parses the settings based on the environment
// specified.
func NewDB() (*DB, error) {
	env := os.Getenv("GOENV")
	if env == "" {
		env = "dev"
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to get working directory with error %v", err)
	}

	filename, err := findConfigFile(workingDir)
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read db.yml with error %s", err.Error())
	}

	var configs map[string]DB
	if err := yaml.Unmarshal(yamlFile, &configs); err != nil {
		return nil, fmt.Errorf("Unable to parse db.yml with error %s", err.Error())
	}

	db, ok := configs[env]
	if !ok {
		return nil, fmt.Errorf("DB configuration for env %s not found", env)
	}

	if db.Host == "" {
		db.Host = "localhost"
	}

	if db.SSLMode == "" {
		db.SSLMode = "disable"
	}

	return &db, nil
}

// PostgresDSN generates a Postgres connection string from a valid DB.
func (db DB) PostgresDSN() string {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s sslmode=%s", db.Host, db.Database, db.User, db.SSLMode)
	if db.Password != "" {
		dsn = fmt.Sprintf("%s password=%s", dsn, db.Password)
	}

	return dsn
}

func containsFile(path, filename string, isDir bool) (bool, error) {
	filesAndDirs, err := ioutil.ReadDir(path)
	if err != nil {
		return false, fmt.Errorf("Failed to read directory '%s' with error %v", path, err)
	}

	for _, file := range filesAndDirs {
		if file.IsDir() == isDir && file.Name() == filename {
			return true, nil
		}
	}

	return false, nil
}

func findConfigFile(path string) (string, error) {
	dirFound, err := containsFile(path, configFolder, true)
	if err != nil {
		return "", fmt.Errorf(errDirNotFound, configFolder, path, err)
	}

	if dirFound {
		configPath := filepath.Join(path, configFolder)
		fileFound, err := containsFile(configPath, configFileName, false)
		if err != nil {
			return "", fmt.Errorf(errFileNotFound, configFileName, configPath, err)
		}

		if fileFound {
			return filepath.Join(configPath, configFileName), nil
		}

	}

	if filepath.Base(path) == projectRoot {
		return "", fmt.Errorf("%s not found in current project", configFileName)
	}

	parent := filepath.Dir(path)
	return findConfigFile(parent)
}
