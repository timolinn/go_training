# Composite Types

## Arrays

Arrays can be referred to as a __list__ of items of the same type. Go's Arrays are fixed size, this means that during declaration, you must specify the size of the array and type it will contain, these cannot be changed at runtime.

```go
    var children [2]string
    children[0] = "Elves"

    // or

    children := [2]string{"Elves", "Men"}
```

We can use `len` to get the length of the array. For cases where we don't know the number of elements we'll be dealing with upfront, we can turn to slices.

## Slices

Slices are lightweight structures that represent a portion of an array. You'll most likely use Slices way more than you'll be using Arrays in your programs:

```go
    // example underlying array: nums := [2, 3, 5, 6, 7..... 1000000]
    // slice => [ptr => 2, len, cap] == nums[:]

    children := []string{"Elves", "Men"}

    // or
    children := make([]string, 2, 5)
```

A Slice can also be created from an existing array

```go
    valar := [7]string{"Manwe", "Ulmo", "Aule", "Yavanna", "Varda", "Este", "Nienna"}
    ladies := valar[2:]
```

A Slice is always a pointer to an underlying array, it has a _length_ and _capacity_ properties, that can be retrived using the _len()_ and _cap()_ built-in functions respectively.

## Maps

Maps are key - value pairs structure. In some other languages, they are called Hashtables or Dictionaries.

```go
    // [key => value]
    // [ 0 => 1, 1 => 2]
    // [ "one" => 1, "two" => 2]

    clans := make(map[string]string)
    clans["Feanor"] = "Noldo Elves"

    // or
    clans := map[string]string{
        "Feanor" : "Noldo",
    }

    // you can access the data in a map like this
    fmt.Println(clans["Feanor"])

    // get the length
    l := len(clans)

    // Delete
    delete(clans, "Noldo")
```

## Structs

A struct is a data type that aggregates different kinds of data values of different types. A classic example of a struct would be the data representation of the of COVID19 status in Nigeria.

```go
    type COV19Stats struct {
        Country string
        NumberOfCases int
        Safety SafetyMeasures
    }

    type SafetyMeasures struct {
        First string
        Second string
        Third string
    }
```

We can initialize a struct using the built-in `new` function which returns a Pointer to the initialized struct value, to access the data within a struct, we use the `dot notation` as show below:

```go
    ng := new(COV19Stats)
    ng.Country = "Nigeria"
    ng.NumberOfCases = 46,
    ng.Safety = SafetyMeasure{}
    // or
    ng := COV19{
        Country: "Nigeria",
        NumberOfCases: 46,
        Safety: SafetyMeasure{}
    }
```

Structs can also be given behaviour by attaching methods to it:

```go
    func(c COV19Stats) GetCountry() string {
        return c.Country
    }
```
