
SHELL 	:= /bin/bash
BINDIR  := bin

#-----------------
#-- e2e
#-----------------
.PHONY: $(BINDIR)/e2e $(BINDIR)/envoy integration integration.ads integration.xds integration.rest integration.ads.tls

$(BINDIR)/e2e:
	@go build -race -o $@ ./test/e2e

$(BINDIR)/envoy:
	@go build -race -o $@ ./cmd/envoy

e2e: e2e.xds e2e.ads e2e.rest e2e.ads.tls

e2e.ads: $(BINDIR)/e2e $(BINDIR)/envoy
	env XDS=ads hack/e2e.sh

e2e.xds: $(BINDIR)/e2e $(BINDIR)/envoy
	env XDS=xds hack/e2e.sh

e2e.rest: $(BINDIR)/e2e $(BINDIR)/envoy
	env XDS=rest hack/e2e.sh

e2e.ads.tls: $(BINDIR)/e2e $(BINDIR)/envoy
	env XDS=ads hack/e2e.sh -tls
