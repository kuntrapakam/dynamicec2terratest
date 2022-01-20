provider "aws" {
  region = var.awsRegion
}

resource "aws_instance" "webb" {
  ami               = var.ami
  availability_zone = var.availabilty_zone
  instance_type     = var.instance_type
  user_data = <<EOF
  #!/bin/bash
yum update -y
yum install python3 -y
curl -O https://bootstrap.pypa.io/get-pip.py
python3 get-pip.py
pip3 install --upgrade requests
pip3 install botocore
pip3 install boto3
pip3 install ec2instanceconnectcli
yum install -y https://s3.us-east-1.amazonaws.com/amazon-ssm-us-east-1/latest/linux_amd64/amazon-ssm-agent.rpm
EOF

  tags = {
    Name = "TERRATEST"
  }
}
