# CEP-Provider (GO)
GoLang service to find Brazilian postal codes (CEP)

## 1. Install:
  - `$ apt-get install golang`

Check if golang was succefully instaled
  - `$ go version`

## 2. Configure:
- Create folder at a suitable location. Usually at the user home directory
  - `$ mkdir go`
  - `$ mkdir go/src`
  - `$ mkdir go/bin`
  - `$ mkdir go/pkg`

- Execute command to download go dependecies
  - `$ cd go`
  - `$ go get`

- Export env variables:
  - `$ export GOPATH="/home/user/go"`
  - `$ export GOBIN="/home/user/go/bin"`

- Check if variables were set succefully
  - `$ go env`

## 3. Clone the project:
  - `$ cd go/src`
  - `$ git clone https://github.com/marsal-silveira/Go-CepProvider.git`

## 4. Build and Run the project:
  - `$ go get ./`
  - `$ go build`
  - `$ go run app.go`