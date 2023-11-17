# OCI SDK Go Module

Este módulo fornece funcionalidades para interagir com o Oracle Cloud Infrastructure (OCI) usando Go.

## Instalação

Para usar este módulo em seu projeto Go:

```bash
go get github.com/goudev/oci-resources@v1.0.0
```

## Uso

Aqui está uma visão geral de como utilizar uma das funções.

### Listar todas as instâncias

```go
package main

import (
	"fmt"

	"github.com/goudev/oci-resources/oci"
)

func main(){
	fmt.Println(oci.GetAllInstances())
}
```

# Lista de funções disponíveis

## Analytics

- GetAnalyticsInstance(analyticsInstanceOCID
- ListAllAnalyticsInstances()

## Blockstorage

- GetBootVolume(bootVolumeOCID, region ...string)
- ListAllBootVolumes()
- GetBootVolumeAttachment(attachmentOCID, region ...string)
- ListAllBootVolumeAttachments()
- GetBlockVolume(blockVolumeOCID, region ...string)
- ListAllBlockVolumeAttachments()
- GetBlockVolumeAttachment(volumeAttachmentOCID, region ...string)
- ListAllBlockVolumes()

## Database

- ListAllDatabases()
- ListAllCloudExadataInfrastructures()
- ListAllDbNodes()
- ListAllCloudVmClusters()

## Identity

- ListAllCompartments()

## Instance

- GetInstance(instanceOcid, region ...string)
- ListAllInstances()

## Network

- GetVCN(vcnOCID, region ...string)
- ListAllVCNs()
- GetSubnet(subnetOCID, region ...string)
- ListAllSubnets()
- GetNatGateway(natGatewayOCID, region ...string)
- ListAllNatGateways()
- GetServiceGateway(serviceGatewayOCID, region ...string)
- ListAllServiceGateways()
- GetRouteTable(routeTableOCID, region ...string)
- ListAllRouteTables()
- GetSecurityList(securityListOCID, region ...string)
- ListAllSecurityLists()
- GetInternetGateway(internetGatewayOCID, region ...string)
- ListAllInternetGateways()
- GetDrg(drgOCID, region ...string)
- ListAllDrgs()
- GetDrgAttachment(drgAttachmentOCID, region ...string)
- ListAllDrgAttachments()
- GetDrgRouteTable(drgRouteTableOCID, region ...string)
- ListAllDrgRouteTables()
- GetDrgRouteDistribution(drgRouteDistributionOCID, region ...string)
- ListAllDrgRouteDistributions()
- GetLocalPeeringGateway(localPeeringGatewayOCID, region ...string)
- ListAllLocalPeeringGateways()
- GetDhcpDnsOption(dhcpOptionsOCID, region ...string)
- ListAllDhcpDnsOptions()
- GetNetworkSecurityGroup(networkSecurityGroupOCID, region ...string)
- ListAllNetworkSecurityGroups()
- GetIPSecConnection(ipSecConnectionId, region ...string)
- ListAllIPSecConnection()

## Region

- GetAllRegions()

## Resourcesearch

- ResourceSearch(query, region ...string)

# Contribuindo com o plugin

✨ Obrigado por contribuir com o modulo oci-resources ✨

## Como posso contribuir?

### Melhore a documentação

Você é o candidato perfeito para nos ajudar a melhorar a nossa documentação. Correções de erro de digitação, correções de erro, melhores explicações e mais exemplos, etc. Issues em aberto para as coisas que poderiam ser melhoradas.

### Melhore issues

Algumas issues são criadas sem informações suficientes, não podem ser reproduzidas, ou não são válidas. Ajude-nos a torná-las mais fáceis de serem resolvidas. Lidar com issues ocupa muito tempo que poderia ser usado para corrigir bugs e adicionar funcionalidades.

### Forneça feedback sobre issues

Estamos sempre à procura de mais opiniões em discussões no issue tracker. É uma boa oportunidade de influenciar o rumo do plugin.

### Submetendo uma issue

- Pesquise o issue tracker antes de abrir uma issue.
- Certifique-se de que você está usando a versão mais recente do plugin.
- Use um título claro e descritivo.
- Inclua o máximo possível de informações: etapas para reproduzir a issue, mensagem de erro, sistema operacional, etc.
- Quanto mais tempo você colocar em uma issue, quanto mais nós colocaremos.

### Submetendo uma pull request

- Mudanças não-triviais muitas vezes são primeiramente melhor discutidas em uma issue, para impedí-lo de fazer trabalho desnecessário.
- Novas funcionalidades devem ser acompanhadas de testes e documentação.
- Não inclua mudanças não relacionadas.
- Faça lint e teste antes de submeter a pull request.
- Use um título claro e descritivo para a pull request e os commits.
- Escreva uma descrição convincente do motivo pelo qual devemos aceitar sua pull request. É seu trabalho nos convencer.
- Podemos pedir que faça mudanças à sua pull request. Nunca há necessidade de abrir outra pull request.
