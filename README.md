# go-pocket
Pocket (getpocket.com) API client for Go (golang).

## Usage

#### Install

`go get github.com/motemen/go-pocket/...`

#### Use
```
mkdir ~/.pocket/config
echo "MY_POCKET_API_CONSUMER_KEY" > ~/.pocket/config/consumer_key
pocket list
# Visit the URL listed in order to authenticate with Pocket
# After succesful authentication, your Pocket article list will appear
```

