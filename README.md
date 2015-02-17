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
client := &gotek.Client{Login: "login_x", Password: "unix_password"}
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
docker run -it --rm --name my-running-app my-sample-app
```

Manually:

```sh
cd sample
go get -v -d
go get -v
$GOPATH/bin/sample
```

LICENSE
======

```
The MIT License (MIT)

Copyright (c) 2014 Guillaume Fillon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
