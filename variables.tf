# Archivo de definicion de variables de entrada del modulo
# Contiene todas las variables que pueden ser configuradas al usar este modulo

variable "ambiente" {
  type        = string
  description = "Nombre del ambiente (dev, staging, prod)"
  default     = "dev"
}

variable "bloque_red" {
  type        = string
  description = "El bloque de red para la red virtual (ej., 10.0.0.0/16)"
  default     = "10.0.0.0/16"
}

variable "crear_cuenta_almacenamiento" {
  type        = bool
  description = "Si crear o no una nueva cuenta de almacenamiento"
  default     = true
}

variable "crear_maquina_virtual" {
  type        = bool
  description = "Si crear o no una nueva máquina virtual"
  default     = true
}

variable "nombre_grupo_recursos" {
  type        = string
  description = "Nombre del grupo de recursos nuevo o existente a utilizar"
  default     = ""
}

variable "nombre_proyecto" {
  type        = string
  description = "El nombre del proyecto, utilizado para nombrar recursos"
}

variable "subredes" {
  type = map(object({
    name             = string
    address_prefixes = list(string)
  }))
  description = "Configuración de subredes para la red virtual"
  default = {
    "subred1" = {
      name             = "subred1"
      address_prefixes = ["10.0.0.0/24"]
    }
    "subred2" = {
      name             = "subred2"
      address_prefixes = ["10.0.1.0/24"]
    }
  }
}

variable "ubicacion" {
  type        = string
  description = "Región de Azure donde se crearán los recursos"
}
