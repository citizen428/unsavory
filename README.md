# unsavory: get rid of those stale bookmarks!

`unsavory` checks your Pinboard bookmarks for dead links (HTTP status code 404) and removes them.
Additionally it will also inform you about links which return a status code other than 200 (OK).
This is a Go re-implementation of the [original Ruby version](https://github.com/citizen428/unsavory-legacy).

## Installation

### Users

For now `unsavory` can only be installed via `go install`:

```sh
$ go install github.com/citizen428/unsavory...
```

### Developers

```sh
go get -u github.com/citizen428/unsavory...
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

## Warning

Any link that returns an HTTP status code of 404 will be deleted without warning. There's no undo,
use at your own risk!

## Todo

- [ ] Add tests
- [ ] Feature parity with [legacy Ruby version](https://github.com/citizen428/unsavory-legacy)
- [ ] Add option to replace links with Archive.org links (?)
- [ ] Add option to update redirects (?)

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/citizen428/unsavory.

## License

The program is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
