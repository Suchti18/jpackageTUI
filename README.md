<div align="center">

[![jpackageTUI](./.github/banner.svg)](#Installation)

</div>

# jpackageTUI

jpackageTUI is a Textual User Interface (TUI) designed to simplify the usage of the jpackage tool for
exporting java projects. The tool is currently in its early development phase, providing basic
functionality to streamline simple tasks.

---


## Installation
You can install it by [building](#Building) it yourself or by downloading the newest release

## Usage

Run the jpackageTUI executable file by typing:
```bash
.\jpackageTUI.exe
```

### Requirements

* jpackage must be installed and properly configured in the system's PATH.

### ⌨️ Keybindings

| Keybinding   | Description                   |
|--------------|-------------------------------|
| `esc`        | Quit the program              |
| `tab`        | Move between the input fields |
| `enter`      | Select                        |
| `down arrow` | Move down in a dropdown field |
| `up arrow`   | Move up in a dropdown field   |

### Example Output

Below is a sample of what the TUI looks like when running:

[![jpackageTUI Screenshot](./.github/sample.png)](#Installation)

---

## Building

### Requirements

* Git
* Go 1.24 or later

### System Requirements

* Windows (tested on Windows 11)
* ~~Linux (Tested on Ubuntu)~~
* ~~MacOS (Tested on MacOS Sequoia 15.5)~~

### Instructions

1. Clone the repository :
    ```bash
    git clone https://github.com/Suchti18/jpackageTUI.git
    ```
2. Change into the repository folder:
    ```bash
    cd jpackageTUI
    ```
3. Build the project:
    ```bash
    go build .\cmd\jpackageTUI\
    ```
   
After running `go build` an executable file was generated in the project folder.

---

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

---

## AI Disclaimer

AI helped in following areas:
* Debugging
* README creation

---

## License

[Unlicense](https://unlicense.org)