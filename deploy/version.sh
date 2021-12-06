
version="$1"
path="./web/handler/user/UserHandler.go"
#	Version = "1.0.8"
eval $(sed -ip 's@	Version = ".*"@	Version = "'$(echo $version)'"@g' $path)
rm "${path}p"
