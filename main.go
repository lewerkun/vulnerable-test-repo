package main

import (
	"fmt"
	"log"

	"vuln-test-repo/internal/config"
	"vuln-test-repo/internal/sast"
)

func main() {
	fmt.Println("Vulnerable Test Application Running (Multiple Vulnerabilities)...")

	// --- Demonstrate Secret Vulnerabilities ---
	fmt.Println("\n--- Secret Leaks ---")
	apiKey := config.GetAPIKey()
	dbPass := config.GetDBPassword()
	fmt.Printf("Retrieved SUPER_SECRET_API_KEY_DO_NOT_COMMIT: %s\n", apiKey)
	fmt.Printf("Retrieved DATABASE_PASSWORD_HARDCODED: %s\n", dbPass)

	// --- Demonstrate Multiple SAST Vulnerabilities ---
	fmt.Println("\n--- SAST Examples ---")

	// 1. SQL Injection
	fmt.Println("\n* SQL Injection Example *")
	maliciousSQLInput := "' OR '1'='1"
	fmt.Printf("Simulating SQLi with input: %s\n", maliciousSQLInput)
	sast.QueryUserData(nil, maliciousSQLInput)

	// 2. Command Injection
	fmt.Println("\n* Command Injection Example *")
	maliciousCommandInput := "ls -la / || id"
	fmt.Printf("Simulating Command Injection with input: %s\n", maliciousCommandInput)
	output, err := sast.ExecuteCommand(maliciousCommandInput)
	if err != nil {
		log.Printf("Command execution error: %v", err)
	}
	fmt.Printf("Command Output:\n%s\n", string(output))

	// 3. Weak Cryptography
	fmt.Println("\n* Weak Cryptography Example *")
	password := "mySecurePassword123"
	hashedPassword := sast.WeakHash(password)
	fmt.Printf("MD5 hash of '%s': %s\n", password, hashedPassword)

	// 4. Hardcoded Credentials
	fmt.Println("\n* Hardcoded Credentials Example *")
	sast.ConnectToDatabase()

	// 5. Auth Bypass
	fmt.Println("\n* Authentication Bypass Example *")
	fmt.Printf("Authentication with empty credentials: %v\n", sast.CheckCredentials("", ""))

	// 6. XSS Vulnerability (not executed, just shown)
	fmt.Println("\n* XSS Vulnerability Example *")
	maliciousXSS := "<script>alert('XSS')</script>"
	fmt.Printf("Vulnerable to XSS with input: %s\n", maliciousXSS)
	// In a web server, this would be: sast.BuildResponse(w, maliciousXSS)

	fmt.Println("\nApplication finished - Multiple vulnerabilities available for scanning.")

	// Start a minimal HTTP server to demonstrate some vulnerabilities
	// Uncomment if you want to run it
	/*
		http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
			sast.HandleRedirect(w, r)
		})
		http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
			data, _ := sast.GetFile(r)
			w.Write(data)
		})
		http.HandleFunc("/xss", func(w http.ResponseWriter, r *http.Request) {
			userInput := r.URL.Query().Get("input")
			sast.BuildResponse(w, userInput)
		})
		fmt.Println("Starting server on :8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
