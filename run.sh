#!/bin/sh

set -e

# Check usage
if [ $# -ne 0 ]; then
	echo "usage: $0" >&2
	exit 1
fi

# Require secure entry of password
read -sp "Enter the API password for user spectator: " api_pass
echo ""

# Run the docker daemon
IP="$(ifconfig | grep inet | grep -v inet6 | grep -v 127.0.0.1 | awk 'NR==1{print $2}')"
echo "Starting proxy on ${IP}:8080"
docker run -d \
	-p 80:8080 \
	-e PORT=8080 \
	-e ARC_API_URL="https://api.spectator.arcpublishing.com" \
	-e ARC_API_USERNAME="spectator" \
	-e ARC_API_PASSWORD="$api_pass" \
	spectech/arc-api-proxy