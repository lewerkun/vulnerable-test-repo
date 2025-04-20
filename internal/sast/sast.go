package sast

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"unsafe"
	// Import a sql driver if you were actually connecting
	// _ "github.com/go-sql-driver/mysql"
)

// WARNING: Contains SAST vulnerabilities for testing purposes only.
// DO NOT USE IN PRODUCTION.

// SAST Vulnerability 1: SQL Injection (multiple patterns)
func QueryUserData(db *sql.DB, userInput string) {
	// Using string concatenation (owasp-top-ten:A03:2021-injection)
	query1 := "SELECT * FROM users WHERE name = '" + userInput + "'"
	fmt.Println("Vulnerable Query 1:", query1)

	// Using fmt.Sprintf (security-audit)
	query2 := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userInput)
	fmt.Println("Vulnerable Query 2:", query2)

	// Directly in db.Query (uncommented)
	if db != nil {
		rows, err := db.Query("SELECT * FROM users WHERE name = '" + userInput + "'")
		if err != nil {
			log.Println(err)
			return
		}
		defer rows.Close()
	}
}

// SAST Vulnerability 2: Command Injection (already detected)
func ExecuteCommand(userInput string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", userInput) // detected by semgrep
	return cmd.CombinedOutput()
}

// SAST Vulnerability 3: Weak Crypto (MD5)
func WeakHash(password string) string {
	hasher := md5.New()            // MD5 is insecure
	hasher.Write([]byte(password)) // This should be detected
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// SAST Vulnerability 4: Open Redirect
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// Vulnerable open redirect
	url := r.URL.Query().Get("url")
	if url != "" {
		http.Redirect(w, r, url, http.StatusFound) // Redirects to any URL (owasp-top-ten:A01:2021-broken-access-control)
	}
}

// SAST Vulnerability 5: Hardcoded Credentials
func ConnectToDatabase() {
	username := "admin"
	password := "password123" // Hardcoded credentials
	fmt.Println("Connecting with:", username, password)
}

// SAST Vulnerability 6: Directory Traversal
func GetFile(r *http.Request) ([]byte, error) {
	filename := r.URL.Query().Get("file")
	// Directory traversal vulnerability (owasp-top-ten:A01:2021-broken-access-control)
	data, err := os.ReadFile(filename) // No path sanitization
	return data, err
}

// SAST Vulnerability 7: Regex DoS
func ValidateEmail(email string) bool {
	// Regex DoS (ReDoS) vulnerability
	pattern := regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`)
	return pattern.MatchString(email)
}

// SAST Vulnerability 8: Unsafe pointer use
func UnsafePointer(s string) {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(strHeader.Len)
}

// SAST Vulnerability 9: Auth bypass with default credentials
func CheckCredentials(user, pass string) bool {
	if user == "" || pass == "" {
		return true // Default access (owasp-top-ten:A07:2021-authentication-failures)
	}
	return false
}

// SAST Vulnerability 10: Insecure deserialization
func DecodeBase64(input string) string {
	decoded, _ := base64.StdEncoding.DecodeString(input) // Error ignored
	return string(decoded)
}

// SAST Vulnerability 11: Improper input validation
func ValidateInput(input string) bool {
	// No input validation at all
	return true
}

// SAST Vulnerability 12: XSS - Response splitting
func BuildResponse(w http.ResponseWriter, userInput string) {
	w.Header().Set("X-Custom-Header", userInput) // HTTP response splitting

	// XSS vulnerability (owasp-top-ten:A03:2021-injection)
	fmt.Fprintf(w, "<div>%s</div>", userInput) // No escaping of user input
}
