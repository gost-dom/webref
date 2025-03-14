module github.com/gost-dom/webref

go 1.23.4

retract (
	v0.3.1 // Was published a little hastily with poor naming. Skip this and jump to 0.4
	v1.0.1 // Contains retractions only.
	v1.0.0 // Published accidentally. This is identical to 0.1.0
)

require (
	github.com/onsi/gomega v1.36.2
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/onsi/ginkgo/v2 v2.22.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
