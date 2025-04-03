ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "owlhook"
DOCKER_NAME = "owlhook"

include ./hack/hack-cli.mk
include ./hack/hack.mk


fmt:
	command -v gofumpt || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go install mvdan.cc/gofumpt@latest && cd $(WORK))
	gofumpt -w -d .