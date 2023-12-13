# Utils

A Set of useful functions for working with files, errors and strings.

## Functions

### If

Generate quick if result.

```go
// Signature:
func If[T any](cond bool, yes T, no T) T

// Example:
res := If[string](name == "john", "it's john", "anonymous")
```

### Try

Simulate try catch block.

```go
// Signature:
func Try(e error, fail func(e error), pass func(), finally ...func(e error))

// Example:
Try(
    someErr,
    func(e) {
        // e error happend!
    },
    func() {
        // no error happend
    },
    func(e) {
        // call on all state
    },
)
```

### Contains

Check if slice contains item.

```go
// Signature:
func Contains[T comparable](items []T, item T) bool

// Example:
res := Contains[string](items, "john")
```

### Alter

Return fallback if value is empty.

```go
// Signature:
func Alter[T comparable](value T, fallback T) T

// Example:
name := ""
res := Alter(name, "John") // "John"
```

## Error

### TaggedError

Generate a tagged error.

```go
// Signature:
TaggedError(tags []string, format string, args ...any) error

// Example:
TaggedError([]string{"MyLib","MyMethod"}, "failed on %s file!", "main.json")
// [MyLib] [MyMethod] failed on main.json file!
```

### IsErrorOf

Check if error has tag.

```go
// Signature:
IsErrorOf(tag string, err error) bool

// Example:
IsErrorOf("MyLib", err) // true
```

### HasError

Check if error is nil or not.

```go
// Signature:
HasError(err error) bool
```

### PanicOnError

Generate panic from error if error not nil.

```go
// Signature:
PanicOnError(err error)
```

### VarOrPanic

Get function result (T, error), panic error if result has error and return T otherwise.

```go
// Signature:
VarOrPanic[T any](res T, err error) T
```

### PrettyLog

Log data to output using json format.

```go
// Signature
PrettyLog(data any)
```

### PrettyLogIndent

Log data to output using json format with indent.

```go
// Signature
PrettyLogIndent(data any)
```

## File

### FileExists

Check if file exists or not.

```go
// Signature:
FileExists(path string) (bool, error)

// Example:
import "github.com/gomig/utils"
exists, err := utils.FileExists("path/to/file")
```

### IsDirectory

Check if path is a directory.

```go
// Signature:
IsDirectory(path string) (bool, error)

// Example
import "github.com/gomig/utils"
ok, err := utils.IsDirectory("path/to/dir")
```

### FindFile

Search for files in directory by a regex pattern.

```go
// Signature:
FindFile(dir string, pattern string) []string

// Example:
import "github.com/gomig/utils"
files := utils.FindFile("path/to/dir", ".+\.sql") // => Get All file with sql extension
```

### ClearDirectory

Delete all files and sub-directory in directory.

```go
// Signature:
ClearDirectory(dir string) error
```

### GetSubDirectory

Get list of sub directories.

```go
// Signature:
GetSubDirectory(dir string) ([]string, error)
```

### CreateDirectory

Create nested directory.

```go
// Signature:
CreateDirectory(path string) error

// Example:
import "github.com/gomig/utils"
err := utils.CreateDirectory("a/b/c/d") // => Create all a, b, c and d directory
```

### DetectMime

Detect file mime info from content

```go
// Signature:
DetectMime(data []byte) *mimetype.MIME

// Example:
if mime := DetectMime(myFileData); mime != nil {
    // do something
}
```

### Extension

Get file extension.

```go
// Signature:
Extension(file string) string

// Example
Extension("file") // ""
Extension("file.JPG") // ".jpg"
Extension("file.png") // ".png"
Extension("file.") // "."
```

### NumberedFile

Generate unique numbered file until 10000000, e.g. file.txt file-1.txt, file-2.txt

```go
// Signature:
NumberedFile(dir, name, file string) (string, error)

// Example:
import "github.com/gomig/utils"
files := utils.NumberedFile("path/to/dist", "my-file" ,"tempfile-abcd.zip")
```

## String

### ExtractNumbers

Extract numbers from string.

```go
import "github.com/gomig/utils"
numbers := utils.ExtractNumbers("(+1) 234-56789") // => 123456789
```

### ExtractAlphaNum

Extract alpha and numbers from string `[a-zA-Z0-9]`. You can add extra character to add in extraction.

```go
import "github.com/gomig/utils"
numbers := utils.ExtractAlphaNum("this is a: 123", ":") // => "thisisa:123"
```

### ExtractAlphaNumPersian

Extract persian alpha, alpha and numbers from string `[ا-یa-zA-Z0-9]`. You can add extra character to add in extraction.

```go
import "github.com/gomig/utils"
numbers := utils.ExtractAlphaNumPersian("My name is: مجتبی", " ") // => "My name is مجتبی"
```

### RandomStringFromCharset

Generate random string from character list.

```go
import "github.com/gomig/utils"
str, err := utils.RandomStringFromCharset(5, "1234567890") // => "59102"
str2, err2 := utils.RandomStringFromCharset(3, "ABCDEFGH") // => "DFC"
```

### RandomString

Generate random string from Alpha-Num Chars

```go
import "github.com/gomig/utils"
str, err := utils.RandomString(5) // => "AB5S2"
```

### Slugify

Generate dash separated string.

```go
import "github.com/gomig/utils"
str := utils.Slugify("welcome to", "my site") // => "welcome-to-my-site"
```

### SlugifyPersian

Make slugify string for persian string. this function only keep `persian alphabet`, `a-z`, `A-Z` and `0-9` characters.

```go
import "github.com/gomig/utils"
str := utils.SlugifyPersian("خوش آمدید \n \r \t - to گچپژ") // => "خوش-آمدید-to-گچپژ"
```

### ConcatStr

Join strings with separator.

```go
// Signature:
ConcatStr(sep string, str ...string) string

// Example:
import "github.com/gomig/utils"
str := utils.ConcatStr(" ", "John", "", "Doe") // => "John Doe"
```

### FormatNumber

Format number with comma separator.

```go
import "github.com/gomig/utils"
func FormatNumber(format string, v ...any) string {
str := utils.FormatNumber("$ %d [total] $ %d [remain]", 10000, 2500) // => "$ 10,000 [total] $ 2,500 [remain]"
```
