package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/usageapi"
)

// SummaryDataWithNewResources é uma estrutura para armazenar os dados resumidos com informações adicionais
type SummaryDataWithNewResources struct {
    Referencia          string
    AmountTotal         float32
    AmountNewResources  float32
    AmountPartial       float32
}

// SummaryData é a estrutura para armazenar os dados resumidos
type SummaryData struct {
    Referencia string
    Amount     float32
}

// Supondo que ResourceSearch retorna uma slice de uma estrutura com um campo 'Identifier'
type Resource struct {
    Identifier string
}

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

// ParseAndZeroTime é uma função para analisar e zerar a hora de uma string de data no formato "aaaa/mm/dd"
func ParseAndZeroTime(dateStr string) (time.Time, error) {
    // Troca "/" por "-" para compatibilidade com o formato de data do Go
    normalizedDateStr := strings.Replace(dateStr, "/", "-", -1)
    parsedTime, err := time.Parse("2006-01-02", normalizedDateStr)
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

// GetCostByFilter obtém o resumo de uso por SKU entre datas específicas com múltiplos OCIDs
func GetCostByFilter(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum, ocids []string) []SummaryData {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatalf("Erro ao analisar a data inicial: %v", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatalf("Erro ao analisar a data final: %v", err)
    }

    if granularity == "DAILY" && endDate.Sub(startDate).Hours() > 90*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 90 dias.")
    }

    if granularity == "MONTHLY" && endDate.Sub(startDate).Hours() > 365*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 365 dias.")
    }

    endDate = endDate.AddDate(0, 0, 1)

    configProvider := common.CustomProfileConfigProvider("", "")
    tenancyID, _ := configProvider.TenancyOCID()
    client, err := usageapi.NewUsageapiClientWithConfigurationProvider(configProvider)
    if err != nil {
        log.Fatalf("Erro ao criar o cliente da API de Uso: %v", err)
    }

    var allUsageSummaries []SummaryData

    // Verifica se a lista de OCIDs é nil ou vazia
    if ocids == nil || len(ocids) == 0 {
        // Cria a requisição sem filtro de OCID específico
        req := usageapi.RequestSummarizedUsagesRequest{
            RequestSummarizedUsagesDetails: usageapi.RequestSummarizedUsagesDetails{
                Granularity:       granularity,
                TenantId:          common.String(*common.String(tenancyID)),
                TimeUsageStarted:  &common.SDKTime{Time: startDate},
                TimeUsageEnded:    &common.SDKTime{Time: endDate},
                CompartmentDepth:  common.Float32(2),
                IsAggregateByTime: common.Bool(false),
                QueryType:         "COST",
            },
        }

        response, err := client.RequestSummarizedUsages(context.Background(), req)
        if err != nil {
            log.Printf("Erro ao solicitar dados de uso: %v", err)
            return nil
        }

        for _, item := range response.Items {
            if *item.Currency != " " {
                var ref string
                if granularity == "DAILY" {
                    ref = item.TimeUsageStarted.Format("2006/01/02")
                } else {
                    ref = item.TimeUsageStarted.Format("2006/01")
                }
                allUsageSummaries = append(allUsageSummaries, SummaryData{
                    Referencia: ref,
                    Amount:     *item.ComputedAmount,
                })
            }
        }
    } else {
        // Processa cada OCID na lista fornecida
        for _, ocid := range ocids {
            req := usageapi.RequestSummarizedUsagesRequest{
                RequestSummarizedUsagesDetails: usageapi.RequestSummarizedUsagesDetails{
                    Granularity:       granularity,
                    TenantId:          common.String(*common.String(tenancyID)),
                    TimeUsageStarted:  &common.SDKTime{Time: startDate},
                    TimeUsageEnded:    &common.SDKTime{Time: endDate},
                    CompartmentDepth:  common.Float32(2),
                    IsAggregateByTime: common.Bool(false),
                    QueryType:         "COST",
                    Filter: &usageapi.Filter{
                        Dimensions: []usageapi.Dimension{
                            {Key: common.String("resourceId"), Value: common.String(ocid)},
                        },
                        Operator: usageapi.FilterOperatorAnd,
                    },
                },
            }

            response, err := client.RequestSummarizedUsages(context.Background(), req)
            if err != nil {
                log.Printf("Erro ao solicitar dados de uso para OCID %s: %v", ocid, err)
                continue
            }

            for _, item := range response.Items {
                if *item.Currency != " " {
                    var ref string
                    if granularity == "DAILY" {
                        ref = item.TimeUsageStarted.Format("2006/01/02")
                    } else {
                        ref = item.TimeUsageStarted.Format("2006/01")
                    }
                    allUsageSummaries = append(allUsageSummaries, SummaryData{
                        Referencia: ref,
                        Amount:     *item.ComputedAmount,
                    })
                }
            }
        }
    }

    return allUsageSummaries
}

