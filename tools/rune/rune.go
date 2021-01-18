package main

/*

	see this doc on rune and strings: https://blog.golang.org/strings

*/

import "fmt"

func main() {
	// test1()
	// test2()
	// test3()
	printBr()
}

// Some simple stuff for testing
func printBr() {
	// https://unicode.org/cldr/utility/character.jsp?a=2834
	chars := []string{"\u2834", "\u2835"}
	for c := 0; c <= 10; c++ {
		fmt.Println("\n")
		for _, v := range chars {
			fmt.Print(v)

		}
	}

}

func test3() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

// Use a btye slices not a string litteral
func test2() {

	var sample []byte = []byte{'\xbd', '\xb2', '\x3d', '\xbc', '\x20', '\xe2', '\x8c', '\x98'}

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}

func test1() {

	// Store a string of bytes in hex
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	// Print Gibberish
	fmt.Println(sample)

	// formats the bytes as hex - the space is added for formatting "% x"
	fmt.Printf("% x\n", sample)

	// %q escapes unprintable stuff "+" escapes non-ascii
	// Notice that \xbd is the escaped non printable. and \u2318 is non-ascii
	fmt.Printf("%q\n", sample)  // escape non-print                "\xbd\xb2=\xbc ⌘"
	fmt.Printf("%+q\n", sample) // escape non-print and non-ascii  "\xbd\xb2=\xbc \u2318"

}
