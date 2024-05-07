FROM golang:1.22.1-alpine AS BUILDER

WORKDIR /app

COPY . .

RUN go build -o live-editor-api

# main stage
FROM node:20.11.1-alpine

ENV NODE_VERSION 20.11.1
ENV API http://api.live-editor.ikeoha.xyz

# Install necessary packages
RUN apk update && \
    apk add --no-cache \
    curl \
    wget

RUN npm install -g sass \
    && npm install -g typescript

WORKDIR /api

COPY --from=BUILDER /app/live-editor-api .

EXPOSE 8000

CMD ["./live-editor-api"]