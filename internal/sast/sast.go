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

// SAST Vulnerability 1: SQL Injection
// QueryUserData constructs an SQL query unsafely using string concatenation.
func QueryUserData(db *sql.DB, userID string) {
	// This is vulnerable to SQL Injection because userID is directly concatenated.
	query := "SELECT * FROM users WHERE id = '" + userID + "'"
	fmt.Println("Executing query:", query) // Simulate execution

	// Example of how it might be used (commented out as DB is not connected)
	// rows, err := db.Query(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// ... process rows ...
}

// SAST Vulnerability 2: Command Injection
// ExecuteCommand executes a system command unsafely using user input.
func ExecuteCommand(command string, args string) ([]byte, error) {
	// This is vulnerable to Command Injection because args are passed directly.
	cmdPath, err := exec.LookPath(command)
	if err != nil {
		log.Printf("Command not found: %s", command)
		return nil, err
	}

	fmt.Printf("Executing command: %s %s\n", cmdPath, args) // Simulate execution

	// This line allows command injection via the 'args' parameter.
	cmd := exec.Command(cmdPath, args) // Vulnerable line

	// For demonstration, let's just capture output. In a real scenario, this might modify the system.
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command: %v", err)
		return output, err
	}
	return output, nil
}
