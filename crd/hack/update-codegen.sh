#!/bin/bash

vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/shohagrana64/crd/pkg/client \
github.com/shohagrana64/crd/pkg/apis \
stable.example.com:v1alpha1