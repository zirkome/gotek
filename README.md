Gotek
======

A Golang library for Epitech Intra Webservice

Installation
======

```sh
go get -u https://github.com/kokaz/gotek
```

Usage
======

Here is a little sample to log into the Epitech Intranet, get the dashboard and get your schedule

```go
client := &gotech.Client{Login: "login_x", Password: "unix_password"}
dashboard, _ := client.SignIn()
schedule, _ := client.GetSchedule()
```

Check out the [sample](https://github.com/kokaz/gotek/blob/master/sample "Sample") folder to get the whole code

Run the sample
======

You can use your own `golang` binary or you can use the Dockerfile provided in the repository

With Docker:

```sh
cd sample
docker build -t my-sample-app .
docker run --it --rm --name my-running-app my-sample-app
```

Manually:

```sh
cd sample
go get -v -d
go get -v
$GOPATH/bin/sample
```
