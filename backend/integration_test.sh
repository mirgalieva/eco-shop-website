#!/bin/bash

function check_http_code(){
    if [[ $1 -ne $2 ]]; then
        echo FAIL
        exit 1
    else
        echo OK
    fi
}

API_ADDR="http://127.0.0.1:8080"

echo "1. test whether server is accessible"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null $API_ADDR)
check_http_code code 200
echo ""

echo "2. cleaning database test.db before tests"
sqlite3 -line test.db "DELETE FROM users; DELETE FROM products;"
echo OK
echo ""

echo "3. trying to register an user"


echo "3.3. trying to register a user - should be CORRECT"
code=$(curl --write-out "%{http_code}\n" -v -o /dev/null -X POST $API_ADDR/users/register -d '{"Email":"test@mail.com", "Password":"weoufnksbef", "FirstName":"Test", "LastName":"TestTest"}')
check_http_code code 201
echo ""

echo "3.4. trying to register a DUPLICATE user"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null -X POST $API_ADDR/users/register -d '{"Email":"test@mail.com", "Password":"weoufnksbef", "FirstName":"Test", "LastName":"TestTest"}')
check_http_code code 400
echo ""

echo "4. trying to get list of users"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null $API_ADDR/users)
user_id=$(curl -s $API_ADDR/users | jq '.[0] | .ID')
check_http_code code 200
echo ""

echo "4.1. trying to login"
authz_token=$(curl -s -w "%header{Authorization}\n" -o /dev/null $API_ADDR/user/login -d '{"Email":"test@mail.com", "Password":"weoufnksbef"}')
echo ""


echo "5. trying to get user with id=${user_id}"
code=$( curl --write-out "%{http_code}\n" -s -o /dev/null $API_ADDR/users/${user_id} -H "Authorization: ${authz_token}" )
check_http_code code 200
echo ""

echo "6. trying to get unexisting user with id=$(($user_id-1))"
code=$( curl --write-out "%{http_code}\n" -v -o /dev/null $API_ADDR/users/$((user_id-1)) -H "Authorization: ${authz_token}" )
check_http_code code 403
echo ""



echo "9. trying to get user with id=${user_id}"
curl -s $API_ADDR/users/${user_id} -H "Authorization: ${authz_token}"
echo ""
