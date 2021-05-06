FROM golang as development
WORKDIR /build
ADD go.mod go.sum ./
RUN export GO111MODULE=on \
    && export GOPROXY=https://goproxy.cn,direct \
    && export CGO_ENABLED=0 \
    && go mod download

COPY . .
RUN go build -o app main.go

FROM golang AS production
ARG APP_PORT
WORKDIR /app
COPY --from=development /build/app .
EXPOSE ${APP_PORT}