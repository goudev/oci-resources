package main

import (
	"encoding/json"
	"fmt"
	"log"
	"oci-sdk-go/pkg/oci"
)

func main() {
	result, err := oci.ListAllAnalyticsInstances()
	if err != nil {
		log.Fatalf("Erro ao obter instâncias: %v", err)
	}

	// Convertendo o resultado para JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Erro ao converter dados para JSON: %v", err)
	}

	// Imprimindo o JSON resultante
	fmt.Println(string(jsonData))
}

