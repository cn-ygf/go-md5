export GOOS=darwin
export GOARCH=amd64
go build -o go-md5-darwin-amd64

export GOOS=linux
export GOARCH=amd64
go build -o go-md5-linux-amd64

export GOOS=linux
export GOARCH=386
go build -o go-md5-linux-386

export GOOS=linux
export GOARCH=arm
go build -o go-md5-linux-arm

export GOOS=windows
export GOARCH=386
go build -o go-md5-windows-386.exe

export GOOS=windows
export GOARCH=amd64
go build -o go-md5-windows-amd64.exe
