package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	ignore "github.com/sabhiram/go-gitignore"
)

func captureStdout(_ *testing.T, f func(writer *bytes.Buffer)) string {
	var buf bytes.Buffer
	f(&buf)
	return buf.String()
}

func TestPrintTree_FilteringAndDepth(t *testing.T) {
	tempDir := t.TempDir()

	// tempDir/
	// ├── file.txt
	// ├── .hidden_file
	// ├── node_modules/ (Ignorado por Git)
	// │   └── dep.js
	// └── src/
	//     └── index.go

	os.Mkdir(filepath.Join(tempDir, "src"), 0755)
	os.Mkdir(filepath.Join(tempDir, "node_modules"), 0755)
	os.WriteFile(filepath.Join(tempDir, "file.txt"), []byte("data"), 0644)
	os.WriteFile(filepath.Join(tempDir, ".hidden_file"), []byte("data"), 0644)
	os.WriteFile(filepath.Join(tempDir, "src", "index.go"), []byte("data"), 0644)
	os.WriteFile(filepath.Join(tempDir, "node_modules", "dep.js"), []byte("data"), 0644)

	os.WriteFile(filepath.Join(tempDir, ".gitignore"), []byte("node_modules"), 0644)

	ignoreObj, _ := ignore.CompileIgnoreFile(filepath.Join(tempDir, ".gitignore"))

	tests := []struct {
		name                string
		maxDepth            int
		showHidden          bool
		expectedLines       int
		expectedContains    string
		expectedNotContains string
	}{
		{
			name:                "Defecto_sin_flags",
			maxDepth:            0,
			showHidden:          false,
			expectedLines:       6,
			expectedContains:    "file.txt",
			expectedNotContains: ".hidden_file",
		},
		{
			name:                "Mostrar_Ocultos_-a",
			maxDepth:            0,
			showHidden:          true,
			expectedLines:       7,
			expectedContains:    ".hidden_file",
			expectedNotContains: "node_modules",
		},
		{
			name:                "Límite_Profundidad_-L_1",
			maxDepth:            1,
			showHidden:          false,
			expectedLines:       3,
			expectedContains:    "src/",
			expectedNotContains: "index.go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxDepth = tt.maxDepth
			showHidden = tt.showHidden
			directoriesOnly = false

			output := captureStdout(t, func(buf *bytes.Buffer) {
				fmt.Fprintf(buf, "%s\n", tempDir)
				printTree(tempDir, tempDir, "", ignoreObj, 0, buf)
			})

			if !strings.Contains(output, tt.expectedContains) {
				t.Errorf("FAIL: La salida NO contiene la cadena esperada '%s'. Salida:\n%s", tt.expectedContains, output)
			}

			if strings.Contains(output, tt.expectedNotContains) {
				t.Errorf("FAIL: La salida contiene la cadena NO esperada '%s'. Salida:\n%s", tt.expectedNotContains, output)
			}

			if tt.maxDepth == 1 {
				lines := strings.Split(strings.TrimSpace(output), "\n")
				actualLines := len(lines)
				if len(lines) != tt.expectedLines {
					t.Errorf("FAIL: Se espraba %d linesa, se obtuvo %d. Salida:\n%s", tt.expectedLines, actualLines, output)
				}
			}
		})
	}
}
