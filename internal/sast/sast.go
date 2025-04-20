package sast

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	// Import a sql driver if you were actually connecting
	// _ "github.com/go-sql-driver/mysql"
)

// WARNING: Contains SAST vulnerabilities for testing purposes only.
// DO NOT USE IN PRODUCTION.

// SAST Vulnerability 1: Obvious SQL Injection via Sprintf
// QueryUserData constructs an SQL query VERY unsafely using fmt.Sprintf.
func QueryUserData(db *sql.DB, userInput string) {
	// VERY VULNERABLE: Direct use of user input in Sprintf for SQL query.
	// This is a classic SQL Injection vulnerability (gosec G201).
	// A scanner using gosec rules should definitely flag this line.
	query := fmt.Sprintf("SELECT username, password FROM users WHERE user_id = '%s'", userInput)

	fmt.Println("Executing potentially malicious query:", query)

	// Example of how it might be used (commented out as DB is not connected)
	// _, err := db.Exec(query) // Using Exec instead of Query for variation
	// if err != nil {
	// 	log.Fatal("SQL execution error (expected if injected):", err)
	// }
}

// SAST Vulnerability 2: Obvious Command Injection
// ExecuteCommand executes a system command VERY unsafely.
func ExecuteCommand(userSuppliedCommand string) ([]byte, error) {
	// VERY VULNERABLE: The user input is directly used as the command.
	// This allows arbitrary command execution.
	// A scanner should definitely flag this line.
	fmt.Printf("Executing potentially malicious command: %s\n", userSuppliedCommand)

	// No path lookup, no argument separation. Raw execution.
	// This is extremely dangerous and a clear command injection.
	cmd := exec.Command("sh", "-c", userSuppliedCommand) // Classic injection pattern

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command (expected if injected): %v", err)
		// Returning output even on error for demonstration
	}
	return output, err
}
