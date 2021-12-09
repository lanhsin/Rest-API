#!/bin/bash

error=0

echo "###########################"
echo "# Launching services ..."
echo "###########################"
/usr/bin/redis-server&
/usr/sbin/nginx&
/usr/bin/pemgr/pemgr-plat&
/usr/bin/pemgr/pemgr-server&
sleep 5

echo ""
echo "##########################"
echo "# Waiting for 10 seconds ..."
echo "##########################"
sleep 10

echo ""
echo "##########################"
echo "# Checking processes ..."
echo "##########################"
pgrep -x redis-server &> /dev/null
if [ $? -eq 0 ]; then
    echo redis-server: OK
else
    echo redis-server: FAIL
    error=1
fi

pgrep -x nginx &> /dev/null
if [ $? -eq 0 ]; then
    echo nginx: OK
else
    echo nginx: FAIL
    error=1
fi

pgrep -x pemgr-plat &> /dev/null
if [ $? -eq 0 ]; then
    echo pemgr-plat: OK
else
    echo pemgr-plat: FAIL
    error=1
fi

pgrep -x pemgr-server &> /dev/null
if [ $? -eq 0 ]; then
    echo pemgr-server: OK
else
    echo pemgr-server: FAIL
    error=1
fi

echo ""
echo "##########################"
echo "# Checking ports ..."
echo "##########################"
nc -z localhost 6379 &> /dev/null
if [ $? -eq 0 ]; then
    echo redis-server: OK
else
    echo redis-server: FAIL
    error=1
fi

nc -z localhost 80 &> /dev/null
if [ $? -eq 0 ]; then
    echo nginx: OK
else
    echo nginx: FAIL
    error=1
fi

nc -z localhost 50010 &> /dev/null
if [ $? -eq 0 ]; then
    echo pemgr-plat: OK
else
    echo pemgr-plat: FAIL
    error=1
fi

nc -z localhost 8080 &> /dev/null
if [ $? -eq 0 ]; then
    echo pemgr-server: OK
else
    echo pemgr-server: FAIL
    error=1
fi

exit $error