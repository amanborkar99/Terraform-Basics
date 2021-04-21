# Read operation on DataResource

## Prerequisites :-

* **Golang Installed** // To make provider
* **Terraform** // To control and add Resource
* **Docker** // launching the local Instance

***

1.  Clone the Repository with `boilerplate` branch [Hashicups Provider Repository](https://github.com/hashicorp/terraform-provider-hashicups/tree/boilerplate "Boilerplate Repository for Hashicups")

2. One can check the Coffee schema by launching a local instance in docker of Hashicup.

3. Created a Empty Provider inside the `hashicups/provider.go`i.e ResourceMap & DataSourceMap

4. Create a new file named `data_source_coffee.go` in the hashicups directory.

5. Make a function name as `dataSourecCoffees()` with return type *schema.Resource and inside that schema.Resource include **Schema** and `dataSourceCoffeeRead`.

6. Defined the **Schema** in it.

7. Implemented _dataSourceCoffeeRead_ libraries
    *  http.NewRequest - To create new request
    *  diag.Diagnostics - To return a final diagnostic report and error handling.

8. Finally Update the Provider funcion with DataSource value assigned 
``` GO
func Provider() *schema.Provider {
   return &schema.Provider{
       ResourcesMap: map[string]*schema.Resource{},
       DataSourcesMap: map[string]*schema.Resource{
            "hashicups_coffees":     dataSourceCoffees(),
      },
   }
}
```

9. Build the binary file of provider by command 
`go build -o terraform-provider-hashicups.exe`

8. Copy this built .exe file to the location:
`%APPDATA%\terraform.d\plugins\hashicorp.com\edu\hashicups\0.2\windows_amd64` according to OS Architecture.

9. Run main.tf in example directory to test the provider.
