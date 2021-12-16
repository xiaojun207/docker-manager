function push_tag()
{
    #拿出当前匹配的v最近版本号
    latelyTag=$(git describe --match "v*" --abbrev=0 --tags $(git rev-list --tags --max-count=1))

    #版本前缀
    tagPre="v"

    #截取字符串,tag 是变量名，# 号是运算符，*v 表示从左边开始删除第一个 v 号及左边的所有字符
    version=${latelyTag#*$tagPre}
    echo "最近版本："$latelyTag
    #echo $version

    #按.分割字符串，变成3段
    array=(${version//./ })
    arrayLen=${#array[*]}
    if [[ $arrayLen != 3 ]];then
      echo "版本号长度不对，必须为 v1.x.x的格式"
      read -p "按任意键关闭" -n 1
      exit
    fi

    #版本号处理
    newVersion=""
    for (( i = 0; i < $arrayLen; i++ )); do
        v=${array[i]}
        if [[ $i == 2 ]];then
          #第三位小版本+1
          v=`expr $v + 1`
        fi
        newVersion+=$v"."
    done
    newVersion=${newVersion%?}

    #拼接
    newTag=${tagPre}${newVersion}
    echo "新的版本："$newTag

    if git tag -l | grep -q $newTag;then
      echo $newTag"已存在"
      read -p "按任意键关闭" -n 1
      exit
    fi

    #打标签
    #git tag -a $newTag -m ""
    #推送单个标签到远端
    #git push origin $newTag
}

push_tag
read -p "按任意键关闭" -n 1
