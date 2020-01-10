FROM golang AS go-build
ENV GO111MODULE=on
ADD . /work
WORKDIR /work
RUN CGO_ENABLED=0 go build -o todo-go main.go

FROM busybox
COPY --from=go-build /work/todo-go /usr/local/bin/todo-go
EXPOSE 9000
CMD ["/usr/local/bin/todo-go"]