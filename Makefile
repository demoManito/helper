user :=	$(shell whoami)

# check GOBIN
ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
	# check GOPATH
	ifneq ($(GOPATH),)
		BIN=$(GOPATH)/bin
	endif
endif

cm:
	@cd cmd/command && go build -o helper && cd - &> /dev/null

wf:
	@cd cmd/workflow && go build -o helper && cd - &> /dev/null

.PHONY: install-cm
install-cm: cm
ifeq ($(user),root)
#root, install for all user
	@cp ./cmd/command/helper /usr/bin
else
#!root, install for current user
	$(shell if [ -z $(BIN) ]; then read -p "Please select installdir: " REPLY; mkdir -p $${REPLY};\
	cp ./cmd/command/helper $${REPLY}/;else mkdir -p $(BIN);\
	cp ./cmd/command/helper $(BIN); fi)
	@echo "to path: $(BIN)"
endif
	@echo "install bin finished"

.PHONY: install-workflow
install-workflow: wf
# TODO: install to workflow
	@echo "install workflow finished"

.PHONY: install
install: install-cm install-workflow
	@echo "install bin and workflow finished"