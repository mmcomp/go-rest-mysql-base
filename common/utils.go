package common

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	const cost = 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(hash), "$2a$", "$2y$", 1), nil
}

// escapeSQLLikeWildcards escapes SQL LIKE wildcard characters (% and _) and the escape character itself (\)
// to prevent them from being interpreted as wildcards in LIKE queries.
func EscapeSQLLikeWildcards(s string) string {
	// Escape backslash first, then escape wildcards
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "%", "\\%")
	s = strings.ReplaceAll(s, "_", "\\_")
	return s
}
