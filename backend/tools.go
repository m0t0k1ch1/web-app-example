//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
