package oci

import (
	"context"
	"log"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

// Este adiciona a Region ao Vcn, que originalmente não tem
type Vcn struct {
    core.Vcn // Incorporando a struct Vcn

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region a Subnet, que originalmente não tem
type Subnet struct {
    core.Subnet // Incorporando a struct Subnet

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region a NatGateway, que originalmente não tem
type NatGateway struct {
    core.NatGateway // Incorporando a struct NatGateway

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region a SecurityList, que originalmente não tem
type SecurityList struct {
    core.SecurityList // Incorporando a struct SecurityList

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region a RouteTable, que originalmente não tem
type RouteTable struct {
    core.RouteTable // Incorporando a struct RouteTable

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region a ServiceGateway, que originalmente não tem
type ServiceGateway struct {
    core.ServiceGateway // Incorporando a struct ServiceGateway

    // Adicionando o novo campo Region
    Region *string `json:"region"`
}

// Este adiciona a Region a InternetGateway, que originalmente não tem
type InternetGateway struct {
    core.InternetGateway

    // Adiciona o novo cmpo Region
    Region *string `json:"region"`
}

// Definição da struct personalizada
type Drg struct {
    core.Drg
    Region *string `json:"region"`
}

// Definição da struct personalizada para DrgAttachment
type DrgAttachment struct {
    core.DrgAttachment
    Region *string `json:"region"`
}

// Definição da struct personalizada para DrgRouteTable
type DrgRouteTable struct {
    core.DrgRouteTable
    Region *string `json:"region"`
}

// Definição da struct personalizada para DrgRouteDistribution
type DrgRouteDistribution struct {
    core.DrgRouteDistribution
    Region *string `json:"region"`
}

// Definição da struct personalizada para LocalPeeringGateway
type LocalPeeringGateway struct {
    core.LocalPeeringGateway
    Region *string `json:"region"`
}

type DhcpDnsOption struct {
    core.DhcpOptions
    Region *string `json:"region"`
}

// Definição da struct personalizada para NetworkSecurityGroup
type NetworkSecurityGroup struct {
    core.NetworkSecurityGroup
    Region *string `json:"region"`
}

// Definição da struct personalizada para IpSecConnection
type IpSecConnection struct {
    core.IpSecConnection
    Region *string `json:"region"`
}

// Definição da struct personalizada para Cpe
type Cpe struct {
    core.Cpe
    Region *string `json:"region"`
}

// Obtem uma VCN pelo ID
func GetVCN(vcnOCID string, region ...string) (core.Vcn, error) {
    var vcn core.Vcn

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.Vcn{}, err
    }

    // Configura a região se fornecida
    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    // Usando RetryWithBackoff para retentativas
    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetVcnRequest{
            VcnId: common.String(vcnOCID),
        }

        resp, innerErr := client.GetVcn(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        vcn = resp.Vcn
        return nil
    })

    if err != nil {
        return core.Vcn{}, err
    }

    return vcn, nil
}

// Obtem todas as VCNs em todas as regiões
func ListAllVCNs() ([]Vcn, error) {
    var allVCNs []Vcn

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query Vcn resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.ResourceType != nil && *summary.ResourceType == "Vcn" && summary.Identifier != nil {
                    var extendedVcn Vcn
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetVcnRequest{
                            VcnId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetVcn(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        extendedVcn = Vcn{
                            Vcn:    resp.Vcn,
                            Region: common.String(*region.RegionName),
                        }
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allVCNs = append(allVCNs, extendedVcn)
                }
            }
        }
    }

    return allVCNs, nil
}

// Obtem uma subnet pelo ocid
func GetSubnet(subnetOCID string, region ...string) (core.Subnet, error) {
    var subnet core.Subnet

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.Subnet{}, err
    }

    // Configura a região se fornecida
    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    // Usando RetryWithBackoff para retentativas
    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetSubnetRequest{
            SubnetId: common.String(subnetOCID),
        }

        resp, innerErr := client.GetSubnet(ctx, req)
        if innerErr != nil {
            return innerErr // Retorna o erro para que RetryWithBackoff possa decidir retentar
        }

        subnet = resp.Subnet
        return nil
    })

    if err != nil {
        return core.Subnet{}, err
    }

    return subnet, nil
}

