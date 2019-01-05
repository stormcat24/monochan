.PHONY: clean
clean:
	bazel clean

.PHONY: dep
dep:
	go mod download
	bazel run //:gazelle -- update-repos -from_file=./go.mod

.PHONY: build
build:
	bazel build -k -- //cmd/... //pkg/...