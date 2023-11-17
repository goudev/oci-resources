# OCI SDK Go Module

Este módulo fornece funcionalidades para interagir com o Oracle Cloud Infrastructure (OCI) usando Go.

## Instalação

Para usar este módulo em seu projeto Go:

```bash
go get github.com/goudev/oci-resource
```

## Uso

Aqui está uma visão geral das principais funcionalidades e como utilizá-las em seu projeto.

### resourcesearch.go

- `ResourceSearch(query string, region ...string) ([]resourcesearch.ResourceSummary, error)`: Descrição breve da função `ResourceSearch`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.ResourceSearch(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### instance.go

- `GetInstance(instanceOcid string, region ...string) (core.Instance, error)`: Descrição breve da função `GetInstance`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetInstance(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetAllInstances() ([]core.Instance, error)`: Descrição breve da função `GetAllInstances`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.GetAllInstances(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### blockstorage.go

- `GetBootVolume(bootVolumeOCID string, region ...string) (BootVolume, error)`: Descrição breve da função `GetBootVolume`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetBootVolume(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllBootVolumes() ([]BootVolume, error)`: Descrição breve da função `ListAllBootVolumes`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllBootVolumes(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetBootVolumeAttachment(attachmentOCID string) (core.BootVolumeAttachment, error)`: Descrição breve da função `GetBootVolumeAttachment`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetBootVolumeAttachment(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllBootVolumeAttachments() ([]core.BootVolumeAttachment, error)`: Descrição breve da função `ListAllBootVolumeAttachments`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllBootVolumeAttachments(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetBlockVolume(blockVolumeOCID string, region ...string) (Volume, error)`: Descrição breve da função `GetBlockVolume`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetBlockVolume(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllBlockVolumeAttachments() ([]core.VolumeAttachment, error)`: Descrição breve da função `ListAllBlockVolumeAttachments`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllBlockVolumeAttachments(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetBlockVolumeAttachment(volumeAttachmentOCID string) (core.VolumeAttachment, error)`: Descrição breve da função `GetBlockVolumeAttachment`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetBlockVolumeAttachment(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllBlockVolumes() ([]Volume, error)`: Descrição breve da função `ListAllBlockVolumes`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllBlockVolumes(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### database.go

- `ListAllDatabases() ([]database.Database, error)`: Descrição breve da função `ListAllDatabases`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDatabases(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllCloudExadataInfrastructures() ([]database.CloudExadataInfrastructure, error)`: Descrição breve da função `ListAllCloudExadataInfrastructures`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllCloudExadataInfrastructures(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllDbNodes() ([]database.DbNode, error)`: Descrição breve da função `ListAllDbNodes`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDbNodes(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllCloudVmClusters() ([]database.CloudVmCluster, error)`: Descrição breve da função `ListAllCloudVmClusters`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllCloudVmClusters(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### region.go

- `GetAllRegions() ([]identity.RegionSubscription, error)`: Descrição breve da função `GetAllRegions`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.GetAllRegions(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### network.go

- `GetVCN(vcnOCID string, region ...string) (core.Vcn, error)`: Descrição breve da função `GetVCN`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetVCN(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllVCNs() ([]Vcn, error)`: Descrição breve da função `ListAllVCNs`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllVCNs(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetSubnet(subnetOCID string, region ...string) (core.Subnet, error)`: Descrição breve da função `GetSubnet`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetSubnet(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllSubnets() ([]Subnet, error)`: Descrição breve da função `ListAllSubnets`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllSubnets(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetNatGateway(natGatewayOCID string, region ...string) (core.NatGateway, error)`: Descrição breve da função `GetNatGateway`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetNatGateway(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllNatGateways() ([]NatGateway, error)`: Descrição breve da função `ListAllNatGateways`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllNatGateways(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetServiceGateway(serviceGatewayOCID string, region ...string) (core.ServiceGateway, error)`: Descrição breve da função `GetServiceGateway`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetServiceGateway(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllServiceGateways() ([]ServiceGateway, error)`: Descrição breve da função `ListAllServiceGateways`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllServiceGateways(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetRouteTable(routeTableOCID string, region ...string) (core.RouteTable, error)`: Descrição breve da função `GetRouteTable`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetRouteTable(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllRouteTables() ([]RouteTable, error)`: Descrição breve da função `ListAllRouteTables`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllRouteTables(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetSecurityList(securityListOCID string, region ...string) (core.SecurityList, error)`: Descrição breve da função `GetSecurityList`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetSecurityList(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllSecurityLists() ([]SecurityList, error)`: Descrição breve da função `ListAllSecurityLists`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllSecurityLists(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetInternetGateway(internetGatewayOCID string, region ...string) (core.InternetGateway, error)`: Descrição breve da função `GetInternetGateway`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetInternetGateway(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllInternetGateways() ([]InternetGateway, error)`: Descrição breve da função `ListAllInternetGateways`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllInternetGateways(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetDrg(drgOCID string, region ...string) (core.Drg, error)`: Descrição breve da função `GetDrg`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetDrg(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllDrgs() ([]Drg, error)`: Descrição breve da função `ListAllDrgs`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDrgs(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetDrgAttachment(drgAttachmentOCID string, region ...string) (core.DrgAttachment, error)`: Descrição breve da função `GetDrgAttachment`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetDrgAttachment(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllDrgAttachments() ([]DrgAttachment, error)`: Descrição breve da função `ListAllDrgAttachments`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDrgAttachments(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetDrgRouteTable(drgRouteTableOCID string, region ...string) (core.DrgRouteTable, error)`: Descrição breve da função `GetDrgRouteTable`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetDrgRouteTable(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllDrgRouteTables() ([]DrgRouteTable, error)`: Descrição breve da função `ListAllDrgRouteTables`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDrgRouteTables(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetDrgRouteDistribution(drgRouteDistributionOCID string, region ...string) (core.DrgRouteDistribution, error)`: Descrição breve da função `GetDrgRouteDistribution`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetDrgRouteDistribution(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllDrgRouteDistributions() ([]DrgRouteDistribution, error)`: Descrição breve da função `ListAllDrgRouteDistributions`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDrgRouteDistributions(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetLocalPeeringGateway(localPeeringGatewayOCID string, region ...string) (core.LocalPeeringGateway, error)`: Descrição breve da função `GetLocalPeeringGateway`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetLocalPeeringGateway(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllLocalPeeringGateways() ([]LocalPeeringGateway, error)`: Descrição breve da função `ListAllLocalPeeringGateways`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllLocalPeeringGateways(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetDhcpDnsOption(dhcpOptionsOCID string, region ...string) (core.DhcpOptions, error)`: Descrição breve da função `GetDhcpDnsOption`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetDhcpDnsOption(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllDhcpDnsOptions() ([]DhcpDnsOption, error)`: Descrição breve da função `ListAllDhcpDnsOptions`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllDhcpDnsOptions(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetNetworkSecurityGroup(networkSecurityGroupOCID string, region ...string) (core.NetworkSecurityGroup, error)`: Descrição breve da função `GetNetworkSecurityGroup`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetNetworkSecurityGroup(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllNetworkSecurityGroups() ([]NetworkSecurityGroup, error)`: Descrição breve da função `ListAllNetworkSecurityGroups`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllNetworkSecurityGroups(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `GetIPSecConnection(ipSecConnectionId, regionName string) (*IpSecConnection, error)`: Descrição breve da função `GetIPSecConnection`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetIPSecConnection(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllIPSecConnection() ([]IpSecConnection, error)`: Descrição breve da função `ListAllIPSecConnection`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllIPSecConnection(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### identity.go

- `ListAllCompartments() ([]identity.Compartment, error)`: Descrição breve da função `ListAllCompartments`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllCompartments(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

### analytics.go

- `GetAnalyticsInstance(analyticsInstanceOCID string, region ...string) (analytics.AnalyticsInstance, error)`: Descrição breve da função `GetAnalyticsInstance`.

  ```go
  // Exemplo de uso
  resultado, err := pkg.GetAnalyticsInstance(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

- `ListAllAnalyticsInstances() ([]AnalyticsInstance, error)`: Descrição breve da função `ListAllAnalyticsInstances`.
  ```go
  // Exemplo de uso
  resultado, err := pkg.ListAllAnalyticsInstances(args)
  if err != nil {
    // Handle error
  }
  fmt.Println(resultado)
  ```

## Contribuição

Contribuições para este projeto são bem-vindas. Por favor, envie pull requests para melhorias e correções.

## Licença

[Inserir informações da licença aqui]
