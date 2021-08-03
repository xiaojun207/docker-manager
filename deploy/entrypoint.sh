#!/bin/sh
param=""

if [[ "$driverName" != "" ]]; then
    param="$param -driverName $driverName"
fi

if [[ "$dataSourceUrl" != "" ]]; then
    param="$param -dataSourceUrl $dataSourceUrl"
fi

/app/App $param $@
