# Exercise for Week 2 class

## Build a bi-directional Unit converter

The task for week 2 is to build a bi-directional (converts both ways) unit converter.

Create a `struct` type called `Converter`. This type will contain `methods` that can convert various units in a bi-directional way, that is, if we `attach` a method to the `Converter` struct called `FeetToCentimeter` that converts every value in `feet` to `centimeter`, we must have a corresponding `method` that converts `centimeter` to `feet` called `CentimeterToFeet`. Every unit must be represented as a __named type__.

The following demonstrate the steps to create a `FeetToCentimeter` & `CentimeterToFeet`:

+ Create a struct type called `Converter`.
+ Define the __named types__ for both units:

  + ```go

        type Feet float64
        type Centimeter float64

  ```go

  > Hint 1: All units "named types" should be float64

+ Define `FeetToCentimeter` and `CentimeterToFeet` methods with both _attached_ to the `Converter` struct.
+ Add a line in your `main` function that prints the result with sample data.

```go
    func main() {
        fmt.Println(CentimeterToFeet(13.50))
    }
```

> Hint 2: your function signatures should look like this `func CentimeterToFeet(c Centimeter) Feet { // code }`

## Compulsory exercises

+ Minutes to Seconds
+ Seconds to Milliseconds
+ Celsius to Fahrenheit
+ Fahrenheit to Celsius
+ Radian to Degree (`π radians` is equal to `180°`)
  + use the `math` package from the standard library to access `PI` _(math.Pi)_ value for better precision.
+ Degree to Radian
+ Centimeter to Feet
+ Feet to Centimeter
+ Kilogram to Pounds
+ Pounds to Kilogram

## Optional

+ Feel free to add extra conversions
