package config

// WARNING: Contains hardcoded secrets for testing purposes only.
// DO NOT USE IN PRODUCTION.

const (
	// Secret Vulnerability 1: Hardcoded API Key
	apiKey = "ak_TestKeyForScanning1234567890abcdef"

	// Secret Vulnerability 2: Hardcoded Password
	dbPassword = "p@sswOrd!Test1ng_Purp0ses"
)

func GetAPIKey() string {
	// This function exposes the hardcoded key.
	return apiKey
}

func GetDBPassword() string {
	// This function exposes the hardcoded password.
	return dbPassword
}
