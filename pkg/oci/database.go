package oci

import (
	"context"
	"oci-sdk-go/pkg/util"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func ListAllDatabases() ([]database.Database, error) {
    var allDatabases []database.Database

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            // Obtem os resumos dos recursos do tipo 'database' para a região atual
            resourceSummaries, err := ResourceSearch("query database resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.ResourceType != nil && *summary.ResourceType == "Database" && summary.Identifier != nil {
                    var databaseDetails database.Database
                    err := util.RetryWithBackoff(5, func() error {
                        dbClient, innerErr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        dbClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := database.GetDatabaseRequest{
                            DatabaseId: summary.Identifier,
                        }

                        resp, innerErr := dbClient.GetDatabase(ctx, req)
                        if innerErr != nil {
                            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
                        }

                        databaseDetails = resp.Database
                        return nil
                    })

                    if err != nil {
                        continue // ou retorne, dependendo da lógica de erro
                    }

                    allDatabases = append(allDatabases, databaseDetails)
                }
            }
        }
    }

    return allDatabases, nil
}

func ListAllCloudExadataInfrastructures() ([]database.CloudExadataInfrastructure, error) {
    var allCloudExadataInfrastructures []database.CloudExadataInfrastructure

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            // Obtem os resumos dos recursos do tipo 'CloudExadataInfrastructure' para a região atual
            resourceSummaries, err := ResourceSearch("query CloudExadataInfrastructure resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.ResourceType != nil && *summary.ResourceType == "CloudExadataInfrastructure" && summary.Identifier != nil {
                    var cloudExadataInfrastructure database.CloudExadataInfrastructure
                    err := util.RetryWithBackoff(5, func() error {
                        dbClient, innerErr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        dbClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := database.GetCloudExadataInfrastructureRequest{
                            CloudExadataInfrastructureId: summary.Identifier,
                        }

                        resp, innerErr := dbClient.GetCloudExadataInfrastructure(ctx, req)
                        if innerErr != nil {
                            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
                        }

                        cloudExadataInfrastructure = resp.CloudExadataInfrastructure
                        return nil
                    })

                    if err != nil {
                        continue // ou retorne, dependendo da lógica de erro
                    }

                    allCloudExadataInfrastructures = append(allCloudExadataInfrastructures, cloudExadataInfrastructure)
                }
            }
        }
    }

    return allCloudExadataInfrastructures, nil
}

func ListAllDbNodes() ([]database.DbNode, error) {
    var allDbNodes []database.DbNode

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            // Obtem os resumos dos recursos do tipo 'DbNode' para a região atual
            resourceSummaries, err := ResourceSearch("query DbNode resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.ResourceType != nil && *summary.ResourceType == "DbNode" && summary.Identifier != nil {
                    var dbNode database.DbNode
                    err := util.RetryWithBackoff(5, func() error {
                        dbClient, innerErr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        dbClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := database.GetDbNodeRequest{
                            DbNodeId: summary.Identifier,
                        }

                        resp, innerErr := dbClient.GetDbNode(ctx, req)
                        if innerErr != nil {
                            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
                        }

                        dbNode = resp.DbNode
                        return nil
                    })

                    if err != nil {
                        continue // ou retorne, dependendo da lógica de erro
                    }

                    allDbNodes = append(allDbNodes, dbNode)
                }
            }
        }
    }

    return allDbNodes, nil
}

func ListAllCloudVmClusters() ([]database.CloudVmCluster, error) {
    var allCloudVmClusters []database.CloudVmCluster

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            // Obtem os resumos dos recursos do tipo 'CloudVmCluster' para a região atual
            resourceSummaries, err := ResourceSearch("query CloudVmCluster resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.ResourceType != nil && *summary.ResourceType == "CloudVmCluster" && summary.Identifier != nil {
                    var cloudVmCluster database.CloudVmCluster
                    err := util.RetryWithBackoff(5, func() error {
                        dbClient, innerErr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        dbClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := database.GetCloudVmClusterRequest{
                            CloudVmClusterId: summary.Identifier,
                        }

                        resp, innerErr := dbClient.GetCloudVmCluster(ctx, req)
                        if innerErr != nil {
                            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
                        }

                        cloudVmCluster = resp.CloudVmCluster
                        return nil
                    })

                    if err != nil {
                        continue // ou retorne, dependendo da lógica de erro
                    }

                    allCloudVmClusters = append(allCloudVmClusters, cloudVmCluster)
                }
            }
        }
    }

    return allCloudVmClusters, nil
}