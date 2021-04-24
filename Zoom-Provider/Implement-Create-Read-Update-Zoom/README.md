# Implement Create

* ## Request & Response :-
     1. Created the Request structure and Schema for ResourceData.
     ``` GO
     type Create struct {
	Action    string `json:"action"`
	User_info Uz     `json:"user_info"`
    }

     type Uz struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Type       int    `json:"type"`
    }
     ```
     2. Assigned the schema value to the structure and then that strcuture is `Marsheled` and then `POST` to the url for user creation.

     3. Resquest Response is captured into the `Resp` structure.
     ``` GO
     type Resp struct {
	Id         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Type       int    `json:"type"`
    }
     ```
     4. Then created user is read by `func resourceUserRead` and made change in terraform statefile.

# Created Read
 1. Set Id as email and to retrive the information of the specific user.

 2. Using that **USER_ID / EMAIL** making a `GET` Request and we get `Response` for it which includes Id, email, etc.


# Implement Update
  1. First made the changes in already created user in terraform configuration file.
  
  2. Using the same structure of `struct Uz` changed the response and assigned values to it.

  3. Then using `PATCH` and `email` made the chnages and written it using Read function.
  
