package oci

import (
	"context"
	"oci-sdk-go/pkg/util"

	"github.com/oracle/oci-go-sdk/v65/analytics" // Este import pode variar
	"github.com/oracle/oci-go-sdk/v65/common"
)

type AnalyticsInstance struct {
    analytics.AnalyticsInstance // Incorporação anônima da struct AnalyticsInstance
    Region *string `json:"region"`
}

func GetAnalyticsInstance(analyticsInstanceOCID string, region ...string) (analytics.AnalyticsInstance, error) {
    client, err := analytics.NewAnalyticsClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return analytics.AnalyticsInstance{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    ctx := context.Background()
    req := analytics.GetAnalyticsInstanceRequest{
        AnalyticsInstanceId: common.String(analyticsInstanceOCID),
    }

    resp, err := client.GetAnalyticsInstance(ctx, req)
    if err != nil {
        return analytics.AnalyticsInstance{}, err
    }

    // Retorne a instância de análise obtida aqui
    return resp.AnalyticsInstance, nil
}


// ListAllAnalyticsInstances lista todas as AnalyticsInstances nas regiões disponíveis
func ListAllAnalyticsInstances() ([]AnalyticsInstance, error) {
    var allAnalyticsInstances []AnalyticsInstance
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query AnalyticsInstance resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var analyticsInstance AnalyticsInstance
                    err := util.RetryWithBackoff(5, func() error {
                        client, innerErr := analytics.NewAnalyticsClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        client.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := analytics.GetAnalyticsInstanceRequest{
                            AnalyticsInstanceId: summary.Identifier,
                        }

                        resp, innerErr := client.GetAnalyticsInstance(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        analyticsInstance = AnalyticsInstance{
                            AnalyticsInstance: resp.AnalyticsInstance, // Ajuste para a incorporação anônima
                            Region:            region.RegionName,
                        }
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allAnalyticsInstances = append(allAnalyticsInstances, analyticsInstance)
                }
            }
        }
    }

    return allAnalyticsInstances, nil
}