#!bin/bash

echo "Simple API Client"
read -p "insert endpoint: " endpoint
read -p "insert HTTP Method: " method
read -p "insert request body: " body
read -p "insert request body type: " type

if [ $method = "GET" ]; then
    curl -X $method -Ls $endpoint
elif [ $method = "POST" ]; then
    curl -X $method -H "Content-Type: $type" -d "$body" $endpoint
elif [ $method = "PUT" ]; then
    curl -X $method -H "Content-Type: $type" -d "$body" $endpoint
elif [ $method = "DELETE" ]; then
    curl -X $method $endpoint
else
    echo "HTTP Method harus benar"
fi
