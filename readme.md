# GoCo: Go Bencoding Parser

GoCo is a lightweight and efficient bencoding parser implemented in Golang. Bencoding is a simple data serialization format used primarily in BitTorrent applications but can be useful in other contexts as well. This parser allows you to decode bencoded data into native Go data structures for easy manipulation and processing.

## Features

- **Fast Parsing**: GoCo is designed for speed and efficiency, making it suitable for parsing large bencoded data sets quickly.
- **Simple API**: With a straightforward API, GoCo makes it easy to decode bencoded data with just a few lines of code.
- **Native Go Types**: Decoded data is represented using native Go data types, making it easy to work with in your Go applications.
- **Error Handling**: GoCo provides comprehensive error handling, ensuring robustness in parsing various types of bencoded data.

## Installation

To install GoCo, use the `go get` command:

```bash
go get github.com/adiryjoshi/goco
```

## Usage
To run the program, execute the following command:
```bash
go run main.go
```
## Sample Inputs and Outputs

| Input Type     | Bencoded Data                | Decoded Output    |
|----------------|------------------------------|-------------------|
| Integer        | i42e                         | 42                |
| String         | 4:spam                       | spam              |
| List           | l4:spam4:eggse               | [spam eggs]       |
| Dictionary     | d3:cow3:moo4:spam4:eggse     | map[cow:moo spam:eggs] |

![ScreenRecording2024-05-08at12 15 36PM-ezgif com-speed](https://github.com/adityjoshi/GoCo/assets/111140014/4a6125cb-9088-471a-a647-2330d90d347c)


## Contributing

Contributions to GoCo are welcome! If you encounter any bugs or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/adityjoshi/goco).
