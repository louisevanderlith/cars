go mod tidy
GOOS="linux" CGO_ENABLED="0" go build
pub get
webdev build