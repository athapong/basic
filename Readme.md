Go resources
- https://pkg.go.dev/
- https://go.dev/ref/spec
- https://awesome-go.com/
- https://github.com/pallat/uber-go-style-guide-th
- https://github.com/golang/vscode-go/blob/master/docs/settings.md
- https://brew.sh/

Install go on macos
- https://go.dev/dl/ or manual steps below

```bash
curl -OL https://go.dev/dl/go1.18.3.darwin-amd64.tar.gz
tar xfz go1.18.3.darwin-amd64.tar.gz
sudo mv go /usr/local
sudo ln /usr/local/go/bin/go /usr/local/bin/go
```
VS Code extensions
- Go https://marketplace.visualstudio.com/items?itemName=golang.Go
- Error Lens https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens

VS Code configuration
```json
    "go.coverOnSave": true,
    "go.coverOnSingleTest": true,
    "go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred",
    },
    "go.buildFlags": [
        "-tags=integration"
    ]
```    