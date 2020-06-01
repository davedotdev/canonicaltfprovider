go build -o terraform-provider-canonical
rm -rf .terraform
rm terraform.tfstate
rm crash.log 
terraform init