// Obtem todas as subnets de todas as regiões
func ListAllSubnets() ([]Subnet, error) {
    var allSubnets []Subnet

    // Obtém a lista de todas as regiões
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query Subnet resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var extendedSubnet Subnet
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetSubnetRequest{
                            SubnetId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetSubnet(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        extendedSubnet = Subnet{
                            Subnet: resp.Subnet,
                            Region: common.String(*region.RegionName),
                        }
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allSubnets = append(allSubnets, extendedSubnet)
                }
            }
        }
    }

    return allSubnets, nil
}

// Ontem um natgateway
func GetNatGateway(natGatewayOCID string, region ...string) (core.NatGateway, error) {
    var natGateway core.NatGateway

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.NatGateway{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetNatGatewayRequest{
            NatGatewayId: common.String(natGatewayOCID),
        }

        resp, innerErr := client.GetNatGateway(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        natGateway = resp.NatGateway
        return nil
    })

    if err != nil {
        return core.NatGateway{}, err
    }

    return natGateway, nil
}

// Obtem a lista de todos os natgateway
func ListAllNatGateways() ([]NatGateway, error) {
    var allNatGateways []NatGateway
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query NatGateway resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var natGateway NatGateway
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetNatGatewayRequest{
                            NatGatewayId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetNatGateway(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        natGateway = NatGateway{NatGateway: resp.NatGateway, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allNatGateways = append(allNatGateways, natGateway)
                }
            }
        }
    }
    return allNatGateways, nil
}

// Obtem um service gateway pelo ocid
func GetServiceGateway(serviceGatewayOCID string, region ...string) (core.ServiceGateway, error) {
    var serviceGateway core.ServiceGateway

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.ServiceGateway{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetServiceGatewayRequest{
            ServiceGatewayId: common.String(serviceGatewayOCID),
        }

        resp, innerErr := client.GetServiceGateway(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        serviceGateway = resp.ServiceGateway
        return nil
    })

    if err != nil {
        return core.ServiceGateway{}, err
    }

    return serviceGateway, nil
}

// Obtem todos os service gateways em todas as regiões
func ListAllServiceGateways() ([]ServiceGateway, error) {
    var allServiceGateways []ServiceGateway
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query ServiceGateway resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var serviceGateway ServiceGateway
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetServiceGatewayRequest{
                            ServiceGatewayId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetServiceGateway(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        serviceGateway = ServiceGateway{ServiceGateway: resp.ServiceGateway, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allServiceGateways = append(allServiceGateways, serviceGateway)
                }
            }
        }
    }
    return allServiceGateways, nil
}

// Obtem uma route table pelo ocid
func GetRouteTable(routeTableOCID string, region ...string) (core.RouteTable, error) {
    var routeTable core.RouteTable

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.RouteTable{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetRouteTableRequest{
            RtId: common.String(routeTableOCID),
        }

        resp, innerErr := client.GetRouteTable(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        routeTable = resp.RouteTable
        return nil
    })

    if err != nil {
        return core.RouteTable{}, err
    }

    return routeTable, nil
}

// Obtem todas as route tables de uma região
func ListAllRouteTables() ([]RouteTable, error) {
    var allRouteTables []RouteTable
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query RouteTable resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var routeTable RouteTable
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetRouteTableRequest{
                            RtId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetRouteTable(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        routeTable = RouteTable{RouteTable: resp.RouteTable, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allRouteTables = append(allRouteTables, routeTable)
                }
            }
        }
    }
    return allRouteTables, nil
}

// Obtem a security list pelo ocid
func GetSecurityList(securityListOCID string, region ...string) (core.SecurityList, error) {
    var securityList core.SecurityList

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.SecurityList{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetSecurityListRequest{
            SecurityListId: common.String(securityListOCID),
        }

        resp, innerErr := client.GetSecurityList(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        securityList = resp.SecurityList
        return nil
    })

    if err != nil {
        return core.SecurityList{}, err
    }

    return securityList, nil
}

// Obtem a lista de todas as security lists de uma região
func ListAllSecurityLists() ([]SecurityList, error) {
    var allSecurityLists []SecurityList
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query SecurityList resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var securityList SecurityList
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetSecurityListRequest{
                            SecurityListId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetSecurityList(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        securityList = SecurityList{SecurityList: resp.SecurityList, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allSecurityLists = append(allSecurityLists, securityList)
                }
            }
        }
    }
    return allSecurityLists, nil
}

