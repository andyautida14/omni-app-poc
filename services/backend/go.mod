module github.com/andyautida/omni-app-poc/services/backend

go 1.23.2

replace github.com/andyautida/omni-app-poc/lib/db => ../../lib/db

replace github.com/andyautida/omni-app-poc/lib/handler => ../../lib/handler

require (
	github.com/andyautida/omni-app-poc/lib/db v0.0.0-00010101000000-000000000000
	github.com/andyautida/omni-app-poc/lib/handler v0.0.0-00010101000000-000000000000
	github.com/gocraft/dbr/v2 v2.7.7
	github.com/google/uuid v1.6.0
	github.com/sethvargo/go-envconfig v1.1.0
)

require (
	github.com/golang-migrate/migrate/v4 v4.18.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)
