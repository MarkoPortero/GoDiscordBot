resource "local_file" "dockerconfig" {
  depends_on = [module.acr_sa]
  filename   = "./generated/dockerconfig_gobot"
  content    = module.acr_sa.dockerconfigjson
}