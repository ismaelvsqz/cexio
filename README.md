# Unofficial CEX.IO Rest API Go Client.

This library simplifies the CEX.IO Rest API consumption.

# Example

The following example shows two calls. One is private and the other is public.

    package main

    import (
        "fmt"

        "github.com/ismaelvsqz/cexio"
    )

    func main() {
        api := cexio.API{
            Username:  "up00000001",
            APIKey:    "XXXXXXXXXXXXX",
            APISecret: "YYYYYYYYYYYYY"}

        resp, err := api.Do(cexio.Request{
            Method: cexio.MethodActiveOrdersStatus,
            Data:   map[string]interface{}{"orders_list": []string{"111111111", "22222222222"}},
        })

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%s\n", resp)

        resp, err = api.Do(cexio.Request{
            Method:     cexio.MethodTicker,
            Parameters: []string{"BTC", "USD"},
        })

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%s\n", resp)
    }