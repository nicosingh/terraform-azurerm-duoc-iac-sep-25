# Archivo principal de recursos del modulo
# Contiene la definicion de todos los recursos de Azure que se van a crear

# Modulo para crear Red Virtual usando Azure Verified Module (AVM)
# Crea una red virtual con las subredes especificadas
module "red_virtual" {
  source              = "git::https://github.com/Azure/terraform-azurerm-avm-res-network-virtualnetwork.git?ref=92d91187f566fc47313e1d54cda366a5acd3be55"

  address_space       = [var.bloque_red]
  location            = var.ubicacion
  name                = "${local.prefijo_recursos}-vnet"
  resource_group_name = var.nombre_grupo_recursos
  subnets             = var.subredes
}

# Modulo para crear Cuenta de Almacenamiento usando Azure Verified Module (AVM)
# Solo se crea si la variable crear_cuenta_almacenamiento es true
module "cuenta_almacenamiento" {
  count = var.crear_cuenta_almacenamiento ? 1 : 0

  source = "git::https://github.com/Azure/terraform-azurerm-avm-res-storage-storageaccount.git?ref=9d977b5d1a5412a2b79113cfdbcac457c8b5858c"

  location            = var.ubicacion
  name                = local.nombre_cuenta_almacenamiento
  resource_group_name = var.nombre_grupo_recursos

  # Habilita el acceso mediante clave compartida
  shared_access_key_enabled = true

  # Crea un contenedor dentro de la cuenta de almacenamiento
  containers = {
    (local.nombre_cuenta_almacenamiento) = {
      name = local.nombre_cuenta_almacenamiento
    }
  }
}

# Modulo para crear Maquina Virtual usando Azure Verified Module (AVM)
# Solo se crea si la variable crear_maquina_virtual es true
module "maquina_virtual" {
  count = var.crear_maquina_virtual ? 1 : 0

  source = "git::https://github.com/Azure/terraform-azurerm-avm-res-compute-virtualmachine.git?ref=c47eeb60116a6bd7a4073f96d6239f355e661f8e"

  name                       = "${local.prefijo_recursos}-vm"
  resource_group_name        = var.nombre_grupo_recursos
  location                   = var.ubicacion
  zone                       = 1
  encryption_at_host_enabled = false

  # Configuracion del sistema operativo
  os_type  = "Linux"
  sku_size = "Standard_DS1_v2"

  # Imagen de Ubuntu Server 22.04 LTS
  source_image_reference = {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts"
    version   = "latest"
  }

  # Configuracion de la interfaz de red
  # Se conecta a la primera subred de la red virtual
  network_interfaces = {
    interfaz_red_1 = {
      name = "${local.prefijo_recursos}-nic"
      ip_configurations = {
        configuracion_ip_1 = {
          name                          = "${local.prefijo_recursos}-nic"
          private_ip_subnet_resource_id = module.red_virtual.subnets["subred1"].resource_id
        }
      }
    }
  }
}
