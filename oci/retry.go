package oci

import (
	"math"
	"time"
)

func RetryWithBackoff(retryCount int, operation func() error) error {
    var err error
    for i := 0; i < retryCount; i++ {
        err = operation()
        if err == nil {
            return nil
        }

        // Calcula o tempo de espera para retentativa
        wait := time.Duration(math.Pow(2, float64(i))) * time.Second
        time.Sleep(wait)
    }
    return err
}