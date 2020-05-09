# gojaconv

## 日本語変換package

- かな -> ローマ字(ヘボン式)
  - ToHebon(string) string

## Install

```
go get -u github.com/kotaroooo0/gojaconv/jaconv
```

## Usage

```Go
import 	"github.com/kotaroooo0/gojaconv/jaconv"

hebon := jaconv.ToHebon("おはよう")
fmt.Println(hebon)
// Output: ohayo

hebon = jaconv.ToHebon("こんにちは")
fmt.Println(hebon)
// Output: konnichiha
```

## Author

kotaroooo0

## LICENSE

MIT License
