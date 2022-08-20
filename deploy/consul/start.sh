#!/bin/sh

# brew install consul

nohup consul agent -server -bind 127.0.0.1 -ui -bootstrap-expect 1 -data-dir ~/.consul  > /dev/null &

echo "visit manager ui: http://localhost:8500/ui"