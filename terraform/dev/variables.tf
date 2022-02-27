variable "resource_group_name" {
  type        = string
  description = "RG name in Azure"
}

variable "acr_name" {
  type        = string
  description = "ACR name"
}

variable "acr_email" {
  type        = string
  description = "Email to be used for the ACR?"
}
