#!/bin/bash
find . -name "*_vgen.go" -exec rm {} +
rm main_wasm.go

# probably it would be nice to add this to vugugen ("vugugen -c" for example)