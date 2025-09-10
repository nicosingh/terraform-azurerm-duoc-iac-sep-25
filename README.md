# Módulo Terraform: Infraestructura Base Azure con AVM

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

## Integración Continua (CI/CD)

Este proyecto incluye dos pipelines de GitHub Actions que garantizan la calidad y mantenimiento automático del código:

**Pipeline de Validaciones** (`validaciones.yml`):
- Se ejecuta en cada push y pull request
- Verifica el formato del código con `terraform fmt`
- Ejecuta análisis estático con TFLint para detectar errores potenciales
- Valida la sintaxis y configuración con `terraform validate`
- Realiza análisis de seguridad y mejores prácticas con Checkov

**Pipeline de Auto-fixes** (`auto-fixes.yml`):
- Se ejecuta automáticamente en cada push y pull request
- Formatea el código Terraform automáticamente
- Genera y actualiza la documentación del módulo usando terraform-docs
- Commitea automáticamente los cambios de formato y documentación

Ambos pipelines utilizan Terraform 1.13.0 y se ejecutan en todas las ramas, asegurando consistencia y calidad en todo el ciclo de desarrollo.

## Pruebas Unitarias

Este proyecto incluye pruebas unitarias automatizadas implementadas con Terratest, una biblioteca de Go diseñada específicamente para probar código de infraestructura como código.

**Cómo ejecutar las pruebas**:

Desde el directorio del proyecto, ejecuta los siguientes comandos:

```bash
# Navegar al directorio de pruebas
cd tests

# Descargar dependencias de Go
go mod tidy

# Ejecutar todas las pruebas
go test -v

# Ejecutar una prueba específica
go test -v -run TestTerraformInit
```

**Requisitos**: Go 1.19+ y acceso a internet para descargar módulos de Terraform durante las pruebas.

Las pruebas validan la funcionalidad del módulo sin crear recursos reales en Azure, enfocándose en la validación de configuración y inicialización.

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
| <a name="module_cuenta_almacenamiento"></a> [cuenta\_almacenamiento](#module\_cuenta\_almacenamiento) | git::https://github.com/Azure/terraform-azurerm-avm-res-storage-storageaccount.git | 9d977b5d1a5412a2b79113cfdbcac457c8b5858c |
| <a name="module_maquina_virtual"></a> [maquina\_virtual](#module\_maquina\_virtual) | git::https://github.com/Azure/terraform-azurerm-avm-res-compute-virtualmachine.git | c47eeb60116a6bd7a4073f96d6239f355e661f8e |
| <a name="module_red_virtual"></a> [red\_virtual](#module\_red\_virtual) | git::https://github.com/Azure/terraform-azurerm-avm-res-network-virtualnetwork.git | 92d91187f566fc47313e1d54cda366a5acd3be55 |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ambiente"></a> [ambiente](#input\_ambiente) | Nombre del ambiente (dev, staging, prod). | `string` | `"dev"` | no |
| <a name="input_bloque_red"></a> [bloque\_red](#input\_bloque\_red) | El bloque de red para la red virtual (ej., 10.0.0.0/16). | `string` | `"10.0.0.0/16"` | no |
| <a name="input_crear_cuenta_almacenamiento"></a> [crear\_cuenta\_almacenamiento](#input\_crear\_cuenta\_almacenamiento) | Si crear o no una nueva cuenta de almacenamiento. | `bool` | `true` | no |
| <a name="input_crear_maquina_virtual"></a> [crear\_maquina\_virtual](#input\_crear\_maquina\_virtual) | Si crear o no una nueva máquina virtual. | `bool` | `true` | no |
| <a name="input_nombre_grupo_recursos"></a> [nombre\_grupo\_recursos](#input\_nombre\_grupo\_recursos) | Nombre del grupo de recursos nuevo o existente a utilizar. | `string` | `""` | no |
| <a name="input_nombre_proyecto"></a> [nombre\_proyecto](#input\_nombre\_proyecto) | El nombre del proyecto, utilizado para nombrar recursos. | `string` | n/a | yes |
| <a name="input_subredes"></a> [subredes](#input\_subredes) | Configuración de subredes para la red virtual. | <pre>map(object({<br/>    name             = string<br/>    address_prefixes = list(string)<br/>  }))</pre> | <pre>{<br/>  "subred1": {<br/>    "address_prefixes": [<br/>      "10.0.0.0/24"<br/>    ],<br/>    "name": "subred1"<br/>  },<br/>  "subred2": {<br/>    "address_prefixes": [<br/>      "10.0.1.0/24"<br/>    ],<br/>    "name": "subred2"<br/>  }<br/>}</pre> | no |
| <a name="input_ubicacion"></a> [ubicacion](#input\_ubicacion) | Región de Azure donde se crearán los recursos. | `string` | n/a | yes |

<!-- END_TF_DOCS -->

## Contribuciones

¡Las contribuciones son bienvenidas y valoradas! Este proyecto está diseñado para ser una solución colaborativa que beneficie a toda la comunidad de desarrolladores trabajando con infraestructura en Azure.

Formas de contribuir:
- **Reportar problemas**: Si encuentras errores o tienes sugerencias de mejora, abre un issue
- **Proponer características**: Comparte ideas para nuevas funcionalidades o componentes
- **Mejorar documentación**: Ayuda a hacer el proyecto más accesible con mejor documentación
- **Enviar pull requests**: Implementa correcciones, mejoras o nuevas características

Proceso de contribución:
1. Fork del repositorio
2. Crea una rama con nombre descriptivo (`feature/nueva-funcionalidad` o `fix/correccion-problema`)
3. Implementa tus cambios siguiendo las mejores prácticas de Terraform
4. Asegúrate de que los pipelines de validación pasen exitosamente
5. Envía un pull request con descripción clara de los cambios

Todos los cambios pasan por un proceso de revisión automática que incluye validaciones de formato, linting, seguridad y documentación, garantizando la calidad y consistencia del módulo.
