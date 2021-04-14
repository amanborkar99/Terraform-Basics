# Maps in GO

* A **map** is an unordered collection of key-value pairs. It maps keys to values. The keys are unique within a map while the values may not be.

* Declaration
  ``` GO
    var m map[key_type]value_type
  ```

* The map data structure is used for fast lookups, retrieval, and deletion of data based on key.

* **Maps** are reference types. When you assign a map to a new variable, they both refer to the same underlying data structure. Therefore changes done by one variable will be visible to the other