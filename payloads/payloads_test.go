package payloads

// todo: replace tests with testify

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name  string
		lhost string
		lport int
	}{
		{
			name:  "bash_tcp_1",
			lhost: "127.0.0.1",
			lport: 8080,
		},
		{
			name:  "python_1",
			lhost: "192.168.1.100",
			lport: 4444,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Get(tc.name, tc.lhost, tc.lport)

			// Test that the result is not empty
			if result == "" {
				t.Fatal("Get() returned empty string")
			}

			// Test that SHELLDROP_HOST is replaced with actual host
			if !strings.Contains(result, tc.lhost) {
				t.Errorf("Expected result to contain host '%s', got: %s", tc.lhost, result)
			}
			if strings.Contains(result, "SHELLDROP_HOST") {
				t.Error("Result still contains SHELLDROP_HOST placeholder")
			}

			// Test that SHELLDROP_PORT is replaced with actual port
			portStr := strconv.Itoa(tc.lport)
			if !strings.Contains(result, portStr) {
				t.Errorf("Expected result to contain port '%s', got: %s", portStr, result)
			}
			if strings.Contains(result, "SHELLDROP_PORT") {
				t.Error("Result still contains SHELLDROP_PORT placeholder")
			}
		})
	}
}

func TestGetUrlEncoded(t *testing.T) {
	testCases := []struct {
		name  string
		lhost string
		lport int
	}{
		{
			name:  "bash_tcp_1",
			lhost: "127.0.0.1",
			lport: 8080,
		},
		{
			name:  "python_1",
			lhost: "192.168.1.100",
			lport: 4444,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			regular := Get(tc.name, tc.lhost, tc.lport)
			encoded := GetUrlEncoded(tc.name, tc.lhost, tc.lport)

			if encoded == "" {
				t.Fatal("GetUrlEncoded() returned empty string")
			}

			decoded, err := url.QueryUnescape(encoded)
			if err != nil {
				t.Fatalf("Failed to decode URL encoded payload: %v", err)
			}
			if decoded != regular {
				t.Errorf("URL decoded payload doesn't match regular payload.\nRegular: %s\nDecoded: %s", regular, decoded)
			}
		})
	}
}

func TestPlaceholderReplacement(t *testing.T) {
	name := "bash_tcp_1"
	lhost := "10.0.0.1"
	lport := 9999

	result := Get(name, lhost, lport)

	// Count occurrences of host and port in result
	hostCount := strings.Count(result, lhost)
	portCount := strings.Count(result, strconv.Itoa(lport))

	// There should be at least one occurrence of each
	if hostCount == 0 {
		t.Error("Host was not replaced in payload")
	}
	if portCount == 0 {
		t.Error("Port was not replaced in payload")
	}

	// Verify no placeholders remain
	if strings.Contains(result, "SHELLDROP_HOST") {
		t.Error("SHELLDROP_HOST placeholder was not fully replaced")
	}

	if strings.Contains(result, "SHELLDROP_PORT") {
		t.Error("SHELLDROP_PORT placeholder was not fully replaced")
	}
}

func TestMultipleReplacements(t *testing.T) {
	testPayload := "connect to SHELLDROP_HOST:SHELLDROP_PORT and also SHELLDROP_HOST on port SHELLDROP_PORT"
	payloads["test_multiple"] = testPayload

	lhost := "example.com"
	lport := 1337

	result := Get("test_multiple", lhost, lport)

	expectedHostCount := 2
	expectedPortCount := 2

	actualHostCount := strings.Count(result, lhost)
	actualPortCount := strings.Count(result, strconv.Itoa(lport))

	if actualHostCount != expectedHostCount {
		t.Errorf("Expected %d host replacements, got %d", expectedHostCount, actualHostCount)
	}

	if actualPortCount != expectedPortCount {
		t.Errorf("Expected %d port replacements, got %d", expectedPortCount, actualPortCount)
	}
}

func TestAllPayloadFilesAreValid(t *testing.T) {
	txtFiles, err := filepath.Glob("*.txt")
	if err != nil {
		t.Fatalf("Failed to find .txt files: %v", err)
	}

	if len(txtFiles) == 0 {
		t.Fatal("No .txt files found in payloads directory")
	}

	for _, filename := range txtFiles {
		t.Run(filename, func(t *testing.T) {
			content, err := os.ReadFile(filename)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", filename, err)
			}

			fileContent := string(content)
			if !strings.Contains(fileContent, "SHELLDROP_HOST") {
				t.Errorf("File %s does not contain SHELLDROP_HOST placeholder", filename)
			}
			if !strings.Contains(fileContent, "SHELLDROP_PORT") {
				t.Errorf("File %s does not contain SHELLDROP_PORT placeholder", filename)
			}
			if strings.TrimSpace(fileContent) == "" {
				t.Errorf("File %s is empty", filename)
			}
			if strings.Contains(fileContent, "\n") {
				t.Errorf("File %s should not contain newlines", filename)
			}
		})
	}
}
