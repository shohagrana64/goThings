CONTROLLER_GEN="./bin/controller-gen"

.PHONY: manifests
# Generate manifests for CRDs
manifests:
	$(CONTROLLER_GEN) crd:crdVersions="v1" paths="./..." output:crd:artifacts:config=config/crd/bases