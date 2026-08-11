<p align="center">
  <img src="./assets/spider.svg" alt="Spider" width="240"/>
</p>

# LinkSpider

A CLI tool written in Go that crawls a web page, extracts all links, and checks them for broken/dead status (404, timeout, DNS errors).


## How it works

1. **Fetch** — downloads the HTML of a given page (`internal/fetcher`)
2. **Parse** — extracts all `<a href="...">` links from the page (`internal/parser`)
3. **Resolve** — turns relative links (`/about`) into absolute URLs using the page as base
4. **Check** — sends a request to each link and records its status (`internal/checker`)
5. **Report** — prints each link with its result: alive, broken, or unreachable

## Usage

```bash
go run main.go -url=https://example.com
```

### Example output

```
Link - https://example.com - STATUS: 200 [OK]
Link - https://example.com/this-page-does-not-exist - STATUS: 404 [BROKEN]
Link - https://this-domain-does-not-exist-abc123xyz.com - STATUS: 0 [FAIL]: dial tcp: lookup ...: no such host
```

## Project structure

```
linkspider/
├── go.mod
├── main.go                    # wires fetcher → parser → checker together
└── internal/
    ├── fetcher/
    │   └── fetcher.go         # downloads HTML by URL
    ├── parser/
    │   ├── parser.go          # extracts links from HTML
    │   └── parser_test.go     # table-driven tests
    └── checker/
        └── checker.go         # checks a single link's status
```

## Status

- ✅ Working end-to-end: fetch → parse → resolve → check → report (sequential)
- ✅ Basic test coverage for the parser
- ⬜ No concurrency yet — links are checked one at a time
- ⬜ Some servers reject HEAD requests, producing a false "unreachable" result instead of the real status code

## Roadmap

### Next up
- [ ] **HEAD → GET fallback** in the checker: some servers don't support `HEAD` requests and drop the connection, which currently gets misreported as a network failure instead of a real status code
- [ ] **Concurrency**: check links in parallel using goroutines + a worker pool (with a configurable concurrency limit), instead of one at a time
- [ ] **Request timeouts** via `context.Context`, so a single slow/hanging link can't stall the whole run

### Later
- [ ] Table-driven tests for `fetcher` and `checker` (using `httptest`)
- [ ] Summary at the end of the run (total links, broken count, unreachable count)
- [ ] Recursive crawling: follow internal links to check an entire site, not just one page
- [ ] Wrap the tool as a small REST API (`POST /check {"url": "..."}` → JSON report)

