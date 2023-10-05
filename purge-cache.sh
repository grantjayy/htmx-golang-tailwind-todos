#!/bin/bash

# Purge Everything
curl -X POST "$API_ENDPOINT" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $CLOUDFLARE_API_KEY" \
    --data '{"purge_everything":true}'