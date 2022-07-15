# VxHash
Simple IoC crawler tool.

# Build
You can use simply the following command.<br>
```bash
go build -ldflags "-s -w"
```

# Usage
```bash
./VxHash -dtype ipaddr -output ips.txt # Crawl IP address data
./VxHash -dtype hash -output hash.txt  # Crawl hash data
```