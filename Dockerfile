FROM node:latest as html
COPY html /root/html
WORKDIR /root/html
RUN npm i && npm run build

FROM golang:1.19 as server
WORKDIR /root
COPY main.go go.sum go.mod ./.
COPY ./parser ./parser
RUN mkdir -p /root/html/build
COPY --from=html /root/html/build /root/html/build
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o music

FROM alpine:latest
COPY --from=server /root/music /root/music
WORKDIR /root
EXPOSE 8080
CMD ["./music"]
