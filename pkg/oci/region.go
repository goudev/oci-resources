package oci

import (
	"context"
	"oci-sdk-go/pkg/util"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func GetAllRegions() ([]identity.RegionSubscription, error) {
    client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    var regionSubscriptions []identity.RegionSubscription
    err = util.RetryWithBackoff(5, func() error {
        ctx := context.Background()

        // Primeiro, obtemos o OCID da tenancy do cliente
        tenancyID, innerErr := common.DefaultConfigProvider().TenancyOCID()
        if innerErr != nil {
            return innerErr
        }

        // Listar as subscrições de região para a tenancy
        req := identity.ListRegionSubscriptionsRequest{
            TenancyId: common.String(tenancyID),
        }

        resp, innerErr := client.ListRegionSubscriptions(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        regionSubscriptions = resp.Items
        return nil
    })

    if err != nil {
        return nil, err
    }

    return regionSubscriptions, nil
}
