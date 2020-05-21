go build -o terraform-provider-canonical
rm -rf .terraform
terraform init
terraform providers schema -json

