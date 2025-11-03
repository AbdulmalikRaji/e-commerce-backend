package authenticator

import (
	"os"

	"github.com/supabase-community/auth-go"
)

// New instantiates the *Authenticator.
func New() (auth.Client, error) {
	// Initialise client
	client := auth.New(
		os.Getenv("PROJECT_REFERENCE"),
		os.Getenv("SUPABASE_API_KEY"),
	)
	return client, nil
}
