# Archivo de outputs del modulo
# Define los valores que el modulo devuelve para uso en otros modulos o configuraciones

# Informacion de la Red Virtual creada
output "nombre_red_virtual" {
  value       = module.red_virtual.name
  description = "Nombre de la Red Virtual"
}

output "id_red_virtual" {
  value       = module.red_virtual.resource_id
  description = "ID de la Red Virtual"
}

# Informacion de la Cuenta de Almacenamiento
# Utiliza try() para manejar el caso donde no se crea la cuenta de almacenamiento
output "nombre_cuenta_almacenamiento" {
  value       = try(module.cuenta_almacenamiento[0].name, null)
  description = "Nombre de la Cuenta de Almacenamiento"
}

# Informacion de la Maquina Virtual
# Utiliza try() para manejar el caso donde no se crea la maquina virtual
output "ip_privada_maquina_virtual" {
  value       = try(module.maquina_virtual[0].private_ip_address, null)
  description = "Dirección IP privada de la Máquina Virtual"
}

output "dns_maquina_virtual" {
  value       = try(module.maquina_virtual[0].dns_name, null)
  description = "DNS de la Máquina Virtual"
}
