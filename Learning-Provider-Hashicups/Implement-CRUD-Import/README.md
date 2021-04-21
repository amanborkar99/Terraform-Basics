# Implement Create in hashicup-Provider

1. Defined the `order` resource.
   * Add a scaffold that defines an empty schema and functions to create, read, update and delete orders.

2. Defined `order` schema.
   * The schema contains fields used to create a new order and retrieved by the read functionality.

3. Add create functionality.
   * The create function invokes a `POST` request to `/orders` to create a new order.

4. Add read functionality.
   * Read function invokes a `GET` request to `/orders/{orderID}` to retrieve the order's data. 

   * The orders schema contains a nested object, therefore used multiple flattening functions to map the response to the order schema.

5. Then finally, Add `order` resource in provider.go
   * Allow to use resource in Configuration file.


# Implement Update 

1. Update the Schema
   * Added `last_update` attribute to the `order` schema. This attribute will update to the current date time whenever the order resource is updated.

2. Implement update
   * This fucntion invokes a `PUT` request to the `/orders/{orderID}` endpoint with update order items in the request body, then it invoke resorceorderread to get the updated values.

# Implement Delete

1. The delete function invokes a `DELETE` request to the `/orders/{orderId}` endpoint.

# Implement Import

1. Implemented a `ResourceImporter()` function. This function retrieves an existing HashiCups order and brings it into Terraform state, enabling Terraform to manage it.

