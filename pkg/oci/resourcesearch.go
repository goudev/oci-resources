package oci

import (
	"context"
	"oci-sdk-go/pkg/util"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/resourcesearch"
)

// ResourceSearch executa uma consulta estruturada na Oracle Cloud e retorna os resultados em JSON.
// O parâmetro region é opcional. Se fornecido, a pesquisa será realizada na região especificada.
func ResourceSearch(query string, region ...string) ([]resourcesearch.ResourceSummary, error) {
    client, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    var allResults []resourcesearch.ResourceSummary
    var currentPage *string = nil

    for {
        err = util.RetryWithBackoff(5, func() error {
            ctx := context.Background()
            searchReq := resourcesearch.SearchResourcesRequest{
                SearchDetails: resourcesearch.StructuredSearchDetails{
                    Query: common.String(query),
                },
                Page: currentPage,
            }

            resp, innerErr := client.SearchResources(ctx, searchReq)
            if innerErr != nil {
                return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
            }

            allResults = append(allResults, resp.Items...)

            if resp.OpcNextPage == nil || (currentPage != nil && *resp.OpcNextPage == *currentPage) {
                return nil // Sai do loop de retentativas
            }

            currentPage = resp.OpcNextPage
            return nil
        })

        if err != nil {
            return nil, err // Retorna o erro se todas as retentativas falharem
        }

        if currentPage == nil {
            break // Sai do loop principal se concluímos a paginação
        }
    }

    return allResults, nil
}
