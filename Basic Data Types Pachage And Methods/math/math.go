package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	n := -12.75

	fmt.Println(math.Abs(n))      // 12.75
	fmt.Println(math.Max(10, 25)) // 25
	fmt.Println(math.Min(10, 25)) // 10
	fmt.Println(math.Pow(2, 3))   // 8
	fmt.Println(math.Pow10(3))    // 1000
	fmt.Println(math.Sqrt(81))    // 9 // Square Root 9 × 9 = 81
	fmt.Println(math.Cbrt(27))    // 3 // Cube Root 3 × 3 × 3 = 27
	fmt.Println(math.Hypot(3, 4)) // 5 // √(3² + 4²)
	fmt.Println(math.Mod(10, 3))  // 1// remainder after division:
	fmt.Println(math.Dim(10, 6))  // 4
	fmt.Println(math.Dim(6, 10))  // 0
	/*
		Dim(x, y) means “positive difference”: it returns x - y when x > y; otherwise it returns 0.
	*/

	n := 2.6

	fmt.Println(math.Floor(n))         // 2
	fmt.Println(math.Ceil(n))          // 3
	fmt.Println(math.Trunc(n))         // 2
	fmt.Println(math.Round(n))         // 3
	fmt.Println(math.RoundToEven(2.5)) // 2
	fmt.Println(math.RoundToEven(3.5)) // 4

	//? For negative values, the difference matters:

	n := -2.6

	fmt.Println(math.Floor(n)) // -3: toward negative infinity
	fmt.Println(math.Ceil(n))  // -2: toward positive infinity
	fmt.Println(math.Trunc(n)) // -2: toward zero
	fmt.Println(math.Round(n)) // -3: nearest whole number

	/*
		//* Use Round for typical UI display rounding. RoundToEven is useful in some financial/statistical contexts
		//* because values exactly halfway between two integers go to the nearest even integer.
	*/

	fmt.Println(math.Exp(1))   // e raised to 1
	fmt.Println(math.Exp2(3))  // 2 raised to 3 = 8
	fmt.Println(math.Expm1(1)) // e^1 - 1

	fmt.Println(math.Log(math.E)) // 1: natural logarithm
	fmt.Println(math.Log10(100))  // 2
	fmt.Println(math.Log2(8))     // 3
	fmt.Println(math.Log1p(0.5))  // ln(1 + 0.5)

	angle := math.Pi / 2

	fmt.Println(math.Sin(angle)) // 1
	fmt.Println(math.Cos(0))     // 1
	fmt.Println(math.Tan(0))     // 0

	fmt.Println(math.Asin(1)) // pi/2
	fmt.Println(math.Acos(0)) // pi/2
	fmt.Println(math.Atan(1)) // pi/4

	// Angle of point (y=1, x=1)
	fmt.Println(math.Atan2(1, 1)) // pi/4

	positiveInfinity := math.Inf(1)
	negativeInfinity := math.Inf(-1)
	notANumber := math.NaN()

	fmt.Println(positiveInfinity)                // +Inf
	fmt.Println(negativeInfinity)                // -Inf
	fmt.Println(math.IsInf(positiveInfinity, 0)) // true
	fmt.Println(math.IsInf(positiveInfinity, 1)) // true
	fmt.Println(math.IsNaN(notANumber))          // true
	fmt.Println(math.Signbit(-0.5))              // true
	//! A NaN is an invalid numeric result, such as some impossible floating-point operations.
	//? Never check it with value == math.NaN() because NaN is not equal to anything, including itself;
	//* use math.IsNaN(value).

	fmt.Println(rand.Int())       // Random int
	fmt.Println(rand.IntN(10))    // 0 through 9
	fmt.Println(rand.Int64N(100)) // 0 through 99

	fmt.Println(rand.Uint32()) // Random unsigned 32-bit integer
	fmt.Println(rand.Uint64()) // Random unsigned 64-bit integer

	fmt.Println(rand.Float32()) // >= 0.0 and < 1.0
	fmt.Println(rand.Float64()) // >= 0.0 and < 1.0

	// 1 through 6: dice roll
	dice := rand.IntN(6) + 1

	// 50 through 100, inclusive
	score := rand.IntN(51) + 50

	// Decimal from 10.0 up to (but excluding) 20.0
	price := 10 + rand.Float64()*10

	fmt.Println(dice, score, price)

	//* IntN(n) always gives a value from 0 through n-1.
	//* Therefore, to generate an inclusive range from min through max:
	//? random := rand.IntN(max-min+1) + min

	languages := []string{"Go", "PHP", "JavaScript", "Python"}

	randomLanguage := languages[rand.IntN(len(languages))]

	fmt.Println(randomLanguage)

	numbers := rand.Perm(5)

	fmt.Println(numbers) // A random ordering of [0 1 2 3 4]

	cards := []string{"A", "K", "Q", "J"}

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	fmt.Println(cards)
	/*
	* Perm(n) produces a random permutation of integers from 0 through n-1;
	* Shuffle rearranges an existing slice.
	 */

	fmt.Println(rand.NormFloat64()) // Normally distributed value around 0
	fmt.Println(rand.ExpFloat64())  // Exponentially distributed positive value
	/*
	? These are mostly useful for
	? simulations, statistics, procedural generation, and probability modeling—not typical CRUD web apps.
	*/
}
