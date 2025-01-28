.PHONY: test
test:
	gow -c test ./...

.PHONY: watch
watch:
	gow -s -S "Build done" build


.PHONY: watch
build-webref:
	# git submodule update --init
	# cd specs/sources && npm i
	# cd specs/sources && npm run curate
	cp specs/sources/curated/idlparsed/* specs/curated/idlparsed
	cp specs/sources/curated/elements/* specs/curated/elements

ED_DIR := specs/sources/ed
IDL_DIR := $(ED_DIR)/idlparsed
ELEMENTS_DIR := $(ELEMENTS_DIR)/elements
OUTPUT_DIR := specs/curated

# Pattern to match all JSON files
JSON_FILES := $(wildcard $(IDL_DIR)/*.json)

# The jq transformation command
JQ_CMD := jq -c 'del(.. | .fragment? // empty, .href? // empty)'

# Default target
all: $(JSON_FILES:$(ED_DIR)/%=$(OUTPUT_DIR)/%)

diag:
	echo $(JSON_FILES:$(ED_DIR)/%=$(OUTPUT_DIR)/%)

# Rule to process each JSON file
$(OUTPUT_DIR)/%.json: $(ED_DIR)/%.json
	@echo "Processing $@..."
	$(JQ_CMD) $< > $@