// Converte datas do formato 'aaaa/mm/dd' para o formato ISO 8601 'yyyy-mm-ddTHH:MM:SSZ'
func convertToISO8601(dateStr string) string {
    normalizedDateStr := strings.Replace(dateStr, "/", "-", -1)
    return normalizedDateStr + "T00:00:00Z"
}

// GetCostByNewResources obtém o custo de novos recursos criados em um intervalo de tempo
func GetCostByNewResources(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []SummaryData {
    startDateISO := convertToISO8601(startDateStr)
    endDateISO := convertToISO8601(endDateStr)

    query := fmt.Sprintf("query instance, BootVolume, Volume, AnalyticsInstance, VolumeBackup, VolumeReplica, BootVolumeBackup, BootVolumeReplica, Bucket, FileSystem, Database resources where timeCreated >= '%s' && timeCreated <= '%s'", startDateISO, endDateISO)
    resources, err := ResourceSearch(query)
    if err != nil {
        log.Fatalf("Erro ao buscar recursos: %v", err)
    }

    var ocids []string
    for _, resource := range resources {
        ocids = append(ocids, *resource.Identifier)
    }

    if(len(ocids)>=0){
        return GetCostByFilter(startDateStr, endDateStr, granularity, ocids)
    }else{
        return nil
    }
}

// GetCostSummaryDataWithNewResources retorna um resumo dos custos
func GetCostSummaryDataWithNewResources(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []SummaryDataWithNewResources {
    
    totalCostData := GetCostByFilter(startDateStr, endDateStr, granularity, nil) // nil significa sem filtro de OCID específico
    newResourcesCostData := GetCostByNewResources(startDateStr, endDateStr, granularity)

    summaryMap := make(map[string]SummaryDataWithNewResources)

    // Processa o custo total
    for _, data := range totalCostData {
        summaryMap[data.Referencia] = SummaryDataWithNewResources{
            Referencia:      data.Referencia,
            AmountTotal:     data.Amount,
            AmountNewResources:  0,
            AmountPartial:       0,
        }
    }

    // Processa o custo de novos recursos
    for _, newData := range newResourcesCostData {
        if summary, ok := summaryMap[newData.Referencia]; ok {
            summary.AmountNewResources = newData.Amount
            summary.AmountPartial = summary.AmountTotal - newData.Amount
            summaryMap[newData.Referencia] = summary
        }
    }

    var summaryData []SummaryDataWithNewResources
    for _, summary := range summaryMap {
        summaryData = append(summaryData, summary)
    }

    return summaryData
}

// Função para obter o custo entre datas específica
func getCostByService(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []usageapi.UsageSummary {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data inicial:", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data final:", err)
    }

	 if granularity=="DAILY" && endDate.Sub(startDate).Hours() > 90*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 90 dias.")
    }

    if granularity=="MONTHLY" && endDate.Sub(startDate).Hours() > 365*24 {
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
            Granularity:       granularity,
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
            GroupBy: []string{"service"},
        },
    }

    // Executa a requisição
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    if err != nil {
        log.Fatal("Erro ao solicitar dados de uso:", err)
    }

    return response.Items
}

