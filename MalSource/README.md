# MalSource
Simple IoC crawler tool.

# Build
You can use simply the following command.<br>
```bash
go build -ldflags "-s -w"
```

# Usage
```bash
./Malsource -apikey MALSHARE_API_KEY -dtype hash -output hash.txt      # Crawl hash data
./Malsource -apikey MALSHARE_API_KEY -dtype domain -output domain.txt  # Crawl domain data
```