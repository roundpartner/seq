[![Build Status](https://travis-ci.org/roundpartner/seq.svg?branch=master)](https://travis-ci.org/roundpartner/seq)

# SEQ
A Task Queue In Go
## Compiling
```bash
go build
```

## Running
```bash
./seq
```

## Usage

```bash
curl http://0.0.0.0:6060 -X POST -d "\"hello world\""
```

```bash
curl http://0.0.0.0:6060
```

> [{"id":1,"body":"hello world"}]

```bash
curl http://0.0.0.0:6060/1 -X DELETE
```
