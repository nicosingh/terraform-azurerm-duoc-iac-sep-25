# Archivo de valores locales calculados
# Define valores que se reutilizan en multiples recursos del modulo

locals {
  # Prefijo comun para todos los recursos basado en proyecto y ambiente
  prefijo_recursos = lower("${var.nombre_proyecto}-${var.ambiente}")

  # Nombre de cuenta de almacenamiento que cumple con las restricciones de Azure
  # Solo letras minusculas y numeros, maximo 24 caracteres
  nombre_cuenta_almacenamiento = substr(replace(lower("${local.prefijo_recursos}-storage"), "/[^a-z0-9]/", ""), 0, 24)
}
