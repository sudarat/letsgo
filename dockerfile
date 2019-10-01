FROM golang:1.13
WORKDIR /Users/veracharttungwinyoo/Desktop/Prare/godev/src/pp
COPY . .

RUN go get github.com/qiangxue/fasthttp-routing
RUN go get github.com/go-sql-driver/mysql
