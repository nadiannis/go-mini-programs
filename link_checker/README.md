# Link Checker

Check & get information about a URL.

## Usage

### Examples

**Example 1**

```sh
go run main.go check https://nadiannis.xyz
```

Output:

```
GET https://nadiannis.xyz

HTTP/2.0 200 OK
Cache-Control: public, max-age=0, must-revalidate
Content-Disposition: inline
Content-Type: text/html; charset=utf-8
Etag: W/"4788dd2e465d7eb309eea6f6f45ce91c"
Access-Control-Allow-Origin: *
Age: 26657750
Date: Fri, 01 Dec 2023 06:30:02 GMT
Server: Vercel
Strict-Transport-Security: max-age=63072000
```

**Example 2**

```sh
go run main.go check https://myrandomurl.com
```

Output:

```
Get "https://myrandomurl.com": dial tcp: lookup myrandomurl.com: no such host
```
