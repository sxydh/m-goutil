package port_utils

import (
	"testing"
)

func TestGetFreePort(t *testing.T) {
	got, err := GetFreePort()
	if err != nil {
		t.Fatalf("GetFreePort() failed: %v", err)
	} else {
		t.Logf("GetFreePortInRange() = %d", got)
	}
}

func TestGetFreePortInRange(t *testing.T) {
	tests := []struct {
		name  string
		start int
		end   int
	}{
		{"Valid range", 1024, 1100},
		{"Invalid: start > end", 2000, 1000},
		{"Invalid: start < 1", 0, 100},
		{"Invalid: end > 65535", 65530, 70000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFreePortInRange(tt.start, tt.end)
			if err != nil {
				t.Errorf("GetFreePortInRange(%d, %d) failed: %v", tt.start, tt.end, err)
			} else {
				t.Logf("GetFreePortInRange(%d, %d) = %d", tt.start, tt.end, got)
			}
		})
	}
}