// Obtem um internet gateway pelo ocid
func GetInternetGateway(internetGatewayOCID string, region ...string) (core.InternetGateway, error) {
    var internetGateway core.InternetGateway

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.InternetGateway{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetInternetGatewayRequest{
            IgId: common.String(internetGatewayOCID),
        }

        resp, innerErr := client.GetInternetGateway(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        internetGateway = resp.InternetGateway
        return nil
    })

    if err != nil {
        return core.InternetGateway{}, err
    }

    return internetGateway, nil
}

// Obtem a lista de todos os internet gateway de todas as regiões
func ListAllInternetGateways() ([]InternetGateway, error) {
    var allInternetGateways []InternetGateway
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query InternetGateway resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var internetGateway InternetGateway
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetInternetGatewayRequest{
                            IgId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetInternetGateway(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        internetGateway = InternetGateway{InternetGateway: resp.InternetGateway, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allInternetGateways = append(allInternetGateways, internetGateway)
                }
            }
        }
    }

    return allInternetGateways, nil
}

// Obtem um DRG a partir do OCID
func GetDrg(drgOCID string, region ...string) (core.Drg, error) {
    var drg core.Drg

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.Drg{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetDrgRequest{
            DrgId: common.String(drgOCID),
        }

        resp, innerErr := client.GetDrg(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        drg = resp.Drg
        return nil
    })

    if err != nil {
        return core.Drg{}, err
    }

    return drg, nil
}

// Obtem a lista de todos os DRGS
func ListAllDrgs() ([]Drg, error) {
    var allDrgs []Drg
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query Drg resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var drg Drg
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetDrgRequest{
                            DrgId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetDrg(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        drg = Drg{Drg: resp.Drg, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allDrgs = append(allDrgs, drg)
                }
            }
        }
    }

    return allDrgs, nil
}

// Obtem um DRG pelo OCID
func GetDrgAttachment(drgAttachmentOCID string, region ...string) (core.DrgAttachment, error) {
    var drgAttachment core.DrgAttachment

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.DrgAttachment{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetDrgAttachmentRequest{
            DrgAttachmentId: common.String(drgAttachmentOCID),
        }

        resp, innerErr := client.GetDrgAttachment(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        drgAttachment = resp.DrgAttachment
        return nil
    })

    if err != nil {
        return core.DrgAttachment{}, err
    }

    return drgAttachment, nil
}

// Obtem a lista de todos os DRGS Attachment
func ListAllDrgAttachments() ([]DrgAttachment, error) {
    var allDrgAttachments []DrgAttachment
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query DrgAttachment resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var drgAttachment DrgAttachment
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            log.Fatal(innerErr)
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetDrgAttachmentRequest{
                            DrgAttachmentId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetDrgAttachment(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        drgAttachment = DrgAttachment{DrgAttachment: resp.DrgAttachment, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allDrgAttachments = append(allDrgAttachments, drgAttachment)
                }
            }
        }
    }

    return allDrgAttachments, nil
}

// Obtem um DrgRouteTable pelo OCID
func GetDrgRouteTable(drgRouteTableOCID string, region ...string) (core.DrgRouteTable, error) {
    var drgRouteTable core.DrgRouteTable

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.DrgRouteTable{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetDrgRouteTableRequest{
            DrgRouteTableId: common.String(drgRouteTableOCID),
        }

        resp, innerErr := client.GetDrgRouteTable(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        drgRouteTable = resp.DrgRouteTable
        return nil
    })

    if err != nil {
        return core.DrgRouteTable{}, err
    }

    return drgRouteTable, nil
}

// Obtem a lista de todos os DrgRouteTables
func ListAllDrgRouteTables() ([]DrgRouteTable, error) {
    var allDrgRouteTables []DrgRouteTable
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query DrgRouteTable resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var drgRouteTable DrgRouteTable
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            log.Fatal(innerErr)
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetDrgRouteTableRequest{
                            DrgRouteTableId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetDrgRouteTable(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        drgRouteTable = DrgRouteTable{DrgRouteTable: resp.DrgRouteTable, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allDrgRouteTables = append(allDrgRouteTables, drgRouteTable)
                }
            }
        }
    }

    return allDrgRouteTables, nil
}

