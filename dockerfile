# 基础镜像
# docker run -itd name=test golang:1.21.5-alpine3.18 ash
FROM golang:1.21.5-alpine3.18
# go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 备份原始安装源
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
# 修改为国内源
# RUN sed -i 's/https:\/\/dl-cdn.alpinelinux.org/http:\/\/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk upgrade
# 安装基础软件
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o translate main.go
RUN apk add translate-shell sqlite-dev bash git build-base wget ffmpeg yt-dlp
RUN wget -e use_proxy=yes -e http_proxy=192.168.1.5:8889 -e https_proxy=192.168.1.5:8889 https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux --no-check-certificate --output-document=/usr/local/bin/yt-dlp
# 准备软件
# --no-check-certificate --output-document=/usr/local/bin/yt-dlp
# --no-check-certificate --directory-prefix=/usr/local/bin
#CMD ["go","run","/root/freeTranslate/main.go"]
#docker build -t trans:v1 .
#docker run -itd --name=trans1 -v /d/srt:/srt -e APPID=20230419001647901 -e KEY=rsNTVeBCtQ1sF7RSmFVq trans:v1 ash
#docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
