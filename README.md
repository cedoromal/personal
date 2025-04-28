## Deployment
This project uses [Docker](https://docs.docker.com/) for deployment. To build the image, use the following command:
```bash
docker build -t cedoromal/personal .
```
To run the image, use the following command:
```bash
docker run -p 8080:8080 cedoromal/personal
```
By default, Gin is in debug mode. If you wish to set it to release mode, simply add `-e GIN_MODE=release` to the `docker run` command. For example:
```bash
docker run -p 8080:8080 -e GIN_MODE=release cedoromal/personal
```
You may also change port 8080 to whatever port you want to use.

## Development
This project uses [air](https://github.com/air-verse/air) for live reloads, and  [Standalone Tailwind CLI](https://tailwindcss.com/blog/standalone-cli) for styling.

### air
To automatically re-compile on change, install air using your preferred installation method and run the following command on your terminal:
```bash
air
```

### Tailwind CLI
To watch for changes and automatically update `main.css`, download the Standalone Tailwind CLI (for Linux) using the `./download_tailwindcss.sh` script
```bash
./download_tailwindcss.sh
```
and use the downloaded Tailwind CLI with the following command:
```bash
./tailwindcss -i ./input.css -o ./view/static/styles/main.css --watch
```