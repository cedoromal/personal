FROM golang:1.24.2-bookworm AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./server
COPY view ./view
COPY daisyui.js daisyui-theme.js input.css ./
RUN wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x ./tailwindcss-linux-x64
RUN mv ./tailwindcss-linux-x64 /bin/tailwindcss
RUN tailwindcss -i ./input.css -o ./view/static/styles/main.css --minify 


FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/view ./view
COPY --from=build /app/server ./
EXPOSE 8080
ENTRYPOINT [ "./server" ]