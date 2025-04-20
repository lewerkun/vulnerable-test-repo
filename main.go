package main

import (
	"fmt"
	"log"

	"vuln-test-repo/internal/config"
	"vuln-test-repo/internal/sast"
)

func main() {
	fmt.Println("Vulnerable Test Application Running...")

	// --- Demonstrate Secret Vulnerabilities ---
	fmt.Println("\n--- Secret Leaks ---")
	apiKey := config.GetAPIKey()
	dbPass := config.GetDBPassword()
	fmt.Printf("Retrieved API Key (Leak 1): %s\n", apiKey)     // Secret scanner should find this constant
	fmt.Printf("Retrieved DB Password (Leak 2): %s\n", dbPass) // Secret scanner should find this constant

	// --- Demonstrate SAST Vulnerabilities ---
	fmt.Println("\n--- SAST Examples ---")

	// 1. SQL Injection Example
	fmt.Println("\n* SQL Injection Example *")
	// Simulate user input that could be malicious
	maliciousUserID := "' OR '1'='1"         // Classic SQLi payload
	sast.QueryUserData(nil, maliciousUserID) // Pass nil as DB is not connected

	// 2. Command Injection Example
	fmt.Println("\n* Command Injection Example *")
	// Simulate user input providing arguments
	// A safe argument might be "example.com"
	// A malicious argument could be "; ls -la" or similar
	commandToRun := "ping"
	// maliciousArgs := "example.com; id" // Example payload
	maliciousArgs := "-c 1 localhost; echo 'Command Injection Successful!'" // More direct payload
	output, err := sast.ExecuteCommand(commandToRun, maliciousArgs)
	if err != nil {
		log.Printf("Command execution failed (as expected with injection attempt sometimes): %v", err)
	}
	fmt.Printf("Command Output:\n%s\n", string(output))

	fmt.Println("\nApplication finished.")
}
