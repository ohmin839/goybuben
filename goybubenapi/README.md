# goybubenapi
`goybuben` provides minimum APIs.

## Conversion API

`ToAybuben` function converts ASCII strings into Aybuben strings.

```go
import (
    "fmt"

    goybuben "github.com/ohmin839/goybuben/api"  
)

result := goybuben.ToAybuben("Barev Dzez:")
fmt.Println(result); // Բարև Ձեզ։
```

### Conversion rules

#### Upper cases
| Input(ASCII) | Output(Unicode) |
| ------------ | --------------- |
| `A` | `Ա` |
| `B` | `Բ` |
| `G` | `Գ` |
| `D` | `Դ` |
| `E` | `Ե` |
| `Z` | `Զ` |
| `E'` | `Է` |
| `Y'` | `Ը` |
| `T'` | `Թ` |
| `Zh` | `Ժ` |
| `I` | `Ի` |
| `L` | `Լ` |
| `X` | `Խ` |
| `C'` | `Ծ` |
| `K` | `Կ` |
| `H` | `Հ` |
| `Dz` | `Ձ` |
| `Gh` | `Ղ` |
| `Tw` | `Ճ` |
| `M` | `Մ` |
| `Y` | `Յ` |
| `N` | `Ն` |
| `Sh` | `Շ` |
| `Vo` | `Ո` |
| `Ch` | `Չ` |
| `P` | `Պ` |
| `J` | `Ջ` |
| `Rr` | `Ռ` |
| `S` | `Ս` |
| `V` | `Վ` |
| `T` | `Տ` |
| `R` | `Ր` |
| `C` | `Ց` |
| `W` | `Ւ` |
| `P'` | `Փ` |
| `Q` | `Ք` |
| `O` | `Օ` |
| `F` | `Ֆ` |
| `U` | `Ու` |

#### Lower cases
| Input(ASCII) | Output(Unicode) |
| ------------ | --------------- |
| `a` | `ա` |
| `b` | `բ` |
| `g` | `գ` |
| `d` | `դ` |
| `e` | `ե` |
| `z` | `զ` |
| `e'` | `է` |
| `y'` | `ը` |
| `t'` | `թ` |
| `zh` | `ժ` |
| `i` | `ի` |
| `l` | `լ` |
| `x` | `խ` |
| `c'` | `ծ` |
| `k` | `կ` |
| `h` | `հ` |
| `dz` | `ձ` |
| `gh` | `ղ` |
| `tw` | `ճ` |
| `m` | `մ` |
| `y` | `յ` |
| `n` | `ն` |
| `sh` | `շ` |
| `vo` | `ո` |
| `ch` | `չ` |
| `p` | `պ` |
| `j` | `ջ` |
| `rr` | `ռ` |
| `s` | `ս` |
| `v` | `վ` |
| `t` | `տ` |
| `r` | `ր` |
| `c` | `ց` |
| `w` | `ւ` |
| `p'` | `փ` |
| `q` | `ք` |
| `o` | `օ` |
| `f` | `ֆ` |
| `u` | `ու` |
| `ev` | `և` |

#### Others
| Input(ASCII) | Output(Unicode) |
| ------------ | --------------- |
| `$` | `֏` |
| `0123456789` | `0123456789` |
| `,` | `,` |
| `.` | `.` |
| `` ` `` | `՝` |
| `:` | `։` |
| `-` | `-` |
| `(` | `(` |
| `)` | `)` |
| `<<` | `«` |
| `>>` | `»` |
| `?` | `՞` |
| `!` | `՛` |
| `!~` | `՜` |
| `␣` | `␣` |
| `\t` | `\t` |
| `\n` | `\n` |
| `\r\n` | `\r\n` |

## Collection API

`ToHayerenWords` function extracts Armenian words from texts:
```go
import (
    "fmt"

    goybuben "github.com/ohmin839/goybuben/api"  
)

converted := goybuben.ToAybuben("Barev Dzez:");
words := goybuben.ToHayerenWords(converted)
fmt.PrintLn(words); // [Բարև Ձեզ]
```

## External resources
- Armenian Alphabet (https://en.wikipedia.org/wiki/Armenian_alphabet)
- Romanization of Armenian (https://en.wikipedia.org/wiki/Romanization_of_Armenian)
