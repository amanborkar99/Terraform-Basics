# Arrays and Slices in GO

## Arrays
* An array is a fixed-size collection of elements of the same type. The elements of the array are stored sequentially and can be accessed using their index.

* Declared as ` var x [3]int`.

## Slices
* A Slice is a segment of an array. Slices build on arrays and provide more power, flexibility, and convenience compared to arrays.

* Just like arrays, Slices are indexable and have a length. But unlike arrays, they can be resized.

* Internally, A Slice is just a reference to an underlying array.