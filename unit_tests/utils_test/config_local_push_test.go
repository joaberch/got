package utils

import (
	"encoding/json"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigLocalPush_Success(t *testing.T) {
	tempDir := t.TempDir()
	gotDir := filepath.Join(tempDir, ".got")
	err := os.MkdirAll(gotDir, os.ModePerm)
	if err != nil {
		t.Errorf("error creating directory %s: %v", gotDir, err)
	}

	remotePath := filepath.Join(tempDir, "remote")
	err = os.MkdirAll(remotePath, os.ModePerm)
	if err != nil {
		t.Errorf("error creating directory %s: %v", remotePath, err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Errorf("error getting current working directory: %v", err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Errorf("error changing working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Errorf("error changing working directory: %v", err)
		}
	}()

	err = utils.ConfigLocalPush("remote")
	if err != nil {
		t.Errorf("error configuring local push: %v", err)
	}

	configFile := filepath.Join(".got", ".gotconfig")
	data, err := os.ReadFile(configFile)
	if err != nil {
		t.Errorf("error reading config file: %v", err)
	}

	var config map[string]string
	err = json.Unmarshal(data, &config)
	if err != nil {
		t.Errorf("error unmarshalling config file: %v", err)
	}
	if config["local"] != "remote" {
		t.Errorf("Expected 'remote', got %s", config["local"])
	}
}

func TestConfigLocalPush_NoPath(t *testing.T) {
	tempDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Errorf("error getting current working directory: %v", err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Errorf("error changing working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Errorf("error changing working directory: %v", err)
		}
	}()

	err = utils.ConfigLocalPush("not_exist")
	if err == nil {
		t.Errorf("no error copying file %s, one was expected", tempDir)
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("error copying file %s, one was expected", err.Error())
	}
}

func TestConfigLocalPush_ReadOnly(t *testing.T) {
	tempDir := t.TempDir()
	gotDir := filepath.Join(tempDir, ".got")
	err := os.MkdirAll(gotDir, 0555) //read-only
	if err != nil {
		t.Errorf("error creating directory %s: %v", gotDir, err)
	}

	remotePath := filepath.Join(tempDir, "remote")
	err = os.MkdirAll(remotePath, os.ModePerm)
	if err != nil {
		t.Errorf("error creating directory %s: %v", remotePath, err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Errorf("error getting current working directory: %v", err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Errorf("error changing working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Errorf("error changing working directory: %v", err)
		}
	}()

	err = utils.ConfigLocalPush("remote")
	if err != nil {
		t.Errorf("error configuring local push: %v", err)
	}
}
