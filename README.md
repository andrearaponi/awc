# AWC (Andrea's Word Count) - Version 1.0.0

![Version](https://img.shields.io/badge/version-1.0.0-green.svg)

AWC is a custom implementation of the classic `wc` command from Unix coreutils, written in Go. This tool provides a simple and efficient way to count bytes, words, characters, and lines in a text file.

## Features

- Count bytes in a file with `-c` or `--bytes`.
- Count characters in a file with `-m` or `--chars`. Note: Each byte is treated as a character.
- Count words in a file with `-w` or `--words`.
- Count lines in a file with `-l` or `--lines`.
- Display usage information with `--help`.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. You can download it from [Go's official website](https://golang.org/).

### Installation

To install AWC, clone the repository and build the project:

1. Clone the repository:
```bash 
git clone https://github.com/andrearaponi/awc
```
2. Navigate to the repository:
```bash
cd awc
```


### Usage

To use AWC, simply run the following command:

```bash
./awc <file_path> [option]
```

For detailed usage instructions, use:

```bash 
./awc --help
```

## Contributing

Contributions to AWC are welcome! Please feel free to submit issues and pull requests.

## Buy me a coffee

Whether you use this project, have learned something from it, or just like it, please consider supporting it by buying me a coffee, so I can dedicate more time on open-source projects like this 

<a href="https://www.buymeacoffee.com/andrearapoA" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>


## License

This project is licensed under the [MIT License](https://github.com/andrearaponi/awc/blob/main/LICENSE).