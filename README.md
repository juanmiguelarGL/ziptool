# ziptool

`ziptool` is a simple command-line tool for compressing and decompressing JSON data using **GZIP** and **Base64** encoding. It supports two main operations:

1. **encode**: Compresses a JSON file using GZIP and then encodes the result in Base64.  
2. **decode**: Reads Base64-encoded data from a file, decodes it, and then decompresses it from GZIP back to JSON.

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [encode](#encode)
  - [decode](#decode)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Fast and simple**: Quickly compress JSON data into a small, shareable Base64 string.
- **Easy decoding**: Decode the Base64 string back into the original JSON.
- **Pretty printing**: Optionally format the JSON output in a more readable way.

---

## Installation

1. Make sure you have **Go** installed (version 1.18+ recommended).
2. Clone or download this repository:

   ```bash
   git clone https://github.com/juanmiguelarGL/ziptool.git
