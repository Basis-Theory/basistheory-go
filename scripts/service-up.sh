#!/bin/bash

RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"

i=0
timeout=60000

# continue until $n equals 5
while [ $i -le $timeout ]
do
	status=$(curl -s -X PUT -o /dev/null -i -w "%{http_code}" http://localhost:1080/mockserver/status)

    if [ $status == "200" ]
    then
        echo -e "${GREEN}✓${NOCOLOR} API is ready"
        exit 0
    else
        echo -e "❌ API is not ready"

        sleep 3
	    i=$(( i+3000 ))	 # increments $n
    fi
done

curl -X PUT http://localhost:/1080/mockserver/openapi -d '{ "specUrlOrPayload": "https://api-dev.basistheory.com/swagger/v1/swagger.json" }'
echo 'Health check did not pass within timeout'
docker ps -a
#docker logs vault-api
exit 1
