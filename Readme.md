# Gin Layout Boilerplate

## Manual Installation
Go `assetsbuild` and install dependencies
```sh
npm install
```

For building css
```
npx tailwindcss -i ./css/input.css -o ../public/assets/css/index.css
```

Back to main directory and run
```
go mod download
```

Generate assets binary
```
go-bindata ./public/assets/...
```

## Makefile Installation
Install dependencies
```sh
make install
```

Building assets
```sh
make build-assets
```

## Run with Makefile
Run with `local` configuration
```sh
make run-local
```

Run with `prod` configuration
```sh
make run-prod
```
