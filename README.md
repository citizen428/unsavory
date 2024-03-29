# unsavory: get rid of those stale bookmarks!

`unsavory` checks your [Pinboard](https://pinboard.in) bookmarks for dead links (specifically HTTP status codes [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) and [410](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410)) and removes
them. Additionally it will also inform you about links which return a status code other than 200 (OK).

This is a Go re-implementation of the [original Ruby version](https://github.com/citizen428/unsavory-legacy).

If you wonder about the name, this program was originally written for [del.icio.us](https://en.wikipedia.org/wiki/Delicious_(website)) and I can't resist a bad pun.

## Installation

### Homebrew

```sh
$ brew tap citizen428/homebrew-tap
$ brew install unsavory
```

### Release

Head over to the [release page](https://github.com/citizen428/unsavory/releases) and download the archive for your operating system/architecture.

### Docker

```sh
$ docker run --rm citizen428/unsavory -token=user:NNNNNN --dry-run
```

### Manual

```sh
$ go install github.com/citizen428/unsavory...
```

## Options

```
Usage of unsavory:
  -dry-run
    	Enables dry run mode
  -proxy-url string
    	HTTP proxy URL
  -token string
    	Pinboard API token
```

## Usage

Just start `unsavory` from the command line and provide the Pinboard API token from your [settings page](https://pinboard.in/settings/password).

```sh
$ unsavory -token=user:NNNNNN
Retrieving URLs
Retrieved 783 URLS
...
```

If you don't want to actually delete links, add the `-dry-run` option:

```sh
$ unsavory -token=user:NNNNNN -dry-run
You are using dry run mode. No links will be deleted!

Retrieving URLs
Retrieved 783 URLS
...
```

### Proxy servers

If a `HTTP_PROXY` environment variable is present, Go's HTTP client will automatically use it.
Alternatively a proxy server can be specified via the `--proxy-url` option:

```sh
$ unsavory --token=user:NNNNNN --proxy-url=http://example.com:8080
```

## Warning

Any link that returns an HTTP status code of 404 or 410 will be deleted without warning. There's no undo, use at your own risk!

## Todo

- [ ] Add tests
- [ ] Feature parity with [legacy Ruby version](https://github.com/citizen428/unsavory-legacy)
- [ ] Add option to replace links with Archive.org links (?)
- [ ] Add option to update redirects (?)

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/citizen428/unsavory.

## License

The program is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
