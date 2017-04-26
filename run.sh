#!/bin/bash

set -e

# Check usage
if [ $# -ne 0 ]; then
	echo "usage: $0" >&2
	exit 1
fi

# Export necessary environment variables
if [ -z "$ARC_API_USERNAME" ]; then
	export ARC_API_USERNAME="spectator"
fi

if [ -z "$ARC_API_URL" ]; then
	export ARC_API_URL="https://api.spectator.arcpublishing.com"
fi
# Require secure entry of password
if [ -z "$ARC_API_PASSWORD" ]; then
	read -sp "Enter the API password for user spectator: " ARC_API_PASSWORD
	export ARC_API_PASSWORD
	echo ""
fi

# Run server
docker-compose up