// Obtem um DrgRouteDistribution pelo OCID
func GetDrgRouteDistribution(drgRouteDistributionOCID string, region ...string) (core.DrgRouteDistribution, error) {
    var drgRouteDistribution core.DrgRouteDistribution

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.DrgRouteDistribution{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetDrgRouteDistributionRequest{
            DrgRouteDistributionId: common.String(drgRouteDistributionOCID),
        }

        resp, innerErr := client.GetDrgRouteDistribution(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        drgRouteDistribution = resp.DrgRouteDistribution
        return nil
    })

    if err != nil {
        return core.DrgRouteDistribution{}, err
    }

    return drgRouteDistribution, nil
}

// Obtem a lista de todos os DrgRouteDistributions
func ListAllDrgRouteDistributions() ([]DrgRouteDistribution, error) {
    var allDrgRouteDistributions []DrgRouteDistribution
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query DrgRouteDistribution resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var drgRouteDistribution DrgRouteDistribution
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            log.Fatal(innerErr)
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetDrgRouteDistributionRequest{
                            DrgRouteDistributionId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetDrgRouteDistribution(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        drgRouteDistribution = DrgRouteDistribution{DrgRouteDistribution: resp.DrgRouteDistribution, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allDrgRouteDistributions = append(allDrgRouteDistributions, drgRouteDistribution)
                }
            }
        }
    }

    return allDrgRouteDistributions, nil
}

// Obtem um LocalPeeringGateway pelo OCID
func GetLocalPeeringGateway(localPeeringGatewayOCID string, region ...string) (core.LocalPeeringGateway, error) {
    var localPeeringGateway core.LocalPeeringGateway

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.LocalPeeringGateway{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetLocalPeeringGatewayRequest{
            LocalPeeringGatewayId: common.String(localPeeringGatewayOCID),
        }

        resp, innerErr := client.GetLocalPeeringGateway(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        localPeeringGateway = resp.LocalPeeringGateway
        return nil
    })

    if err != nil {
        return core.LocalPeeringGateway{}, err
    }

    return localPeeringGateway, nil
}

// Obtem a lista de todos os LocalPeeringGateways
func ListAllLocalPeeringGateways() ([]LocalPeeringGateway, error) {
    var allLocalPeeringGateways []LocalPeeringGateway
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query LocalPeeringGateway resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var localPeeringGateway LocalPeeringGateway
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            log.Fatal(innerErr)
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetLocalPeeringGatewayRequest{
                            LocalPeeringGatewayId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetLocalPeeringGateway(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        localPeeringGateway = LocalPeeringGateway{LocalPeeringGateway: resp.LocalPeeringGateway, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allLocalPeeringGateways = append(allLocalPeeringGateways, localPeeringGateway)
                }
            }
        }
    }

    return allLocalPeeringGateways, nil
}

// Obtem um DhcpDnsOption pelo OCID
func GetDhcpDnsOption(dhcpOptionsOCID string, region ...string) (core.DhcpOptions, error) {
    var dhcpDnsOption core.DhcpOptions

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.DhcpOptions{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetDhcpOptionsRequest{
            DhcpId: common.String(dhcpOptionsOCID),
        }

        resp, innerErr := client.GetDhcpOptions(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        dhcpDnsOption = resp.DhcpOptions
        return nil
    })

    if err != nil {
        return core.DhcpOptions{}, err
    }

    return dhcpDnsOption, nil
}

// ListAllDhcpDnsOptions lista todos os DhcpDnsOptions nas regiões disponíveis
func ListAllDhcpDnsOptions() ([]DhcpDnsOption, error) {
    var allDhcpDnsOptions []DhcpDnsOption
    regions, err := GetAllRegions() // Certifique-se de ter uma função que retorna todas as regiões
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query DhcpOptions resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var dhcpDnsOption DhcpDnsOption
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }
            
                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()
            
                        req := core.GetDhcpOptionsRequest{
                            DhcpId: summary.Identifier,
                        }
            
                        resp, innerErr := networkingClient.GetDhcpOptions(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }
            
                        dhcpDnsOption = DhcpDnsOption{DhcpOptions: resp.DhcpOptions, Region: region.RegionName}
                        return nil
                    })
            
                    if err != nil {
                        continue
                    }
            
                    allDhcpDnsOptions = append(allDhcpDnsOptions, dhcpDnsOption)
                }
            }
            
        }
    }

    return allDhcpDnsOptions, nil
}

