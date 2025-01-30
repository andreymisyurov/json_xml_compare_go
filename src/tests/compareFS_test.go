package tests

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCompareFS(t *testing.T) {
	oldSnapshot := "data/test_snapshot1.txt"
	newSnapshot := "data/test_snapshot2.txt"

	createTestFile(oldSnapshot, `/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version2.yaml
/var/log/orders.log
/Users/baker/pokemon.avi`)

	createTestFile(newSnapshot, `/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version3.yaml
/var/log/orders.log
/Users/baker/bakery-secret.txt`)

	cmd := exec.Command("../ex02/compareFS", "--old", oldSnapshot, "--new", newSnapshot)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Ошибка при запуске compareFS: %v", err)
	}

	expected := `ADDED /Users/baker/recipes/database_version3.yaml
ADDED /Users/baker/bakery-secret.txt
REMOVED /Users/baker/recipes/database_version2.yaml
REMOVED /Users/baker/pokemon.avi`

	actual := strings.TrimSpace(string(output))

	if actual != expected {
		t.Errorf("compareFS output mismatch.\nExpected:\n%s\nGot:\n%s", expected, actual)
	}

	os.Remove(oldSnapshot)
	os.Remove(newSnapshot)
}

func createTestFile(filename string, content string) {
	file, _ := os.Create(filename)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
}
