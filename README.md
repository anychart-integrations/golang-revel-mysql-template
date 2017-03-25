[<img src="https://cdn.anychart.com/images/logo-transparent-segoe.png?2" width="234px" alt="AnyChart - Robust JavaScript/HTML5 Chart library for any project">](https://anychart.com)
# Go basic template

This example shows how to use Anychart library with the Go programming language, Revel web framework and MySQL database.

## Running

To use this sample you must have Go installed as described [here](https://golang.org/doc/install), and follow the [official Go code organization](https://golang.org/doc/code.html);
MySQL installed and running, if not please check out [MySQL download page](https://dev.mysql.com/downloads/installer/) and follow [these instructions](http://dev.mysql.com/doc/refman/5.7/en/installing.html).

To check your installations, run the following command in the command line:
```
$ go version
go version go1.7.1 linux/amd64 # sample output
$ mysql --version
mysql  Ver 14.14 Distrib 5.5.52, for debian-linux-gnu (x86_64) using readline 6. # sample output
```

To start this example run commands listed below.

Install Revel:
```
$ go get github.com/revel/cmd/revel
```

Clone the repository from github.com:
```
$ go get github.com/anychart-integrations/golang-revel-mysql-template
```

Navigate to the repository folder:
```
$ cd $GOPATH/src/github.com/anychart-integrations/golang-revel-mysql-template
```

Get dependencies:
```
$ go get ./...
```

Set up MySQL database, use -u -p flags to provide username and password:
```
$  mysql < database_backup.sql
```

Run example:
```
$ revel run github.com/anychart-integrations/golang-revel-mysql-template
```

Open browser at http://localhost:9000/

## Workspace
Your workspace should look like:
```
$GOPATH/
    bin/
    pkg/
    src/
        github.com/
            {user}/
                golang-revel-mysql-template/
                    app/
                        controllers/
                            app.go      # main index controller
                            gorp.go     # database controller
                            init.go     # init contoller
                        models/
                            fruit.go    # model for data
                        routes/
                            routes.go   # revel auto-generated file for routes processing
                        tmp/
                            main.go     # revel auto-generated file
                        views/
                            App/
                                index.html  # template for chart render
                            errors
                                404.html
                                500.html
                            debug.html
                            flash.html
                            footer.html
                            header.html
                        init.go         # revel auto-generated init file
                    conf/
                        app.conf        # revel main configuration file
                        routes          # app routes definition
                    messages/
                        sample.en       # revel i18n
                    public/             # public assets
                        css/
                            style.css   # css file
                        fonts/
                        img/
                        js/
                    tests/              # test suites
                    .gitignore
                    database_backup.sql     # MySQL database dump
                    LICENSE
                    README.md

```

## Technologies
Language - [Go](https://golang.org/)<br />
Web framework - [Revel](https://revel.github.io)<br />
ORM - [Gorp](https://github.com/go-gorp/gorp)<br />
Database - [MySQL](https://www.mysql.com/)<br />

## Further Learning
* [Documentation](https://docs.anychart.com)
* [JavaScript API Reference](https://api.anychart.com)
* [Code Playground](https://playground.anychart.com)
* [Technical Support](https://anychart.com/support)

## License
[© AnyChart.com - JavaScript charts](http://www.anychart.com). Released under the [Apache 2.0 License](https://github.com/anychart-integrations/golang-revel-mysql-template/blob/master/LICENSE).
