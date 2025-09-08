# terraform-azurerm-duoc-iac-sep-25

Este es un módulo de Terraform para Azure diseñado para crear una infraestructura básica y completa en Azure. El proyecto utiliza Azure Verified Modules (AVM) para garantizar las mejores prácticas y estándares de la industria.

Características principales:
- Infraestructura modular: Implementa una arquitectura de red completa con máquina virtual y almacenamiento
- Flexibilidad: Permite activar/desactivar componentes mediante variables booleanas
- Nomenclatura consistente: Utiliza un sistema de prefijos basado en proyecto y ambiente
- Compatibilidad: Diseñado para trabajar con Terraform ~> 1.13.0 y Azure Provider ~> 4.41.0

Componentes que despliega:
- Red Virtual (VNet) con subredes configurables
- Cuenta de Almacenamiento con contenedor incluido (opcional)
- Máquina Virtual Linux (Ubuntu 22.04 LTS) conectada a la red (opcional)

Casos de uso: Ideal para entornos de desarrollo, staging o producción que requieren una infraestructura básica pero completa en Azure, con la flexibilidad de activar solo los componentes necesarios según el caso de uso específico.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.13.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 4.41.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_cuenta_almacenamiento"></a> [cuenta\_almacenamiento](#module\_cuenta\_almacenamiento) | Azure/avm-res-storage-storageaccount/azurerm | 0.6.4 |
| <a name="module_maquina_virtual"></a> [maquina\_virtual](#module\_maquina\_virtual) | Azure/avm-res-compute-virtualmachine/azurerm | 0.18.1 |
| <a name="module_red_virtual"></a> [red\_virtual](#module\_red\_virtual) | Azure/avm-res-network-virtualnetwork/azurerm | 0.10.0 |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ambiente"></a> [ambiente](#input\_ambiente) | Nombre del ambiente (dev, staging, prod) | `string` | `"dev"` | no |
| <a name="input_bloque_red"></a> [bloque\_red](#input\_bloque\_red) | El bloque de red para la red virtual (ej., 10.0.0.0/16) | `string` | `"10.0.0.0/16"` | no |
| <a name="input_crear_cuenta_almacenamiento"></a> [crear\_cuenta\_almacenamiento](#input\_crear\_cuenta\_almacenamiento) | Si crear o no una nueva cuenta de almacenamiento | `bool` | `true` | no |
| <a name="input_crear_maquina_virtual"></a> [crear\_maquina\_virtual](#input\_crear\_maquina\_virtual) | Si crear o no una nueva máquina virtual | `bool` | `true` | no |
| <a name="input_nombre_grupo_recursos"></a> [nombre\_grupo\_recursos](#input\_nombre\_grupo\_recursos) | Nombre del grupo de recursos nuevo o existente a utilizar | `string` | `""` | no |
| <a name="input_nombre_proyecto"></a> [nombre\_proyecto](#input\_nombre\_proyecto) | El nombre del proyecto, utilizado para nombrar recursos | `string` | n/a | yes |
| <a name="input_subredes"></a> [subredes](#input\_subredes) | Configuración de subredes para la red virtual | <pre>map(object({<br/>    name             = string<br/>    address_prefixes = list(string)<br/>  }))</pre> | <pre>{<br/>  "subred1": {<br/>    "address_prefixes": [<br/>      "10.0.0.0/24"<br/>    ],<br/>    "name": "subred1"<br/>  },<br/>  "subred2": {<br/>    "address_prefixes": [<br/>      "10.0.1.0/24"<br/>    ],<br/>    "name": "subred2"<br/>  }<br/>}</pre> | no |
| <a name="input_ubicacion"></a> [ubicacion](#input\_ubicacion) | Región de Azure donde se crearán los recursos | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dns_maquina_virtual"></a> [dns\_maquina\_virtual](#output\_dns\_maquina\_virtual) | DNS de la Máquina Virtual |
| <a name="output_id_red_virtual"></a> [id\_red\_virtual](#output\_id\_red\_virtual) | ID de la Red Virtual |
| <a name="output_ip_privada_maquina_virtual"></a> [ip\_privada\_maquina\_virtual](#output\_ip\_privada\_maquina\_virtual) | Dirección IP privada de la Máquina Virtual |
| <a name="output_nombre_cuenta_almacenamiento"></a> [nombre\_cuenta\_almacenamiento](#output\_nombre\_cuenta\_almacenamiento) | Nombre de la Cuenta de Almacenamiento |
| <a name="output_nombre_red_virtual"></a> [nombre\_red\_virtual](#output\_nombre\_red\_virtual) | Nombre de la Red Virtual |
<!-- END_TF_DOCS -->
