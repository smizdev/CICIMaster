#
# Setup our GOPATH to project directory
#
$env:GOPATH=$PSScriptRoot

#
# Disable cgo for best cross-platform support
#
$env:CGO_ENABLED=0

#
# Run cici
#
go run $PSScriptRoot\src\cici\cmd\cici.go
