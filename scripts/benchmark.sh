#!/bin/bash

# Enable exit on error
set -e

# Function to print messages with color
print_message() {
  local color="$1"
  local message="$2"
  tput setaf "$color"
  echo -e "$message"
  tput sgr0
}

# Function to run benchmarks in a directory
run_benchmarks() {
  local dir="$1"
  print_message 2 "⏱️  Running benchmarks in $dir"
  (cd "$dir" && go test -bench=. -benchmem)
}

# Find all directories containing Go files and run benchmarks
find . -type d | while read -r dir; do
  if ls "$dir"/*.go &> /dev/null; then
    run_benchmarks "$dir"
  fi
done

print_message 2 "✅ All benchmarks completed."