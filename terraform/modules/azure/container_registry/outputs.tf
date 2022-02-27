output "id" {
  value = azurerm_container_registry.acr.id
}

output "login_server" {
  value = azurerm_container_registry.acr.login_server
}

output "dockerconfigjson" {
  value = jsonencode(local.dockerconfigjson)
}
