ls -tp --time=atime | grep -v '/$' | awk '{printf "%s,", $0}' | sort