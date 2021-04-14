package section23

import "fmt"

//LexicalScope ...
func LexicalScope() {
	flagString1, flagString2 := "First flagString1", "First flagString2"
	{
		flagString1, flagString2, flagString3 := "Second flagString1", "Second flagString2", "First flagString3"
		fmt.Println(flagString1)
		fmt.Println(flagString2)
		fmt.Println(flagString3)
	}
	flagString3 := "Second flagString3"
	fmt.Println(flagString1)
	fmt.Println(flagString2)
	fmt.Println(flagString3)
}
