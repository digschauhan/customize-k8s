FROM golang:alpine as gobuilder

COPY . /
WORKDIR /
#RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o customize-k8s-mac-amd .
#RUN CGO_ENABLED=0 GOOS=darvin GOARCH=386 go build -a -o customize-k8s-mac-386 .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o customize-k8s-linux-amd .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -o customize-k8s-linux-386 .



FROM golang:alpine
WORKDIR /
COPY --from=gobuilder /customize-k8s-linux-amd /


ENTRYPOINT ["/customize-k8s-linux-amd"]

# docker build -t customize-k8s:0.1.0 . && docker tag customize-k8s:0.1.0 prepdigi/customize-k8s:0.1.0 && docker push prepdigi/customize-k8s:0.1.0
# 
