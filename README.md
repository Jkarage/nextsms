# [Nextsms Golang Client](https://nextsms.co.tz)

Golang Package to easy the integration with nextsms SMS Gateway

## Getting started

- Using this package would require an active nextsms account.

- Add to your environment the username(`NEXT_USERNAME`) and password(`NEXT_PASSWORD`) in your environment.

```bash

export NEXT_USERNAME=<yournextsmsusername> 
export NEXT_PASSWORD=<yournextsmspassword>

```

## Installing the package

To use this in your golang project, add it to the module.

```bash
go get github.com/Jkarage/nextsms
```

### Send SMS

Here is an example on how to send an SMS with this package;

```golang
import (
    "github.com/Jkarage/nextsms"
    "io"
)

func main() {
    client := nextsms.New()
    resp, err := client.SendSDSMS("+255713507067", "Hello world", "")
    if err != nil {
        // handle the error
        log.Fatal(err)
    }

    // handle the response
}
```

### Send SMS to Multiple Phone Numbers

Here is an example on how to send sms to multiple phone numbers, using nextsms client

```golang
import (
    "github.com/Jkarage/nextsms"
    "log"
)

func main() {
    client := nextsms.New()
    resp, err := client.SendMDSMS("Hi There", []string{"2557135070XX", "2557540534XX"}, "")
    if err != nil {
    log.Fatal(err)
    }
    if err != nil {
        // handle the error
        log.Fatal(err)
    }

    // handle the response
}
```

More Features in Progress. Tests on the way as well.