package config

// WARNING: Contains hardcoded secrets for testing purposes only.
// DO NOT USE IN PRODUCTION.

const (
	// SECRET LEAK 1: Obvious hardcoded API Key
	// gitleaks:allow -- used for testing scanners
	SUPER_SECRET_API_KEY_DO_NOT_COMMIT = "ghp_ThisIsAFakeGitHubTokenForTestingScan123"

	// SECRET LEAK 2: Obvious hardcoded Password
	// gitleaks:allow -- used for testing scanners
	DATABASE_PASSWORD_HARDCODED = "AdminP@ssword123!_ThisShouldBeFound"
)

// GetAPIKey exposes the hardcoded key.
// It's intentionally insecure.
func GetAPIKey() string {
	// This function explicitly returns the hardcoded key.
	return SUPER_SECRET_API_KEY_DO_NOT_COMMIT // Secret returned here
}

// GetDBPassword exposes the hardcoded password.
// It's intentionally insecure.
func GetDBPassword() string {
	// This function explicitly returns the hardcoded password.
	return DATABASE_PASSWORD_HARDCODED // Secret returned here
}
