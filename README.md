# FIXtoJSON

FIXtoJSON is a Go library that converts Financial Information eXchange (FIX) protocol messages to JSON format. This library is designed to simplify the process of working with FIX messages by providing an easy-to-use JSON representation.

## Features

- Converts raw FIX messages to structured JSON
- Handles repeating groups in FIX messages
- Supports nested JSON within FIX fields
- Uses FIX data dictionary for accurate field interpretation
- Handles unknown fields gracefully

## Why FIXtoJSON?

FIX protocol messages are widely used in financial trading systems but can be difficult to read and process due to their tag-value pair format. FIXtoJSON addresses this challenge by:

1. **Improving Readability**: Converting FIX messages to JSON makes them human-readable and easier to understand.

2. **Simplifying Integration**: JSON is a universally supported data format, making it easier to integrate FIX data with various systems and programming languages.

3. **Enhancing Debugging**: The JSON format makes it easier to inspect and debug FIX messages during development and troubleshooting.

4. **Supporting Complex Structures**: FIXtoJSON handles repeating groups and nested structures in FIX messages, preserving the hierarchical nature of the data.

5. **Facilitating Data Analysis**: JSON format allows for easier parsing and analysis of FIX message data using standard JSON tools and libraries.

## Installation

To install FIXtoJSON, use `go get -u github.com/goemon-xyz/fixtojson@latest`

## Contributing

Contributions to FIXtoJSON are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
