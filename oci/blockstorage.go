package oci

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

// Este adiciona a Region ao BootVolume, que originalmente não tem
type BootVolume struct {
    core.BootVolume // Incorporando a struct BootVolume

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region ao Volume, que originalmente não tem
type Volume struct {
    core.Volume // Incorporando a struct Volume

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Obtem os dados de um BootVolume pelo OCID e pela região quando fornecida.
func GetBootVolume(bootVolumeOCID string, region ...string) (BootVolume, error) {
	client, err := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		return BootVolume{}, err
	}

	// Configura a região se fornecida
	if len(region) > 0 {
		client.SetRegion(region[0])
	}

	var extendedBootVolume BootVolume

	err = RetryWithBackoff(5, func() error {
		ctx := context.Background()
		req := core.GetBootVolumeRequest{
			BootVolumeId: common.String(bootVolumeOCID),
		}

		resp, innerErr := client.GetBootVolume(ctx, req)
		if innerErr != nil {
			return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
		}

		// Preenchendo os dados obtidos na struct ExtendedBootVolume
		extendedBootVolume = BootVolume{
			BootVolume: resp.BootVolume,
			Region:     nil,
		}
		if len(region) > 0 {
			extendedBootVolume.Region = common.String(region[0])
		}

		return nil
	})

	if err != nil {
		return BootVolume{}, err
	}

	return extendedBootVolume, nil
}


// ListAllBootVolumes retorna detalhes de todos os boot volumes em todas as regiões na OCI.
func ListAllBootVolumes() ([]BootVolume, error) {
    var allBootVolumes []BootVolume

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query BootVolume resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var bootVolumeDetails BootVolume

                    err := RetryWithBackoff(5, func() error {
                        var innerErr error
                        bootVolumeDetails, innerErr = GetBootVolume(*summary.Identifier, *region.RegionName)
                        return innerErr
                    })

                    if err != nil {
                        continue // ou retorne, dependendo da lógica de erro
                    }

                    allBootVolumes = append(allBootVolumes, bootVolumeDetails)
                }
            }
        }
    }

    return allBootVolumes, nil
}

// GetBootVolumeAttachment retorna detalhes de um boot volume attachment específico usando seu OCID.
func GetBootVolumeAttachment(attachmentOCID string) (core.BootVolumeAttachment, error) {
    var attachment core.BootVolumeAttachment

    client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.BootVolumeAttachment{}, err
    }

    ctx := context.Background()

    // Usando RetryWithBackoff para retentativas
    err = RetryWithBackoff(5, func() error {
        req := core.GetBootVolumeAttachmentRequest{
            BootVolumeAttachmentId: common.String(attachmentOCID),
        }

        resp, innerErr := client.GetBootVolumeAttachment(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        attachment = resp.BootVolumeAttachment
        return nil
    })

    if err != nil {
        return core.BootVolumeAttachment{}, err
    }

    return attachment, nil
}


func ListAllBootVolumeAttachments() ([]core.BootVolumeAttachment, error) {
    var allAttachments []core.BootVolumeAttachment

    computeClient, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    ctx := context.Background()

    for _, region := range regions {
        if region.RegionName != nil {
            computeClient.SetRegion(*region.RegionName)

            resourceSummaries, err := ResourceSearch("query BootVolume resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    request := core.ListBootVolumeAttachmentsRequest{
                        CompartmentId:      summary.CompartmentId,
                        AvailabilityDomain: summary.AvailabilityDomain,
                        BootVolumeId:       summary.Identifier,
                    }

                    err := RetryWithBackoff(5, func() error {
                        var innerErr error
                        for {
                            response, innerErr := computeClient.ListBootVolumeAttachments(ctx, request)
                            if innerErr != nil {
                                break // Sai do loop interno e permite que retryWithBackoff tente novamente
                            }

                            allAttachments = append(allAttachments, response.Items...)
                            if response.OpcNextPage == nil {
                                break // Sai do loop interno, pois concluímos a paginação
                            }
                            request.Page = response.OpcNextPage
                        }
                        return innerErr
                    })

                    if err != nil {
                        return nil, err // Retorna o erro se todas as retentativas falharem
                    }
                }
            }
        }
    }
    return allAttachments, nil
}

// GetBlockVolume retorna detalhes de um block volume específico usando seu OCID.
// O parâmetro region é opcional. Se fornecido, a consulta será realizada na região especificada.
func GetBlockVolume(blockVolumeOCID string, region ...string) (Volume, error) {
    client, err := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return Volume{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    var extendedVolume Volume
    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetVolumeRequest{
            VolumeId: common.String(blockVolumeOCID),
        }

        resp, innerErr := client.GetVolume(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        extendedVolume = Volume{
            Volume: resp.Volume,
            Region: nil,
        }
        if len(region) > 0 {
            extendedVolume.Region = common.String(region[0])
        }

        return nil
    })

    if err != nil {
        return Volume{}, err
    }

    return extendedVolume, nil
}

func ListAllBlockVolumeAttachments() ([]core.VolumeAttachment, error) {
    var allAttachments []core.VolumeAttachment

    computeClient, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    ctx := context.Background()

    for _, region := range regions {
        if region.RegionName != nil {
            computeClient.SetRegion(*region.RegionName)

            resourceSummaries, err := ResourceSearch("query Volume resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    request := core.ListVolumeAttachmentsRequest{
                        CompartmentId:      summary.CompartmentId,
                        AvailabilityDomain: summary.AvailabilityDomain,
                        VolumeId:       summary.Identifier,
                    }

                    err := RetryWithBackoff(5, func() error {
                        var innerErr error
                        for {
                            response, innerErr := computeClient.ListVolumeAttachments(ctx, request)
                            if innerErr != nil {
                                break // Sai do loop interno e permite que retryWithBackoff tente novamente
                            }

                            allAttachments = append(allAttachments, response.Items...)
                            if response.OpcNextPage == nil {
                                break // Sai do loop interno, pois concluímos a paginação
                            }
                            request.Page = response.OpcNextPage
                        }
                        return innerErr
                    })

                    if err != nil {
                        return nil, err // Retorna o erro se todas as retentativas falharem
                    }
                }
            }
        }
    }
    return allAttachments, nil
}

// GetBlockVolumeAttachment retorna detalhes de um volume attachment específico usando seu OCID.
func GetBlockVolumeAttachment(volumeAttachmentOCID string) (core.VolumeAttachment, error) {
    client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    var attachment core.VolumeAttachment
    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetVolumeAttachmentRequest{
            VolumeAttachmentId: common.String(volumeAttachmentOCID),
        }

        resp, innerErr := client.GetVolumeAttachment(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        attachment = resp.VolumeAttachment
        return nil
    })

    if err != nil {
        return nil, err
    }

    return attachment, nil
}

// ListAllBlockVolumes retorna detalhes de todos os block volumes em todas as regiões na OCI.
func ListAllBlockVolumes() ([]Volume, error) {
    var allBlockVolumes []Volume

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query Volume resources", *region.RegionName)
            if err != nil {
                continue // ou retorne, dependendo da lógica de erro
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var blockVolumeDetails Volume

                    err := RetryWithBackoff(5, func() error {
                        var innerErr error
                        blockVolumeDetails, innerErr = GetBlockVolume(*summary.Identifier, *region.RegionName)
                        return innerErr
                    })

                    if err != nil {
                        continue // ou retorne, dependendo da lógica de erro
                    }

                    allBlockVolumes = append(allBlockVolumes, blockVolumeDetails)
                }
            }
        }
    }

    return allBlockVolumes, nil
}