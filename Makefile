ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "gweb"
DOCKER_NAME = "gweb"

include ./hack/hack.mk
