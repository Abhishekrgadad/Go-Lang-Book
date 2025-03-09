//Variable syntax 

var name type = expression

var s string
fmt.Println(s) // " "

var i, j, k int // i(int), j(int),k(int)
var b, f,s = true, 2.2 , "four"

//Short variable declaration
 
one := gif.GIF{something}
freq := rand.something
t := 0.0

Keep in mind that := is a declaration, where as = is an assig nment.
if any variable already declared inside a lexical block then := will act as an assignement to those variables.

Ex:
in, err := os.open(infile) //declaration

out, err := os.Create(outfile) //assignment

