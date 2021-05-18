package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	perfHist = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "myapp_endpoint_sort_ms",
		Help: "The response time histogram for the sort endpoint",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func track(h prometheus.Histogram) (prometheus.Histogram, time.Time) {
	return h, time.Now()
}

func duration(h prometheus.Histogram, start time.Time) {
	ms := time.Since(start).Nanoseconds()
	h.Observe(float64(ms))
}

// -------------------------------------------------

func swapPlaces(array []int, j int) {
	array[j], array[j+1] = array[j+1], array[j]
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swapPlaces(arr, j)
			}
		}
	}
}

// -------------------------------------------------

func parseIntArray(s string) []int {
	arr := strings.Split(s, ",")
	n := len(arr)
	nums := make([]int, n)
	for i, v := range arr {
		nums[i], _ = strconv.Atoi(v)
	}
	return nums
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	defer duration(track(perfHist))

	v := r.URL.Query().Get("values")
	nums := parseIntArray(v)
	bubbleSort(nums)
	fmt.Fprintf(w, "%v", nums)
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/sort", sortHandler)
	http.ListenAndServe(":2112", nil)
}
