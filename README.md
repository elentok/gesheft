Gesheft
=======

An SSH tunnel manager in GO
(based on [Shaft](https://github.com/n0nick/shaft) by [n0nick](https://github.com/n0nick))

Install
-------

```
go get github.com/elentok/gesheft
```


Usage
-----

```

gesheft list    - lists all of the tunnels
gesheft active  - lists the active tunnels (removes zombie tunnels)
gesheft info    - shows information about a tunnel

gesheft start   <tunnel_name>
gesheft stop    <tunnel_name>
gesheft restart <tunnel_name>

```

Gehseft uses the same config file as shaft (see the "example-config.yml" file)

Development
-----------

To run the tests, you'll need [ginkgo](https://github.com/onsi/ginkgo)

    go get github.com/onsi/ginkgo/ginkgo
    go get github.com/onsi/gomega
    ginkgo -r
