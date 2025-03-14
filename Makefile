.PHONY: test test-watch
test:
	go test -vet=all ./... 

test-watch:
	gow -c test ./...

.PHONY: build
build:
	go build ./...

.PHONY: watch
watch:
	gow -s -S "Build done" build

SOURCES_DIR := internal/specs/sources/ed
ED_DIR := $(SOURCES_DIR)
IDL_DIR := $(ED_DIR)/idlparsed
ELEMENTS_DIR := $(ED_DIR)/elements
EVENTS_DIR := $(ED_DIR)/events

OUTPUT_DIR := internal/specs/curated

# If you add a new folder to build, include it here.
SOURCE_JSON_FILES := $(wildcard $(IDL_DIR)/*.json) $(wildcard $(ELEMENTS_DIR)/*.json) $(wildcard $(EVENTS_DIR)/*.json)


# The jq transformation command
JQ_CMD := jq -c 'del(.. | .fragment? // empty, .href? // empty)'

TARGET_JSON_FILES := $(SOURCE_JSON_FILES:$(ED_DIR)/%=$(OUTPUT_DIR)/%)

specs: $(TARGET_JSON_FILES)

diag:
	echo $(TARGET_JSON_FILES)

# Rule to process each JSON file
$(OUTPUT_DIR)/%.json: $(ED_DIR)/%.json
	@echo "Processing $@..."
	$(JQ_CMD) $< > $@

.PHONY: release
release: build test
	pnpm release

