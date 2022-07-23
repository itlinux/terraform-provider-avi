resource "azurerm_resource_group" "terraform_resource_group" {
  name     = "${var.project_name}-terraform-resource-group"
  location = var.location
  tags = {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

#data "template_file" "userdata" {
#  count                     = var.controllers == 3 ? var.controllers : 1 
#  #var.controller_ha != 3 ? var.controller_ha : 1
#  template                  = file("files/userdata.json")
#  vars                      = {
#    avi_controller_password = var.admin_password
#    dns                     = var.dns_ip[0]
#    dns1                    = var.dns_ip[1]
#    search_domain           = var.search_default_domain
#    welcome_banner          = var.welcome_banner
#    ntp_server1             = var.ntp_servers[0]
#    ntp_server2             = var.ntp_servers[1]
#    ntp_server3             = var.ntp_servers[2]
#    ntp_server4             = var.ntp_servers[3]
#  }
#}
resource "random_pet" "avi_ctl_name" {
  count  = var.controllers
  length = 1
}

# Create virtual machine
resource "azurerm_linux_virtual_machine" "virtualmachine" {
  count                           = var.controllers
  name                            = "${var.project_name}-terraform-${random_pet.avi_ctl_name[count.index].id}"
  admin_username                  = var.login_name
  admin_password                  = var.admin_password
  computer_name                   = "${var.computer_name}-${count.index + 1}"
  location                        = var.location
  resource_group_name             = azurerm_resource_group.terraform_resource_group.name
  size                            = var.controller_size #"Standard_D8s_v3"
  zone                            = element(local.azs, count.index % length(local.azs))
  #custom_data                     = (count.index % var.controllers) == 1 ? base64encode(data.template_file.userdata[0].rendered) : ""
  #custom_data                     = (count.index % var.controllers) == 0 ? base64encode(data.template_file.userdata.rendered) : ""
  disable_password_authentication = false
  network_interface_ids           = [
  azurerm_network_interface.terraform_network_interface[floor(count.index)].id]

  source_image_reference {
    publisher = var.avi_publisher
    offer     = var.product_and_offer
    sku       = var.sku_and_name
    version   = var.avi_azure_version
  }

  plan {
    name      = var.sku_and_name
    publisher = var.avi_publisher
    product   = var.product_and_offer
  }

  # Disk
  os_disk {
    name                 = "${var.project_name}-terraform-controller-disk-${random_pet.avi_ctl_name[count.index].id}"
    caching              = var.caching
    storage_account_type = var.storage_type
    disk_size_gb         = var.disk_size
  }

  tags = {
    environment     = "${var.project_name}-terraform-${var.project_environment}"
    owner           = var.owner_tag
    shutdown_policy = var.policy
  }
}
