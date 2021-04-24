# Implemented dataSourceUser

1. Created the `main.go` in Terraform-Provider-Zoom which initiates the Provider.

2. Then implemented `provide.go` file which contains the declaration of the features created in `zoom_corp` with schema which takes `token` as input from user and `providerConfigure` function which sets it as environment varaible.

3. Then Created the client.go which initialise the `NewClient` and model.go which defines the structure of user data.

4. Then Implemented the `dataSourceUser` in which `dataSourceUserRead` is defined along with schema to store the input from zoom.
