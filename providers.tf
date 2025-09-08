# Configuracion del proveedor de Azure Resource Manager
# Se configura sin registros automaticos de proveedores de recursos
provider "azurerm" {
  features {}
  resource_provider_registrations = "none"
}
