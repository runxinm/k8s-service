package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "strconv"
    "strings"
    "sync"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
    namespace = "ping"
    address   = ":18080"
    interval  = time.Second * 10
)

var (
    pingLatency = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Namespace: namespace,
            Name:      "latency_ms",
            Help:      "Ping latency in milliseconds",
        },
        []string{"destination"},
    )

    pingCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Namespace: namespace,
            Name:      "count",
            Help:      "Ping count",
        },
        []string{"destination"},
    )

    pingErrors = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Namespace: namespace,
            Name:      "errors",
            Help:      "Ping errors",
        },
        []string{"destination"},
    )

    pingTargets = []string{"8.8.8.8", "10.160.100.101", "10.160.10.1"}
)

func main() {
    prometheus.MustRegister(pingLatency)
    prometheus.MustRegister(pingCount)
    prometheus.MustRegister(pingErrors)
    fmt.Println("Starting exporter on", address)

    go func() {
        for {
            var wg sync.WaitGroup
            for _, target := range pingTargets {
                wg.Add(1)
                go func(target string) {
                    defer wg.Done()
                    output, err := exec.Command("ping", "-c", "1", target).CombinedOutput()
                    if err != nil {
                        pingErrors.WithLabelValues(target).Inc()
                        fmt.Printf("Error pinging %s: %s\n", target, err)
                        return
                    }
                    // Parse ping output for latency
                    outputStr := string(output)
                    latencyIndex := strings.Index(outputStr, "time=")
                    if latencyIndex == -1 {
                        pingErrors.WithLabelValues(target).Inc()
                        fmt.Printf("Error parsing ping output for %s\n", target)
                        return
                    }
                    latencyStr := outputStr[latencyIndex+5:]
                    latencyEndIndex := strings.Index(latencyStr, " ")
                    latencyMs, err := strconv.ParseFloat(latencyStr[:latencyEndIndex], 64)
                    if err != nil {
                        pingErrors.WithLabelValues(target).Inc()
                        fmt.Printf("Error parsing latency for %s: %s\n", target, err)
                        return
                    }
                    pingLatency.WithLabelValues(target).Set(latencyMs)
                    pingCount.WithLabelValues(target).Inc()
                }(target)
            }
            wg.Wait()
            time.Sleep(interval)
        }
    }()

    http.Handle("/tmetrics", promhttp.Handler())
    http.ListenAndServe(address, nil)
    fmt.Println("Exiting")

}
