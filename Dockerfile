FROM golang:latest as build 
 
ENV USER=appuser 
ENV UID=10001  
RUN adduser \     
    --disabled-password \     
    --gecos "" \     
    --home "/nonexistent" \     
    --shell "/sbin/nologin" \     
    --no-create-home \     
    --uid "${UID}" \     
    "${USER}" 
 
 
WORKDIR /build 
 
COPY go.mod ./ 
COPY go.sum ./ 
 
RUN go mod download 
 
COPY . . 
 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /build/app ./cmd/app/main.go 
 
 
FROM alpine:latest
  
COPY --from=build /etc/passwd /etc/passwd 
COPY --from=build /etc/group /etc/group 
 
COPY --from=build /build/config /config 
COPY --from=build /build/app /app 
 
USER appuser:appuser 

CMD [ "/app" ]