// Obtem um NetworkSecurityGroup pelo OCID
func GetNetworkSecurityGroup(networkSecurityGroupOCID string, region ...string) (core.NetworkSecurityGroup, error) {
    var networkSecurityGroup core.NetworkSecurityGroup

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return core.NetworkSecurityGroup{}, err
    }

    if len(region) > 0 {
        client.SetRegion(region[0])
    }

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetNetworkSecurityGroupRequest{
            NetworkSecurityGroupId: common.String(networkSecurityGroupOCID),
        }

        resp, innerErr := client.GetNetworkSecurityGroup(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        networkSecurityGroup = resp.NetworkSecurityGroup
        return nil
    })

    if err != nil {
        return core.NetworkSecurityGroup{}, err
    }

    return networkSecurityGroup, nil
}

// Obtem a lista de todos os NetworkSecurityGroups
func ListAllNetworkSecurityGroups() ([]NetworkSecurityGroup, error) {
    var allNetworkSecurityGroups []NetworkSecurityGroup
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query NetworkSecurityGroup resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var networkSecurityGroupWithRegion NetworkSecurityGroup
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            log.Fatal(innerErr)
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetNetworkSecurityGroupRequest{
                            NetworkSecurityGroupId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetNetworkSecurityGroup(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        networkSecurityGroupWithRegion = NetworkSecurityGroup{
                            NetworkSecurityGroup: resp.NetworkSecurityGroup,
                            Region:               region.RegionName,
                        }
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allNetworkSecurityGroups = append(allNetworkSecurityGroups, networkSecurityGroupWithRegion)
                }
            }
        }
    }

    return allNetworkSecurityGroups, nil
}

// GetIPSecConnection obtém os detalhes de uma conexão IPSec específica com a região, com tentativas de reconexão
func GetIPSecConnection(ipSecConnectionId, regionName string) (*IpSecConnection, error) {
    var ipSecConnection *IpSecConnection

    client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
    if err != nil {
        return nil, err
    }

    client.SetRegion(regionName)

    err = RetryWithBackoff(5, func() error {
        ctx := context.Background()
        req := core.GetIPSecConnectionRequest{
            OpcRequestId: common.String(ipSecConnectionId),
        }

        resp, innerErr := client.GetIPSecConnection(ctx, req)
        if innerErr != nil {
            return innerErr
        }

        ipSecConnection = &IpSecConnection{
            IpSecConnection: resp.IpSecConnection,
            Region:          &regionName,
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    return ipSecConnection, nil
}

// ListAllIPSecConnection lista todas as conexões IPSec disponíveis com as regiões
func ListAllIPSecConnection() ([]IpSecConnection, error) {
    var allIPSecConnections []IpSecConnection
    regions, err := GetAllRegions()
    if err != nil {
        return nil, err
    }

    for _, region := range regions {
        if region.RegionName != nil {
            resourceSummaries, err := ResourceSearch("query IPSecConnection resources", *region.RegionName)
            if err != nil {
                continue
            }

            for _, summary := range resourceSummaries {
                if summary.Identifier != nil {
                    var ipSecConnection IpSecConnection
                    err := RetryWithBackoff(5, func() error {
                        networkingClient, innerErr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
                        if innerErr != nil {
                            return innerErr
                        }

                        networkingClient.SetRegion(*region.RegionName)
                        ctx := context.Background()

                        req := core.GetIPSecConnectionRequest{
                            IpscId: summary.Identifier,
                        }

                        resp, innerErr := networkingClient.GetIPSecConnection(ctx, req)
                        if innerErr != nil {
                            return innerErr
                        }

                        ipSecConnection = IpSecConnection{IpSecConnection: resp.IpSecConnection, Region: region.RegionName}
                        return nil
                    })

                    if err != nil {
                        continue
                    }

                    allIPSecConnections = append(allIPSecConnections, ipSecConnection)
                }
            }
        }
    }

    return allIPSecConnections, nil
}
