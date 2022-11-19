package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// CLONE
	a := "some"
	b := strings.Clone(a)
	fmt.Println(b) // some
	c := ""
	d := strings.Clone(c)
	fmt.Println(len(d)) //
	fmt.Println("_________________________________")

	// COMPARE
	fmt.Println(strings.Compare("a", "b")) //-1
	fmt.Println(strings.Compare("a", "a")) //0
	fmt.Println(strings.Compare("b", "a")) //1
	fmt.Println("_________________________________")

	//CONTAINS
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true
	fmt.Println("_________________________________")

	//CONTAINS ANY
	fmt.Println(strings.ContainsAny("team", "i"))     //false
	fmt.Println(strings.ContainsAny("fail", "ui"))    // true
	fmt.Println(strings.ContainsAny("ure", "ui"))     //true
	fmt.Println(strings.ContainsAny("failure", "ui")) //true
	fmt.Println(strings.ContainsAny("foo", ""))       //false
	fmt.Println(strings.ContainsAny("", ""))          //false
	fmt.Println("_________________________________")

	//CONTAINS RUNE
	fmt.Println(strings.ContainsRune("aardvark", 97)) //true
	fmt.Println(strings.ContainsRune("timeout", 97))  //false
	fmt.Println("_________________________________")

	//COUNT
	fmt.Println(strings.Count("cheese", "e")) //3
	fmt.Println(strings.Count("five", ""))    // 5
	fmt.Println("_________________________________")

	//CUT
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")     //Cut("Gopher", "Go") = "", "pher", true
	show("Gopher", "ph")     //Cut("Gopher", "ph") = "Go", "er", true
	show("Gopher", "er")     //Cut("Gopher", "er") = "Goph", "", true
	show("Gopher", "Badger") //Cut("Gopher", "Badger") = "Gopher", "", false
	fmt.Println("_________________________________")

	//EQUALFOLD
	fmt.Println(strings.EqualFold("Go", "go")) //true
	fmt.Println(strings.EqualFold("AB", "ab")) //true
	fmt.Println(strings.EqualFold("ß", "ss"))  //false
	fmt.Println("_________________________________")

	//FIELDS
	fmt.Println("First field are: ", strings.Fields(" bed good diner ")[0]) //	First field are:  bed
	fmt.Println("_________________________________")

	// FIELDS FUNC
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	fmt.Println("_________________________________") //Fields are:  ["foo1" "bar2" "baz3"]

	// HAS PREFIX
	fmt.Println(strings.HasPrefix("Gopher", "Go")) //true
	fmt.Println(strings.HasPrefix("Gopher", "C"))  // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true
	fmt.Println("_________________________________")

	// HAS SUFFIX
	fmt.Println(strings.HasSuffix("Amigo", "go"))  //true
	fmt.Println(strings.HasSuffix("Amigo", "O"))   //false
	fmt.Println(strings.HasSuffix("Amigo", "Ami")) //false
	fmt.Println(strings.HasSuffix("Amigo", ""))    //true
	fmt.Println("_________________________________")

	// INDEX
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dkr")) //-1
	fmt.Println("_________________________________")

	// INDEX ANY
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) //2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   //-1
	fmt.Println("_________________________________")

	// INDEX BYTE
	fmt.Println(strings.IndexByte("golang", 'g'))  //0
	fmt.Println(strings.IndexByte("gophers", 'h')) //3
	fmt.Println(strings.IndexByte("golang", 'x'))  //-1
	fmt.Println("_________________________________")

	// INDEX FUNC
	f1 := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f1))    //7
	fmt.Println(strings.IndexFunc("Hello, world", f1)) //-1
	fmt.Println("_________________________________")

	// INDEX RUNE
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
	fmt.Println("_________________________________")

	// JOIN
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	fmt.Println("_________________________________")

	// LAST INDEX
	fmt.Println(strings.Index("go gopher", "go"))         // 0
	fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1
	fmt.Println("_________________________________")

	// LAST INDEX ANY
	fmt.Println(strings.LastIndexAny("go gopher", "go"))     //4
	fmt.Println(strings.LastIndexAny("go gopher", "rodent")) //8
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))   // -1
	fmt.Println("_________________________________")

	// LAST INDEX BYTE
	fmt.Println(strings.LastIndexByte("Hello, world", 'l')) //10
	fmt.Println(strings.LastIndexByte("Hello, world", 'o')) //8
	fmt.Println(strings.LastIndexByte("Hello, world", 'x')) //-1
	fmt.Println("_________________________________")

	// LAST INDEX FUNC
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) //5
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) //2
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))     //-1
	fmt.Println("_________________________________")

	// MAP
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	fmt.Println("_________________________________")

	// REPEAT
	fmt.Println("ba" + strings.Repeat("na", 2)) // banana
	fmt.Println("_________________________________")

	// REPLACE
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
	fmt.Println("_________________________________")

	// REPLACE ALL
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) //moo moo moo
	fmt.Println("_________________________________")

	// SPLIT
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        //["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) //["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         //[" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            //[""]
	fmt.Println("_________________________________")

	// SPLIT AFTER
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) //["a," "b," "c"]
	fmt.Println("_________________________________")

	// SPLIT AFTER N
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))  //["a," "b,c"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,", ",", 3)) //["a," "b," "c,"]
	fmt.Println("_________________________________")

	// SPLIT N
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) //["a" "b,c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) //	[] (nil = true)
	fmt.Println("_________________________________")

	// TOLOWER
	fmt.Println(strings.ToLower("Gopher")) //gopher
	fmt.Println("_________________________________")

	// TOLOWER SPECIAL
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")) //önnek iş
	fmt.Println("_________________________________")

	// TOTITLE
	fmt.Println(strings.ToTitle("her royal highness")) // HER ROYAL HIGHNESS
	fmt.Println(strings.ToTitle("loud noises"))        //LOUD NOISES
	fmt.Println(strings.ToTitle("хлеб"))               //ХЛЕБ
	fmt.Println("_________________________________")

	// TOTITLE SPECIAL
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir")) // DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
	fmt.Println("_________________________________")

	// TOUPPER
	fmt.Println(strings.ToUpper("Gopher")) //GOPHER
	fmt.Println("_________________________________")

	// TOUPPER SPECIAL
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş")) //ÖRNEK İŞ
	fmt.Println("_________________________________")

	// TRIM
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers
	fmt.Println("_________________________________")

	// TRIM FUNC
	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) //Hello, Gophers
	fmt.Println("_________________________________")

	// TRIM Left
	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers!!!
	fmt.Println("_________________________________")

	// TRIM Left FUNC
	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) //Hello, Gophers!!!
	fmt.Println("_________________________________")

	// TRIM PREFIX
	var s2 = "¡¡¡Hello, Gophers!!!"
	s2 = strings.TrimPrefix(s2, "¡¡¡Hello, ")
	s2 = strings.TrimPrefix(s2, "¡¡¡Howdy, ")
	fmt.Println(s2) //Gophers!!!
	fmt.Println("_________________________________")

	// TRIM RIGHT
	fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")) //¡¡¡Hello, Gophers
	fmt.Println("_________________________________")

	// TRIM RIGHT FUNC
	fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) //¡¡¡Hello, Gophers
	fmt.Println("_________________________________")

	// TRIM SPACE
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) //Hello, Gophers
	fmt.Println("_________________________________")

	//TRIM SUFFIX
	var s3 = "¡¡¡Hello, Gophers!!!"
	s3 = strings.TrimSuffix(s3, ", Gophers!!!") //¡¡¡Hello
	s3 = strings.TrimSuffix(s3, ", Marmots!!!")
	fmt.Print(s)
	fmt.Println("_________________________________")
}
