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
