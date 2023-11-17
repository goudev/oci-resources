package oci

import (
	"context"
	"log"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/usageapi"
)

func GetCostByLastDays(dias int) float64 {
	if(dias>90){
		log.Fatal("Desculpe, mas a quantidade máxima de dias é 90. Informe uma quantidade de dias inferior")
	}
	configProvider := common.CustomProfileConfigProvider("", "")
    tenancyID, _ := configProvider.TenancyOCID()
    client, err := usageapi.NewUsageapiClientWithConfigurationProvider(configProvider)
    helpers.FatalIfError(err)

	dataFinal := GetZeroTime(time.Now().UTC())
	dataInicial := GetZeroTime(dataFinal.AddDate(0, 0, ((dias - dias) - dias)))

	req := usageapi.RequestSummarizedUsagesRequest{
		RequestSummarizedUsagesDetails: usageapi.RequestSummarizedUsagesDetails{
			Granularity:       "DAILY",
			TenantId:          common.String(*common.String(tenancyID)),
			TimeUsageStarted:  &common.SDKTime{dataInicial},
			TimeUsageEnded:    &common.SDKTime{dataFinal},
			CompartmentDepth:  common.Float32(2),
			IsAggregateByTime: common.Bool(false),
			QueryType:         "COST",
		},
	}

	response, err := client.RequestSummarizedUsages(context.Background(), req)
	helpers.FatalIfError(err)

	total := 0.0
	for _, item := range response.Items {
		if item.ComputedAmount != nil {
			total += float64(*item.ComputedAmount)
		}
	}

	return total
}

func ParseAndZeroTime(dateStr string) (time.Time, error) {
    layout := "2006/01/02" // Formato de data em yyyy/mm/dd
    parsedTime, err := time.Parse(layout, dateStr)
    if err != nil {
        return time.Time{}, err
    }
    return time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 0, 0, 0, 0, parsedTime.Location()), nil
}

func GetCostByLastMonths(meses int) float64 {
    // Verifique se o número de meses não excede o limite máximo permitido
    if meses > 12 {
        log.Fatal("Desculpe, mas a quantidade máxima de meses é 12. Informe uma quantidade de meses inferior")
    }

    // Obtenha o provedor de configuração, cliente da API e ID da tenancy como antes
    configProvider := common.CustomProfileConfigProvider("", "")
    tenancyID, _ := configProvider.TenancyOCID()
    client, err := usageapi.NewUsageapiClientWithConfigurationProvider(configProvider)
    helpers.FatalIfError(err)

    // Calcule as datas de início e fim
    dataFinal := GetZeroTime(time.Now().UTC())
    dataInicial := GetZeroTime(dataFinal.AddDate(0, -meses, 0))

    // Crie a requisição com as datas ajustadas
    req := usageapi.RequestSummarizedUsagesRequest{
        RequestSummarizedUsagesDetails: usageapi.RequestSummarizedUsagesDetails{
            Granularity:       "MONTHLY",
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{dataInicial},
            TimeUsageEnded:    &common.SDKTime{dataFinal},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
        },
    }

    // Execute a requisição e processe a resposta como antes
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    helpers.FatalIfError(err)

    // Calcule o total de custos
    total := 0.0
    for _, item := range response.Items {
        if item.ComputedAmount != nil {
            total += float64(*item.ComputedAmount)
        }
    }

    return total
}




// Função para obter o custo entre datas específicas
func GetCostByDate(startDateStr, endDateStr string) float64 {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data inicial:", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data final:", err)
    }

	 // Verifica se o intervalo de datas é maior que 90 dias
	 if endDate.Sub(startDate).Hours() > 365*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 365 dias.")
    }

    // Ajusta a data final para o início do dia seguinte
    endDate = endDate.AddDate(0, 0, 1)

    // Configuração do cliente OCI como anteriormente
    configProvider := common.CustomProfileConfigProvider("", "")
    tenancyID, _ := configProvider.TenancyOCID()
    client, err := usageapi.NewUsageapiClientWithConfigurationProvider(configProvider)
    if err != nil {
        log.Fatal("Erro ao criar o cliente da API de Uso:", err)
    }

    // Cria a requisição com as datas ajustadas
    req := usageapi.RequestSummarizedUsagesRequest{
        RequestSummarizedUsagesDetails: usageapi.RequestSummarizedUsagesDetails{
            Granularity:       "MONTHLY",
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
        },
    }

    // Executa a requisição
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    if err != nil {
        log.Fatal("Erro ao solicitar dados de uso:", err)
    }

    total := 0.0
    for _, item := range response.Items {
        if item.ComputedAmount != nil {
            total += float64(*item.ComputedAmount)
        }
    }

    return total
}