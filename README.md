# chuckbot
A silly IRC bot programmed in Golang. It posts Chuck Norris jokes.

## Usage
Get it.

```
$  go get github.com/galaktor/chuckbot
```

Navigate to the source code.

```
$  cd $GOPATH/src/github.com/galaktor/chuckbot
```

Build it.

```
chuckbot$  go build
```

Run it. Right now it's hardcoded to use SSL. Change that if you want to. Also, escape the `#` character on the terminal if you must.

```
chuckbot$  ./chuckbot -server cool.irc.server:6697 -channel \#furries
```


# Kudos
Uses the cool little IRC client library [thoj/go-ircevent](https://github.com/thoj/go-ircevent).
