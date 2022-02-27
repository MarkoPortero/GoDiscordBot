resource "azurerm_resource_group" "shared_rg" {
  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_container_registry" "acr" {
  name                = var.acr_name
  resource_group_name = azurerm_resource_group.shared_rg.name
  location            = var.location
  sku                 = "Standard"
  admin_enabled       = true # Turn this off if we get identity to work
}

# https://stackoverflow.com/questions/62137632/create-kubernetes-secret-for-docker-registry-terraform
# https://github.com/hashicorp/terraform-provider-kubernetes/issues/611

locals {
  dockerconfigjson = {
    "auths" : {
      "${var.acr_name}.azurecr.io" : {
        email    = var.acr_email
        username = azurerm_container_registry.acr.admin_username
        password = azurerm_container_registry.acr.admin_password
      auth = base64encode("${azurerm_container_registry.acr.admin_username}:${azurerm_container_registry.acr.admin_password}") }
    }
  }
}

# resource "kubernetes_secret" "acr_secret" {
#   depends_on = [
#     azurerm_container_registry.acr
#   ]
#   metadata {
#     name = "squadassistregistrykey"
#   }

#   data = {
#     ".dockerconfigjson" = jsonencode(locals.dockerconfigjson)
#   }

#   type = "kubernetes.io/dockerconfigjson"
# }
