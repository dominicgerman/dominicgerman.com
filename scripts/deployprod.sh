#!/bin/bash

git pull origin main
/home/dgerman/hosts/dominicgerman.com/scripts/buildprod.sh &> /home/dgerman/hosts/buildprod.log
sudo systemctl restart portfolio.service
curl -X POST "https://api.cloudflare.com/client/v4/zones/$CLOUDFLARE_ZONE_ID/purge_cache" \
     -H "Authorization: Bearer $CLOUDFLARE_TOKEN" \
     -H "Content-Type: application/json" \
     --data '{"purge_everything":true}'
