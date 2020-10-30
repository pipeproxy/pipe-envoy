.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 		:= github.com/pipeproxy/pipe-xds

.PHONY: $(BINDIR)/pipe-xds
$(BINDIR)/pipe-xds:
	@go build -o $@ ./cmd/pipe-xds


.PHONY: format
format:
	@goimports -local $(PKG) -w -l pkg

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test integration integration.ads.v3 integration.ads.v3.tls

$(BINDIR)/test:
	@go build -o $@ test/main.go

integration: integration.ads.v3 integration.ads.v3.tls

integration.ads.v3: $(BINDIR)/pipe-xds $(BINDIR)/test
	env XDS=ads SUFFIX=v3 build/integration.sh

integration.ads.v3.tls: $(BINDIR)/pipe-xds $(BINDIR)/test
	env XDS=ads SUFFIX=v3 build/integration.sh -tls
