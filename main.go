package main

import (
	"fmt"
	"log"

	"vuln-test-repo/internal/config"
	"vuln-test-repo/internal/sast"
)

func main() {
	fmt.Println("Vulnerable Test Application Running (Explicit Vulnerabilities)...")

	// --- Demonstrate Obvious Secret Vulnerabilities ---
	fmt.Println("\n--- Obvious Secret Leaks ---")
	apiKey := config.GetAPIKey()
	dbPass := config.GetDBPassword()
	// Explicitly printing potentially sensitive info
	fmt.Printf("Retrieved SUPER_SECRET_API_KEY_DO_NOT_COMMIT (Leak 1): %s\n", apiKey)
	fmt.Printf("Retrieved DATABASE_PASSWORD_HARDCODED (Leak 2): %s\n", dbPass)

	// --- Demonstrate Obvious SAST Vulnerabilities ---
	fmt.Println("\n--- Obvious SAST Examples ---")

	// 1. Obvious SQL Injection Example
	fmt.Println("\n* Obvious SQL Injection Example *")
	maliciousSQLInput := "' OR '1'='1" // Classic SQLi payload
	fmt.Printf("Simulating SQLi with input: %s\n", maliciousSQLInput)
	sast.QueryUserData(nil, maliciousSQLInput) // Pass nil as DB is not connected

	// 2. Obvious Command Injection Example
	fmt.Println("\n* Obvious Command Injection Example *")
	// User input is the entire command string
	maliciousCommandInput := "ls -la / || id" // Classic command injection payload
	fmt.Printf("Simulating Command Injection with input: %s\n", maliciousCommandInput)
	output, err := sast.ExecuteCommand(maliciousCommandInput)
	if err != nil {
		// Error is expected if injection causes non-zero exit code, but output might still exist
		log.Printf("Command execution potentially failed or partially succeeded (err: %v)", err)
	}
	fmt.Printf("Command Output:\n%s\n", string(output))

	fmt.Println("\nApplication finished.")
}
