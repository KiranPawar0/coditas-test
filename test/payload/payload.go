package payload

import "github.com/KiranPawar0/coditas-test/pkg/user/config"

var (
	ValidUser = config.User{
		Name:   "John Doe",
		PAN:    "ABCDE1234F",
		Mobile: "1234567890",
		Email:  "john.doe@example.com",
	}

	InvalidPANUser = config.User{
		Name:   "John Doe",
		PAN:    "INVALIDPAN",
		Mobile: "1234567890",
		Email:  "john.doe@example.com",
	}

	InvalidMobileUser = config.User{
		Name:   "John Doe",
		PAN:    "ABCDE1234F",
		Mobile: "12345",
		Email:  "john.doe@example.com",
	}
)
