
# Gofiber and NextJS Static HTML

Gofiber with NextJS Static HTML is a small Go program to showcase for bundling a static HTML export of a Next.js app.

## Requirements

- Go 1.17
- Yarn

## Installing

Clone or download the repository:

```sh
$ git clone https://github.com/openmymai/gofiber_nextjsstatic.git
```

## Usage

From the repository root directory, generate the static HTML export of the Next.js
app (dist folder in nextjs), and build the Go binary:

```sh
$ cd nextjs
$ yarn install
$ yarn run export
$ cd ..
$ go mod init
$ go mod tidy
$ go build main.go
```

Then run the binary:

```sh
$ ./portable
```

## License

[MIT](/LICENSE)

---

© 2022 Sirisak Chantanate
