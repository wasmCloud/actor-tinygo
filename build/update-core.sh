#!/bin/sh

# Update core.go and model.go from latest smithy definitions

# Usage:  (run from top-level 'actor-tinygo' folder)
#   build/update-core.sh

WASH=wash
$WASH gen -c build/codegen.toml

