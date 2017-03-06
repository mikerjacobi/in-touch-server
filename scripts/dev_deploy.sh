./scripts/dev.sh
find . | egrep go | xargs fswatch -m poll_monitor -r -o | xargs -n1 scripts/dev.sh

