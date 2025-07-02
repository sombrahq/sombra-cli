.PHONY: qa
qa:
	@go fmt github.com/sombrahq/sombra-cli/internal/...
	@go fmt github.com/sombrahq/sombra-cli/cmd/...


.PHONY: imports
imports:
	go run github.com/quantumcycle/go-import-checks@latest --config qa/.import.yaml
