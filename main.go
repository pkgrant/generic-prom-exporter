package main

import (
    "flag"
    "log"
    "fmt"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    "static-exporter/client"
    "static-exporter/collector"
)

var addr = flag.String("listen-address", ":9999", "The address to listen for HTTP requests.")

func main() {

    values, err := client.GetThresholds()

    if err != nil {
        panic(err)
    }


    fmt.Printf("Value: %+v\n", *values)

//    targets := values.Targets

//    for _, host := range values.Targets {
//        fmt.Printf("Value: %+v\n", host)
//    }
//    fmt.Printf("%+v\n", targets)

//    hostes := targets[1]

//    fmt.Printf("%+v\n", hostes.Host)
    myExporter := prometheus.NewRegistry()
    coll := collector.New()
    myExporter.MustRegister(coll)
    handler := promhttp.HandlerFor(myExporter, promhttp.HandlerOpts{})

    flag.Parse()
    http.Handle("/metrics", handler)
    log.Fatal(http.ListenAndServe(*addr, nil))
}
