#!/bin/bash

# Enable exit on error for non-watch parts of the script
set -e

# Function to print messages with color
print_message() {
  local color="$1"
  local message="$2"
  tput setaf "$color"
  echo "$message"
  tput sgr0
}

# Function to run tests in a directory
run_tests() {
  local dir="$1"
  print_message 8 "☑️  Running tests in $dir"
  (cd "$dir" && go test ./... -v)
}

# Function to get the current checksum of Go files
get_checksum() {
  if command -v md5sum &> /dev/null; then
    find . -name '*.go' -type f -exec md5sum {} \; | md5sum
  else
    find . -name '*.go' -type f -exec md5 {} \; | md5
  fi
}

# Function to watch for changes and run tests
watch_tests() {
  ./scripts/test.sh
  print_message 3 "⏳ Watching for changes..."
  last_checksum=$(get_checksum)

  while true; do
    sleep 2
    current_checksum=$(get_checksum)
    if [ "$last_checksum" != "$current_checksum" ]; then
      clear
      print_message 6 "🏁 Detected changes, running tests..."
      if ! ./scripts/test.sh; then
        print_message 1 "❌ Tests failed, continuing to watch for changes..."
      fi
      last_checksum=$current_checksum
    fi
  done
}

# Check for the -w flag
if [ "$1" == "-w" ]; then
  # Trap SIGINT (CTRL + C) to exit the script gracefully
  trap "print_message 1 '🛑 Stopping watch...'; exit 0" SIGINT
  watch_tests
else
  # Find all directories containing Go files and run tests
  find . -type d | while read -r dir; do
    if ls "$dir"/*.go &> /dev/null; then
      run_tests "$dir"
    fi
  done
  print_message 2 "✅ All tests completed."
fi