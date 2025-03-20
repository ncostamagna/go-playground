

# Rune
In Go (Golang), a rune is a type that represents a Unicode code point.

## 🔤 What is a Unicode code point?
A Unicode code point is a number that maps to a character in the Unicode standard, which includes letters, symbols, emojis, and more from almost every writing system.

## 🧠 Key Points about rune in Go:
A rune is an alias for int32.
It is used to represent a single Unicode character.
Useful when working with multi-byte characters like emojis or non-English letters.

## ✅ Syntax Example:
```go
var ch rune = 'A'
fmt.Println(ch)        // Outputs: 65 (ASCII code)
fmt.Printf("%c\n", ch) // Outputs: A
```
```go
emoji := '😊'
fmt.Println(emoji)        // Outputs: 128522 (Unicode code point)
fmt.Printf("%c\n", emoji) // Outputs: 😊
```

## 🤔 Why use rune instead of byte?
byte is used for ASCII (1-byte) characters.
rune is used for Unicode (can be more than 1 byte).
For example, 'A' fits in a byte, but 'ñ' or '😊' needs a rune.

### ⚔️ rune vs ASCII in Go

| Feature         | `rune` (Unicode)                               | `ASCII`                            |
|-----------------|------------------------------------------------|------------------------------------|
| **Type in Go**  | `rune` (alias for `int32`)                     | Typically `byte` (alias for `uint8`) |
| **Represents**  | Any **Unicode** character                      | Only **ASCII** characters (0–127)  |
| **Character Set** | Supports global characters (`ñ`, `漢`, `😊`) | Only basic English letters, numbers, symbols |
| **Byte size**   | 4 bytes (can store large Unicode code points)  | 1 byte                             |
| **Example char**| `'ñ'`, `'😊'`, `'日'`                           | `'A'`, `'Z'`, `'0'`, `'@'`         |
| **Use case**    | International apps, emojis, symbols            | Simple text, config, English-only  |

