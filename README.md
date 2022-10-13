# checkredirects

Get all URLs from a file and check HTTP Response status codes.
Redirections are not applied and initial HTTP status is returned.

## Build:
```bash
cd <into the code>
go build
```

## Example Input file 'sites.txt':
```bash
https://www.example.com
https://www.example.net
https://www.example.org
```

## Usage:
```bash
./checkredirects sites.txt
```

Results will be presented in stdout.

