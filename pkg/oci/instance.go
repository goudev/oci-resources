package oci

import (
	"context"
	"oci-sdk-go/pkg/util"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func GetInstance(instanceOcid string, region ...string) (core.Instance, error) {
    client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.Instance{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    var instance core.Instance
    err = util.RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetInstanceRequest{
            InstanceId: common.String(instanceOcid),
        }

        resp, innerErr := client.GetInstance(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        instance = resp.Instance
        return nil
    })

    if err != nil {
        return core.Instance{}, err
    }

    return instance, nil
}


func GetAllInstances() ([]core.Instance, error) {
    var allInstances []core.Instance

    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query instance resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var instanceDetails core.Instance
                    err := util.RetryWithBackoff(5, func() error {
                        var innerErr error
                        instanceDetails, innerErr = GetInstance(*summary.Identifier, *region.RegionName)
                        return innerErr
                    })

                    if err != nil {
                        continue
                    }

                    allInstances = append(allInstances, instanceDetails)
                }
            }
        }
    }

    return allInstances, nil
}
