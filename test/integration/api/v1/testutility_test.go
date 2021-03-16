package v1// internal/league/testutility_test.go
import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Do something here.
	fmt.Printf("\033[1;36m%s\033[0m", "> Setting up test conditions\n")
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.
	fmt.Printf("\033[1;36m%s\033[0m", "> Tearing down test conditions\n")
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
	fmt.Printf("\n")
}
