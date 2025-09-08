# Archivo de versionado y configuracion de proveedores
# Define las versiones minimas requeridas de Terraform y sus proveedores
terraform {
  required_version = "~> 1.13.0"

  # Proveedores requeridos para este modulo
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.41.0"
    }
  }
}