// Função para obter o custo entre datas específica
func getCostBySku(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []usageapi.UsageSummary {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data inicial:", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data final:", err)
    }

	 if granularity=="DAILY" && endDate.Sub(startDate).Hours() > 90*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 90 dias.")
    }

    if granularity=="MONTHLY" && endDate.Sub(startDate).Hours() > 365*24 {
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
            Granularity:       granularity,
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
            GroupBy: []string{"skuName"},
        },
    }

    // Executa a requisição
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    if err != nil {
        log.Fatal("Erro ao solicitar dados de uso:", err)
    }

    return response.Items
}

// Função para obter o custo entre datas específica
func getCostByRegion(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []usageapi.UsageSummary {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data inicial:", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data final:", err)
    }

	 if granularity=="DAILY" && endDate.Sub(startDate).Hours() > 90*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 90 dias.")
    }

    if granularity=="MONTHLY" && endDate.Sub(startDate).Hours() > 365*24 {
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
            Granularity:       granularity,
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
            GroupBy: []string{"region"},
        },
    }

    // Executa a requisição
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    if err != nil {
        log.Fatal("Erro ao solicitar dados de uso:", err)
    }

    return response.Items
}

// Função para obter o custo entre datas específica
func getCostByCompartmentPath(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []usageapi.UsageSummary {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data inicial:", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data final:", err)
    }

	 if granularity=="DAILY" && endDate.Sub(startDate).Hours() > 90*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 90 dias.")
    }

    if granularity=="MONTHLY" && endDate.Sub(startDate).Hours() > 365*24 {
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
            Granularity:       granularity,
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
            GroupBy: []string{"compartmentPath"},
        },
    }

    // Executa a requisição
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    if err != nil {
        log.Fatal("Erro ao solicitar dados de uso:", err)
    }

    return response.Items
}

// Função para obter o custo entre datas específica
func GetCost(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum,group string) []usageapi.UsageSummary {

    switch group {
    case "SKU":
        return getCostBySku(startDateStr,endDateStr,granularity)
    case "COMPARTMENT_PATH":
        return getCostByCompartmentPath(startDateStr,endDateStr,granularity)
    case "REGION":
        return getCostByRegion(startDateStr,endDateStr,granularity)
    case "SERVICE":
        return getCostByService(startDateStr,endDateStr,granularity)
    default:
        return getCostByService(startDateStr,endDateStr,granularity)
    }
}


// Função para obter o custo entre datas específicas
func GGetFullCost(startDateStr, endDateStr string,granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum, ocids []string) float64 {
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
            Granularity:       granularity,
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
            GroupBy: []string{"skuName"},
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

func GetFullCost(startDateStr, endDateStr string, granularity usageapi.RequestSummarizedUsagesDetailsGranularityEnum) []usageapi.UsageSummary {
    startDate, err := ParseAndZeroTime(startDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data inicial:", err)
    }
    endDate, err := ParseAndZeroTime(endDateStr)
    if err != nil {
        log.Fatal("Erro ao analisar a data final:", err)
    }

	 if granularity=="DAILY" && endDate.Sub(startDate).Hours() > 90*24 {
        log.Fatal("Erro: O intervalo de datas não pode ser maior que 90 dias.")
    }

    if granularity=="MONTHLY" && endDate.Sub(startDate).Hours() > 365*24 {
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
            Granularity:       granularity,
            TenantId:          common.String(*common.String(tenancyID)),
            TimeUsageStarted:  &common.SDKTime{Time: startDate},
            TimeUsageEnded:    &common.SDKTime{Time: endDate},
            CompartmentDepth:  common.Float32(2),
            IsAggregateByTime: common.Bool(false),
            QueryType:         "COST",
            GroupBy: []string{"compartmentPath"},
        },
    }

    // Executa a requisição
    response, err := client.RequestSummarizedUsages(context.Background(), req)
    if err != nil {
        log.Fatal("Erro ao solicitar dados de uso:", err)
    }

    return response.Items
}
