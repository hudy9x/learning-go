package initializers

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	checkInterval = 5 * time.Second // Adjust this according to your preference
)

func DetectFgApp() {
	fmt.Println("Application Tracker started...")
	for {
		activeApp := getActiveApplication()
		fmt.Printf("Active application: %s\n", activeApp)

		time.Sleep(checkInterval)
	}
}

// getActiveApplication retrieves the name of the currently active application using AppleScript.
func getActiveApplication() string {
	cmd := exec.Command("osascript", "-e", `tell application "System Events" to get name of (processes where frontmost is true)`)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// The output may contain additional characters, so we trim them.
	activeApp := strings.TrimSpace(string(out))

	return activeApp
}
