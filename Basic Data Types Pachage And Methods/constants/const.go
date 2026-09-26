package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// ============================================================
	// 1. MATHEMATICAL CONSTANTS
	// ============================================================

	// Euler's number.
	// Approximately 2.718281828...
	fmt.Println(math.E)

	// Pi.
	// Approximately 3.141592653...
	fmt.Println(math.Pi)

	// Golden ratio (Phi).
	// Approximately 1.618033988...
	fmt.Println(math.Phi)


	// ============================================================
	// 2. SQUARE ROOT CONSTANTS
	// ============================================================

	// Square root of 2.
	fmt.Println(math.Sqrt2)

	// Square root of Euler's number.
	fmt.Println(math.SqrtE)

	// Square root of Pi.
	fmt.Println(math.SqrtPi)

	// Square root of the golden ratio.
	fmt.Println(math.SqrtPhi)


	// ============================================================
	// 3. LOGARITHM CONSTANTS
	// ============================================================

	// Natural logarithm of 2.
	fmt.Println(math.Ln2)

	// Logarithm of E with base 2.
	fmt.Println(math.Log2E)

	// Natural logarithm of 10.
	fmt.Println(math.Ln10)

	// Logarithm of E with base 10.
	fmt.Println(math.Log10E)


	// ============================================================
	// 4. CALCULATING THE AREA OF A CIRCLE
	// ============================================================

	// Circle radius.
	circleRadius := 5.0

	// Circle area formula:
	//
	// Area = Pi × radius²
	//
	// math.Pow(circleRadius, 2)
	// means 5² = 25.
	area := math.Pi * math.Pow(circleRadius, 2)

	fmt.Println(area)
	// 78.53981633974483


	// ============================================================
	// 5. DEGREES TO RADIANS
	// ============================================================

	degrees := 90.0

	// Formula:
	//
	// radians = degrees × Pi / 180
	//
	// 90 degrees = Pi / 2 radians
	radians := degrees * math.Pi / 180

	fmt.Println(radians)
	// 1.5707963267948966


	// ============================================================
	// 6. FLOATING-POINT LIMITS
	// ============================================================

	// Maximum value that a float32 can represent.
	fmt.Println(math.MaxFloat32)

	// Smallest positive non-zero float32.
	fmt.Println(math.SmallestNonzeroFloat32)

	// Maximum value that a float64 can represent.
	fmt.Println(math.MaxFloat64)

	// Smallest positive non-zero float64.
	fmt.Println(math.SmallestNonzeroFloat64)


	// ============================================================
	// 7. CHECKING FLOAT64 RANGE
	// ============================================================

	var price float64 = 19.99

	// MaxFloat64 represents the largest finite float64 value.
	//
	// price is obviously smaller than MaxFloat64,
	// so this condition will be false.
	if price > math.MaxFloat64 {
		fmt.Println("Value is too large")
	}


	// ============================================================
	// 8. FLOATING-POINT OVERFLOW
	// ============================================================

	// IMPORTANT:
	//
	// math.MaxFloat64 is an untyped constant.
	//
	// If we write:
	//
	// value := math.MaxFloat64 * 2
	//
	// Go tries to evaluate the constant expression at compile time.
	// The result is too large to represent as a float64,
	// so Go produces a compile-time overflow error.
	//
	// Instead, explicitly convert MaxFloat64 to float64 first.

	maxFloat := float64(math.MaxFloat64)

	// This multiplication now happens as a float64 operation.
	//
	// The result is larger than the maximum finite float64,
	// so IEEE-754 floating-point arithmetic produces +Inf.
	value := maxFloat * 2

	fmt.Println(value)
	// +Inf

	// math.IsInf(value, 1)
	//
	// The second argument:
	//
	//  1  = positive infinity
	// -1  = negative infinity
	//  0  = either infinity
	fmt.Println(math.IsInf(value, 1))
	// true


	// ============================================================
	// 9. SIGNED INTEGER LIMITS
	// ============================================================

	// Maximum value of int on the current architecture.
	fmt.Println(math.MaxInt)

	// Minimum value of int on the current architecture.
	fmt.Println(math.MinInt)


	// int8 range:
	//
	// -128 to 127
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt8)


	// int16 range:
	//
	// -32768 to 32767
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MinInt16)


	// int32 range:
	//
	// -2147483648 to 2147483647
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)


	// int64 range:
	//
	// -9223372036854775808
	// to
	// 9223372036854775807
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)


	// ============================================================
	// 10. UNSIGNED INTEGER LIMITS
	// ============================================================

	// IMPORTANT:
	//
	// These constants are very large.
	//
	// For example:
	//
	// MaxUint64 =
	// 18446744073709551615
	//
	// An int64 cannot store this value.
	//
	// We explicitly convert the constants to the appropriate
	// unsigned integer type before passing them to Println.

	fmt.Println(uint(math.MaxUint))

	// Maximum uint8:
	//
	// 255
	fmt.Println(uint8(math.MaxUint8))

	// Maximum uint16:
	//
	// 65535
	fmt.Println(uint16(math.MaxUint16))

	// Maximum uint32:
	//
	// 4294967295
	fmt.Println(uint32(math.MaxUint32))

	// Maximum uint64:
	//
	// 18446744073709551615
	fmt.Println(uint64(math.MaxUint64))


	// ============================================================
	// 11. CONVERTING A STRING TO AN INTEGER
	// ============================================================

	// This is a string, not an integer.
	rawAge := "24"

	// strconv.ParseInt converts a string into an integer.
	//
	// Arguments:
	//
	// rawAge = value to convert
	// 10     = base 10 (decimal)
	// 64     = parse as a 64-bit integer
	//
	// It returns TWO values:
	//
	// age = converted number
	// err = possible error
	age, err := strconv.ParseInt(rawAge, 10, 64)

	// Always check the error when parsing external/user input.
	if err != nil {
		fmt.Println("Age must be a whole number.")
		return
	}


	// ============================================================
	// 12. CHECKING WHETHER THE AGE FITS INTO INT8
	// ============================================================

	// int8 can contain values from:
	//
	// -128 to 127
	//
	// Therefore, we check whether age is inside that range.

	if age < math.MinInt8 || age > math.MaxInt8 {
		fmt.Println("Age is outside the int8 range.")
		return
	}


	// ============================================================
	// 13. CONVERTING INT64 TO INT8
	// ============================================================

	// We know the value is inside the int8 range because
	// we checked it above.
	//
	// Therefore, this conversion is safe.
	smallAge := int8(age)

	fmt.Println(smallAge)
	// 24
}