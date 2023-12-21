#
# this curl POST will create a new DNS resource record (RR) in zone the wit.com
# In this case it will map www3.wit.com to a IPv6 address
# replace the auth key (e088...) and zone ID (27b9...) with the ones from your cloudflare account
#
curl --request POST \
  --url https://api.cloudflare.com/client/v4/zones/27llxxPutYourZoneIDherexxx497f90/dns_records \
  --header 'Content-Type: application/json' \
  --header 'X-Auth-Key: e08806adxxxPutYourAPIKeyHerexxxxa7d417a7x' \
  --header 'X-Auth-Email: test@wit.com' \
  --data '{
  "name": "www3",
  "type": "AAAA"
  "content": "2001:4860:4860::5555",
  "ttl": 3600,
  "proxied": false,
  "comment": "WIT DNS Control Panel",
}'

# This will verify an API token
curl -X GET "https://api.cloudflare.com/client/v4/user/tokens/verify" \
     -H "Authorization: Bearer AAAPutYourTokenInHereSoYouCanTestItL5Cl3" \
     -H "Content-Type:application/json"
