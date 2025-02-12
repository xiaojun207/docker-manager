#!/bin/sh
param=""

if [[ "$driverName" != "" ]]; then
    param="$param -driverName $driverName"
fi

if [[ "$useCache" != "" ]]; then
    param="$param -useCache $useCache"
fi

if [[ "$consoleCode" != "" ]]; then
    param="$param -consoleCode $consoleCode"
fi

if [[ "$dataSourceUrl" != "" ]]; then
    param="$param -dataSourceUrl $dataSourceUrl"
fi

/app/App $param $@
