ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "owlhook"
DOCKER_NAME = "owlhook"

include ./hack/hack-cli.mk
include ./hack/hack.mk