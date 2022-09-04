DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
LOCAL_BIN:=$(DIR)/bin

export THIRD_PARTY_PROTO_PATH:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))integrations/proto

include .ra9/make/*.make