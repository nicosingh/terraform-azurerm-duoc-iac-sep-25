# Registro de Cambios

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
y este proyecto se adhiere al [versionado semántico](https://semver.org/spec/v2.0.0.html).

## [1.2.1] - 2025-09-10

### Corregido
- **Configuración de proveedores del módulo**: Eliminación del archivo `providers.tf` innecesario que causaba conflictos en la estructura del módulo
- **Mejoras en pruebas automatizadas**: Refinamiento del test `TestTerraformPlanValidVars` con configuración temporal específica del proveedor Azure
  - Implementación de archivo provider temporal que se crea y elimina automáticamente durante la ejecución del test
  - Mejora en la detección de errores de autenticación vs errores de configuración de variables
  - Adición de import `os` faltante en el archivo de pruebas

### Técnico
- **Optimización de estructura del módulo**: Eliminación de configuración de provider redundante para permitir mayor flexibilidad en el uso del módulo
- **Robustez en testing**: El test de planificación ahora maneja correctamente la configuración del proveedor Azure de forma aislada
- **Limpieza automática**: Los archivos temporales creados durante las pruebas se eliminan automáticamente

## [1.2.0] - 2025-09-09

### Agregado
- **Suite completa de pruebas automatizadas**: Implementación de Terratest para validación integral del módulo
  - **Pruebas de inicialización**: Validación de `terraform init` y configuración básica
  - **Pruebas de validación de variables**: Verificación de todas las reglas de validación de inputs
  - **Pruebas de sintaxis**: Validación de sintaxis de Terraform usando `terraform validate`
  - **Pruebas de estructura**: Verificación de existencia de archivos esenciales del módulo
  - **Prueba de planificación**: Validación de `terraform plan` con variables válidas
- **Gestión completa de dependencias Go**: Configuración de módulos Go con `go.mod` y `go.sum`
- **Mejoras en CI/CD workflows**: Integración de Go y Terratest en pipelines automatizados
  - Configuración de caché de dependencias Go para mejor rendimiento
  - Integración de pruebas automatizadas en proceso de validación

### Cambiado
- **Migración completa de testing**: Transición de `terraform_basic_test.go` a suite completa en `terraform_module_test.go`
- **Mejoras en validación de errores**: Refinamiento de pruebas de autenticación Azure con múltiples escenarios
- **Actualización de configuración de pre-commit**: Inclusión de formateo automático para archivos Go
- **Optimizaciones en workflows**: Mejoras de rendimiento con cache de dependencias y configuración refinada

### Corregido
- **Configuración de Checkov**: Eliminación de parámetro `output_file_path` deprecado en workflow de validaciones
- **Manejo robusto de errores de autenticación Azure**: Implementación de validación mejorada para diferentes tipos de errores de autenticación en entornos de CI/CD

### Características Técnicas
- **Testing robusto**: 8 pruebas automatizadas cubriendo todos los aspectos críticos del módulo
- **Validación integral**: Pruebas de variables, sintaxis, estructura y planificación
- **Gestión de errores avanzada**: Detección inteligente de errores de autenticación vs. errores de configuración
- **Compatibilidad mejorada**: Soporte para múltiples escenarios de despliegue y entornos
- **Integración continua completa**: Pruebas automatizadas integradas en todos los workflows

## [1.1.1] - 2025-09-09

### Agregado
- **Pipeline de CI/CD completo**: Implementación de workflows de GitHub Actions para garantizar calidad del código
  - **Workflow de Validaciones** (`validaciones.yml`): Ejecuta terraform fmt, TFLint, terraform validate y análisis de seguridad con Checkov
  - **Workflow de Auto-fixes** (`auto-fixes.yml`): Formateo automático de código y actualización de documentación con terraform-docs
- **Sección de Integración Continua en README**: Documentación detallada de los pipelines implementados
- **Guía de Contribuciones**: Proceso completo para contribuir al proyecto con pasos claros y mejores prácticas

### Cambiado
- **Título del módulo**: Actualizado de "terraform-azurerm-duoc-iac-sep-25" a "Módulo Terraform: Infraestructura Base Azure con AVM" para mayor claridad
- **Mejora en pipelines de CI/CD**: Múltiples refinamientos en configuración de workflows
  - Actualización de Checkov a versión v12.1347.0
  - Fijación de TFLint a versión v0.59.1 para consistencia
  - Configuración mejorada de permisos y commits automáticos
  - Triggers consistentes en todas las ramas
- **Documentación actualizada automáticamente**: README.md mantenido al día mediante terraform-docs automatizado

### Características Técnicas
- **Automatización completa de calidad**: Todos los cambios pasan por validaciones automáticas de formato, sintaxis, seguridad y documentación
- **Commits automatizados**: Los pipelines pueden commitear automáticamente correcciones de formato y actualizaciones de documentación
- **Análisis de seguridad integrado**: Checkov integrado para detectar vulnerabilidades y problemas de configuración
- **Documentación siempre actualizada**: terraform-docs mantiene la documentación técnica sincronizada automáticamente

## [1.1.0] - 2025-09-08

### Agregado
- **Sistema de Pre-commit**: Integración completa de pre-commit hooks para mejorar la calidad del código
  - Hook de terraform-fmt para formateo automático de código Terraform
  - Hook de terraform-docs para documentación automática del README.md
- **Configuración de Pre-commit**: Archivo `.pre-commit-config.yaml` con hooks configurados para mantener estándares de código

### Corregido
- **Sintaxis de contenedor de almacenamiento**: Corrección en la sintaxis para el nombre del contenedor en el módulo de cuenta de almacenamiento

### Cambiado
- **Actualización de hooks de pre-commit**: Mejoras en la configuración de hooks de pre-commit y fuentes de módulos en configuración Terraform
- **Documentación actualizada**: README.md actualizado automáticamente mediante terraform-docs para reflejar cambios en la configuración
- **Configuración de .gitignore**: Actualizaciones menores en el archivo .gitignore

## [1.0.1] - 2025-09-08

### Cambiado
- **Estandarización de comentarios**: Actualizados todos los comentarios y descripciones en `variables.tf` y `outputs.tf` para incluir punto final, manteniendo consistencia con los mensajes de error de validación
- **Mejora en la documentación**: Normalización del formato de comentarios tanto en comentarios de archivo como en descripciones de variables y outputs

## [1.0.0] - 2025-09-08

### Agregado
- **Lanzamiento inicial** del módulo de Terraform para Infraestructura en Azure
- **Módulo de Red Virtual** utilizando Azure Verified Module (AVM) `Azure/avm-res-network-virtualnetwork/azurerm` v0.10.0
  - Espacio de direcciones y subredes configurables
  - Configuración por defecto con dos subredes (10.0.0.0/24 y 10.0.1.0/24)
- **Módulo de Cuenta de Almacenamiento** utilizando Azure Verified Module (AVM) `Azure/avm-res-storage-storageaccount/azurerm` v0.6.4
  - Creación opcional mediante la variable `crear_cuenta_almacenamiento`
  - Incluye creación de contenedor con clave de acceso compartido habilitada
- **Módulo de Máquina Virtual** utilizando Azure Verified Module (AVM) `Azure/avm-res-compute-virtualmachine/azurerm` v0.18.1
  - Creación opcional mediante la variable `crear_maquina_virtual`
  - Ubuntu Server 22.04 LTS con SKU Standard_DS1_v2
  - Interfaz de red conectada a la primera subred
- **Sistema integral de variables** con valores por defecto sensatos:
  - Configuración de ambiente (`ambiente`)
  - Configuración de bloque de red (`bloque_red`)
  - Nomenclatura de grupo de recursos (`nombre_grupo_recursos`)
  - Nomenclatura de proyecto para prefijos de recursos (`nombre_proyecto`)
  - Configuración de región de Azure (`ubicacion`)
  - Configuración flexible de subredes (`subredes`)
- **Valores de salida** para integración con otros módulos:
  - ID y nombre de Red Virtual
  - Nombre de Cuenta de Almacenamiento
  - IP privada y DNS de Máquina Virtual
- **Valores locales** para nomenclatura consistente de recursos con prefijos de proyecto y ambiente
- **Configuración de proveedor** con Azure Resource Manager provider v4.41.0
- **Restricciones de versión de Terraform** requiriendo Terraform ~> 1.13.0
- **Documentación integral** incluyendo:
  - Descripción del módulo y casos de uso
  - Tablas de requisitos, entradas y salidas
  - Resumen de arquitectura de componentes

### Características
- **Arquitectura modular**: Cada componente (VNet, Storage, VM) puede ser habilitado/deshabilitado independientemente
- **Convención de nomenclatura consistente**: Todos los recursos siguen el patrón `{proyecto}-{ambiente}-{recurso}`
- **Listo para producción**: Construido con Azure Verified Modules para cumplimiento de mejores prácticas
- **Redes flexibles**: Subredes y espacios de direcciones configurables
- **Soporte multi-ambiente**: Diseñado para despliegues de desarrollo, staging y producción

### Componentes de Infraestructura
- Red Virtual con subredes configurables
- Cuenta de Almacenamiento con contenedor (opcional)
- Máquina Virtual Linux con conectividad de red (opcional)
- Estrategia consistente de nomenclatura y etiquetado de recursos

[1.2.1]: https://github.com/nicosingh/terraform-azurerm-duoc-iac-sep-25/releases/tag/v1.2.1
[1.2.0]: https://github.com/nicosingh/terraform-azurerm-duoc-iac-sep-25/releases/tag/v1.2.0
[1.1.1]: https://github.com/nicosingh/terraform-azurerm-duoc-iac-sep-25/releases/tag/v1.1.1
[1.0.2]: https://github.com/nicosingh/terraform-azurerm-duoc-iac-sep-25/releases/tag/v1.0.2
[1.0.1]: https://github.com/nicosingh/terraform-azurerm-duoc-iac-sep-25/releases/tag/v1.0.1
[1.0.0]: https://github.com/nicosingh/terraform-azurerm-duoc-iac-sep-25/releases/tag/v1.0.0
