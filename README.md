# Site reachable CLi

Small CLI tool to check if a site is reachable. Written in Go.

## Lessons
* User agent string is critical
* Examples:
  * `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:137.0) Gecko/20100101 Firefox/137.0` -> facebook responds with 400 for this
  * `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11` -> chatgpt sometimes returns 403 for this