# jsonforwiki

写wiki时, 部分web编辑器在粘贴时会过滤掉空格和制表符, 导致json格式没法看. 这个小程序把空格替换成了不可见字符, 保持json格式.

## Use

```bash
cat << EOF | go run main.go
    {"key":"value"}
EOF
```

## Install
```
go get github.com/Jiacheng-z/jsonforwiki

cat << EOF | jsonforwiki
    {"key":"value"}
EOF
```