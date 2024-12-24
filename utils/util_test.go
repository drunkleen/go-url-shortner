package utils

import (
	"testing"
)

func TestGenerateUUIDFromIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
	}{
		{"Valid IP", "192.168.1.1"},
		{"Empty IP", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("GenerateUUIDFromIP() = %v, want no panics", r)
				}
			}()
			got := GenerateUUIDFromIP(tt.ip)
			if got == "" {
				t.Errorf("GenerateUUIDFromIP() = %v, want non-empty string", got)
			}
		})
	}
}

func TestGenerateUUIDFromIPUniqueness(t *testing.T) {
	ip1 := "192.168.1.1"
	ip2 := "192.168.1.2"
	uuid1 := GenerateUUIDFromIP(ip1)
	uuid2 := GenerateUUIDFromIP(ip2)
	if uuid1 == uuid2 {
		t.Errorf("GenerateUUIDFromIP() = %v, want unique UUIDs for different IP addresses", uuid1)
	}
}

func TestGenerateUUIDFromIPConsistency(t *testing.T) {
	ip := "192.168.1.1"
	uuid1 := GenerateUUIDFromIP(ip)
	uuid2 := GenerateUUIDFromIP(ip)
	if uuid1 != uuid2 {
		t.Errorf("GenerateUUIDFromIP() = %v, want same UUID for same IP address", uuid1)
	}
}
