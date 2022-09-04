DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
LOCAL_BIN:=$(DIR)/bin

include .ra9/make/*.make