# Implement Complex Read

1. Created a new Order via API.
   * There must be an existing order for you to retrieve it.

2. Defined `order` data source.
   * Added a scaffold that defines an empty schema and functions to retrieve order information.

3. Defined `order` schema.
   * The schema reflects the response from invoking the `/order/{orderId}` endpoint.

4. Implemented Complex Read.
   * This read functionality uses multiple nesting functions to flatten the response from `/order/{orderId}` and map it to the `order` schema.

5. Added `order` data source to the provider schema.
   * This is done to to use the data source in during configuration.
