/*
MIT License

Copyright (c) 2022 Zachinquarantine

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
astilectron "github.com/asticode/go-astilectron"
bootstrap "github.com/asticode/go-astilectron-bootstrap"
"math"
"fmt"
"log"
"os"
)

func init() {
// Initialize astilectron
var a,_ = astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
    AppName: "Coal Plant Calculator",
    AppIconDefaultPath: "<your .png icon>", // If path is relative, it must be relative to the data directory
    AppIconDarwinPath:  "<your .icns icon>", // Same here
    BaseDirectoryPath: "<where you want the provisioner to install the dependencies>",
    VersionAstilectron: "0.55.0",
    VersionElectron: "18.3.4",
})
}

// Start astilectron
a.Start()

defer a.Close()

// Blocking pattern
a.Wait()

// Create a new window
var w, _ = a.NewWindow("http://127.0.0.1:4001", &astilectron.WindowOptions{
    Center: astikit.BoolPtr(true),
    Height: astikit.IntPtr(600),
    Width:  astikit.IntPtr(600),
})
w.Create()