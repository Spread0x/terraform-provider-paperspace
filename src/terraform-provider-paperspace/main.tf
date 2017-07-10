provider "paperspace" {
  apiKey = "1be4f97..."
  region = "East Coast (NY2)"
}

resource "paperspace_script" "my-script-1" {
  name = "My Script"
  description = "a short description"
  scriptText = <<EOF
  #!/bin/bash
  echo "Hello, World" > index.html
  nohup busybox httpd -f -p 8080 &
  EOF
  isEnabled = true
  runOnce = false
}

resource "paperspace_machine" "my-machine-1" {
  region = "East Coast (NY2)" // defaults to provider region if not specified
  machineType = "C1"
  size = 50
  billingType = "hourly"
  machineName = "Terraform Test",
  templateId = "tqalmii" // Ubuntu 16.04 Server
}
