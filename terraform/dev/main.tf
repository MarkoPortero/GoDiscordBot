module "acr_sa" {
  source = "../modules/azure/container_registry"

  resource_group_name = var.resource_group_name
  acr_name = var.acr_name
  acr_email = var.acr_email
}
