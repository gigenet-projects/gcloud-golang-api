# gcloud-golang-api
An API library for GigeNET Cloud written in Golang

## Installation

```
  go get github.com/gigenet-projects/gcloud-golang-api/api
```

## Usage
```
package main

import (
        "bytes"
        "fmt"
        gcloud "github.com/gigenet-projects/gcloud-golang-api/api"
)

var APIKEY  = "YOURAPIKEY"
var SECRET  = "YOURSECRETKEY"
var BASEURL = "https://api.thegcloud.com"

func main() {
        // one client can be used for many commands
        client := gcloud.NewClient(BASEURL, APIKEY, SECRET)

        reservations, err := client.DescribeInstances("")

        if err != nil {
                panic(err)
        }

        var buffer bytes.Buffer
        buffer.WriteString("GCloud Servers\n")
        for _, re := range reservations {
                for _, in := range re.Instances {
                        buffer.WriteString(" InstanceId: ")
                        buffer.WriteString(in.Id)
                        buffer.WriteString(" PublicIpAddress: ")
                        buffer.WriteString(in.PublicIpAddress)
                        buffer.WriteString("\n")
                }
        }
        fmt.Println(buffer.String())

}
```
