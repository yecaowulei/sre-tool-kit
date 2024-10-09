# stk 「运维工具」

## 1. 功能介绍

1. 从源mysql库同步表结构及数据到目标mysql库
2. 上传/下载指定云厂商指定bucket的文件列表，支持指定本地路径
3. 修改gitlab nacos-archive仓库yml配置文件内容
4. 修改nacos指定命名空间指定yml配置文件内容
5. 文件加解密
6. 执行数据库脚本
7. 检测数据库脚本规范性

## 2. 工具下载

```bash
# linux
wget -O /usr/local/bin/stk https://github.com/yecaowulei/sre-tool-kit/releases/download/v1.0.0/stk-latest-linux-amd64
chmod +x /usr/local/bin/stk

# windows
curl -o stk.exe https://github.com/yecaowulei/sre-tool-kit/releases/download/v1.0.0/stk-latest-windows64.exe
```

## 3. 帮助信息

```text
Usage:
  stk [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     convert file
  download    download obsfile to local
  help        Help about any command
  transfer    transfer mysql data
  update      update gitlab/nacos file
  version     version information

Flags:
  -h, --help   help for stk
```

## 4. mysql数据同步

```text
Usage:
  stk transfer mysql --source-db-host mysql源库地址 --source-db-port mysql源库端口 --source-db-name mysql源库名字 --source-table-name mysql源库表名 --ignore-table=库名.表名1,库名.表名2 --source-db-user mysql源库用户名 --source-db-password mysql源库用户密码 --dst-db-host mysql目标库地址 --dst-db-port mysql目标库端口 --dst-db-name mysql目标库名字 --dst-db-user mysql目标库用户名 --dst-db-password mysql目标库用户密码 --need-data [true/false]

Examples:
  stk transfer mysql --source-db-host xxx --source-db-port xxx --source-db-name xxx --source-table-name xxx,xxx --ignore-table=xxx,xxx --source-db-user xxx --source-db-password xxx --dst-db-host xxx --dst-db-port xxx --dst-db-name xxx --dst-db-user xxx --dst-db-password xxx --need-data xxx

Flags:
  -h, --help   help for transfer
```

## 5. 上传/下载对象存储的文件

### 文件上传-华为云对象存储
```text
有入参文件指定要上传的本地路径
Usage:
  stk upload obsfile -a ak -s sk -b bucketname -f 文件列表 -c 并发数默认为8

Examples:
  stk upload obsfile -a xxx -s xxx -b bucket-test -f xxx -c xxx

没有入参文件指定要上传的本地路径，入参为本地目录名
Usage:
  stk upload obsfile -a ak -s sk -b bucketname -f 本地目录名 -c 并发数默认为8

Examples:
  stk upload obsfile -a xxx -s xxx -b bucket-test -f xxx -c xxx
Flags:
  -h, --help   help for upload
```

### 文件下载-华为云对象存储
```text
文件全部放在本地一个目录里

Usage:
  stk download obsfile -a ak -s sk -b bucketname -f 文件列表 -c 并发数默认为8

Examples:
  stk download obsfile -a xxx -s xxx -b bucket-test -f xxx -c xxx

指定本地路径
Usage:
  stk download obsfile -a ak -s sk -b bucketname -f 文件列表 -c 并发数默认为8 -l true

Examples:
  stk download obsfile -a xxx -s xxx -b bucket-test -f xxx -c xxx -l true

Flags:
  -h, --help   help for download
```

### 文件下载-阿里云对象存储
```text
指定本地路径
Usage:
  stk download ossfile -a ak -s sk -b bucketname -f 文件列表 -c 并发数默认为8 -l true

Examples:
  stk download ossfile -a xxx -s xxx -b bucket-uat -f xxx -c xxx -l true

Flags:
  -h, --help   help for download
```

## 6. 修改gitlab仓库指定文件内容

```text
Usage:
  stk update update gitlab --gitlab-addr gitlab地址 --gitlab-project-id gitlab仓库ID --gitlab-project-branch gitlab仓库ID分支名 --gitlab-token gitlab认证token --gitlab-filename-list 要修改的gitlab文件列表 --gitlab-commit-username gitlab提交者用户名

Examples:
  stk update gitlab --gitlab-addr http://gitlab.xxx.com/ --gitlab-project-id 201 --gitlab-project-branch master --gitlab-token xxx --gitlab-filename-list test/test.yml,prod/test.yml --gitlab-commit-username nanfei.li

Flags:
  -h, --help   help for update
```

## 7. 修改nacos指定命名空间指定yml文件内容

```text
Usage:
  stk update nacos --nacos-addr nacos地址 --nacos-addr-scheme nacos域名的http协议 --nacos-ns-id nacos命名空间id --nacos-username nacos用户名 --nacos-passwd nacos密码 --nacos-filename-list 要修改的nacos文件列表

Examples:
  stk update nacos --nacos-addr nacos.xxx.com --nacos-addr-scheme https --nacos-ns-id test --nacos-username test --nacos-passwd test --nacos-filename-list test1.yml,test2.yml

Flags:
  -h, --help   help for update
```

## 8. 文件转换

### 文件加解密
```text
Usage:
  stk convert file --key 密钥 --encrypt [true/false] --filename-list 要加解密的文件列表 --output 是否抛屏输出(不抛屏输出会创建一个新的文件"原文件名.new") --file-type [file/string] --content xxx

Examples:
  stk convert file --key xxx --encrypt true --filename-list test1.yml,test2.yml --output true --file-type file
  
  stk convert file --key xxx --encrypt false --filename-list test1.yml,test2.yml --output false --file-type file

  stk convert file --key xxx --encrypt true --file-type string --content abc123
    
Flags:
  -h, --help   help for convert
```

### 文件内容转换为export命令
```text
Usage:
  stk convert yamlfile --filename 要处理的文件名 

Examples:
  stk convert yamlfile --filename common.yml --result_file common_result.yml
 
Flags:
  -h, --help   help for convert
```

## 9. 执行数据库脚本
```text
Usage:
  stk execute mysql --db-host mysql地址 --db-port mysql端口 --db-name mysql库名 --db-user mysql用户名 --db-password mysql用户密码 --sql-filename 要执行的数据库脚本文件 --sql-type sql类型

Examples:
  stk execute mysql --db-host xxx --db-port xxx --db-name xxx --db-user xxx --db-password xxx --sql-filename xxx --sql-type [select/other]

Flags:
  -h, --help   help for execute
```
## 10. 检测数据库脚本规范性
```text
Usage:
  stk check mysql --sql-filename 要检测的数据库脚本文件 

Examples:
  stk check mysql --sql-filename xxx 

Flags:
  -h, --help   help for check
```