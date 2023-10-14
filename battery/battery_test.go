package battery

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestReadBatteryStatus(t *testing.T) {
	// Setup a temp directory with mock battery status
	tmpDir := os.TempDir()
	defer os.RemoveAll(tmpDir)

	statusFile := filepath.Join(tmpDir, "status")
	os.WriteFile(statusFile, []byte("Charging\\n"), os.ModePerm)

	// Test the function
	status := ReadBatteryStatus(tmpDir)
	if status != "Charging" {
		t.Errorf("Expected Charging, got %s", status)
	}

	// Test other battery statuses
	statuses := []string{"Discharging", "Full", "Unknown"}
	for _, s := range statuses {
		os.WriteFile(statusFile, []byte(s + "\\n"), os.ModePerm)
		status := ReadBatteryStatus(tmpDir)
		if status != s {
			t.Errorf("Expected %s, got %s", s, status)
		}
	}
}

func TestReadBatteryLevel(t *testing.T) {
	// Setup a temp directory with mock battery level
	tmpDir := os.TempDir()
	defer os.RemoveAll(tmpDir)

	levelFile := filepath.Join(tmpDir, "capacity")
	os.WriteFile(levelFile, []byte("80\\n"), os.ModePerm)

	// Test the function
	level := ReadBatteryLevel(tmpDir)
	if level != 80 {
		t.Errorf("Expected 80, got %d", level)
	}

	// Test other battery levels
	for i := 0; i <= 100; i++ {
		os.WriteFile(levelFile, []byte(strconv.Itoa(i) + "\\n"), os.ModePerm)
		level := ReadBatteryLevel(tmpDir)
		if level != uint(i) {
			t.Errorf("Expected %d, got %d", i, level)
		}
	}
}


func TestReadBatteryStatus_Error(t *testing.T) {
	// Test reading from a non-existent directory
	status := ReadBatteryStatus("/non/existent/dir")
	if status != "" {
		t.Errorf("Expected an empty string, got %s", status)
	}
}

func TestReadBatteryLevel_Error(t *testing.T) {
	// Test reading from a non-existent directory
	level := ReadBatteryLevel("/non/existent/dir")
	if level != 0 {
		t.Errorf("Expected 0, got %d", level)
	}
}
