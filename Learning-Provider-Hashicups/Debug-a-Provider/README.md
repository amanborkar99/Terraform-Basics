# Add Authentication to a Provider
##  AIM :- To access protected API endpoints.

## 1. Update Provider Schema.
   * Added `username` and `password` parameter to provider schema. 

   * Used `ConfigureContextFunc` property to the the name of the function that will configure provider.

## 2. Defined ProviderConfig.
   * This function actually configures provider. 

   * If both the username and password are not empty, the function returns an authenticated API client; else, the function returns an unauthenticated API client. 
   
   * The authenticated API client is able to access protected endpoints, the unauthenticated one is not.
