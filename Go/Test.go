// ***** Pointer string
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func upperAllLetter(str *string) {
// 	fmt.Println(*str)
// 	*str = strings.ToUpper(*str)
// }

// func main() {
// 	name := "Weerasak"
// 	upperAllLetter(&name)
// 	fmt.Println(name)
// }

// ***** Pointer struct
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type Profile struct {
// 	Name string
// 	Age  int
// }

// func upperProfileName(p *Profile) {
// 	fmt.Println(p)
// 	p.Name = strings.ToUpper("hello")
// }

// func main() {
// 	p := Profile{
// 		Name: "Por",
// 		Age:  35,
// 	}
// 	upperProfileName(&p)
// 	fmt.Println(p)
// }

// ***** Array in golang copy address to function
// package main

// import (
// 	"fmt"
// )

// func doubleAllElement(nums []int) {
// 	for i := range nums {
// 		nums[i] *= 2
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	doubleAllElement(nums)
// 	fmt.Println(nums)
// }

// ***** json.Unmarshal
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	jsonText := []byte(`[1, 2, 3]`)
// 	var nums []int
// 	json.Unmarshal(jsonText, nums)
// 	// ที่ถูกต้องเป็นแบบนี้ json.Unmarshal(jsonText, &nums)
// 	fmt.Println(nums)
// }

// ***** Return address แทน return obj
// package main

// type App struct {
// 	db   *DB
// 	conf *Config
// }

// func NewApp() (*App, error) {
// 	db, err := NewDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	conf, err := NewConf()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &App{
// 		db:   db,
// 		conf: conf,
// 	}, nil
// }

// func main() {
// 	app, err := NewApp()
// 	if err != nil {
// 		panic(err)
// 	}
// 	app.Run()
// }
