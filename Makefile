phony: all dependencies

# Default variables
MAKE_INCLUDE_PATH = make-deps
SRCS = $(MAKE_INCLUDE_PATH)/*.mk
MAKEFILE_PROJECT_ID = 372
GITLAB_TOKEN = $(shell cat ~/.netrc | grep lab.weave.nl -A 3 | grep password | sed 's/^.* //')
GOBIN = $(GOPATH)/bin

# Download default make steps specified at: https://lab.weave.nl/devops/make
dependencies:
	@rm -rf $(MAKE_INCLUDE_PATH)
	@mkdir -p $(MAKE_INCLUDE_PATH)
	@curl --header "PRIVATE-TOKEN: $(GITLAB_TOKEN)" https://lab.weave.nl/api/v4/projects/$(MAKEFILE_PROJECT_ID)/repository/tree?ref=master > tempmakefiles.json
	@jq -c -r '.[].name' tempmakefiles.json | while read i; do \
		if [[ $$i == *?.mk ]]; then \
			curl --header 'PRIVATE-TOKEN: $(GITLAB_TOKEN)' "https://lab.weave.nl/api/v4/projects/$(MAKEFILE_PROJECT_ID)/repository/files/$$i?ref=master" | jq -r .content | base64 -D > $(MAKE_INCLUDE_PATH)/$$i; \
		fi; \
	done
	@rm tempmakefiles.json

-include $(SRCS)