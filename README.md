# Vulnerable Test Repository

This repository is intentionally created with vulnerabilities for testing security scanning tools.

**DO NOT USE ANY CODE FROM THIS REPOSITORY IN PRODUCTION.**

## Vulnerabilities Included:

### SAST Vulnerabilities

1.  **SQL Injection:** Located in `internal/sast/sast.go` in the `QueryUserData` function.
2.  **Command Injection:** Located in `internal/sast/sast.go` in the `ExecuteCommand` function.

### Secret Vulnerabilities

1.  **Hardcoded API Key:** Located in `internal/config/config.go`.
2.  **Hardcoded Password:** Located in `internal/config/config.go`. 