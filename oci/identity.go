package oci

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func ListAllCompartments() ([]identity.Compartment, error) {
    client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
    if err != nil {
        return nil, err
    }

    var compartments []identity.Compartment
    var currentPage *string = nil

    for {
        err = RetryWithBackoff(5, func() error {
            ctx := context.Background()
            request := identity.ListCompartmentsRequest{
                CompartmentId:          common.String(tenancyID),
                AccessLevel:            identity.ListCompartmentsAccessLevelAny,
                CompartmentIdInSubtree: common.Bool(true),
                Page:                   currentPage,
            }

            response, innerErr := client.ListCompartments(ctx, request)
            if innerErr != nil {
                return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
            }

            compartments = append(compartments, response.Items...)
            if response.OpcNextPage == nil {
                return nil // Sai do loop de retentativas
            }

            currentPage = response.OpcNextPage
            return nil
        })

        if err != nil {
            return nil, err // Retorna o erro se todas as retentativas falharem
        }

        if currentPage == nil {
            break // Sai do loop principal se concluímos a paginação
        }
    }

    return compartments, nil
}

