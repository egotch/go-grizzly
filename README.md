# 🐻 grizzly 🐻

A file parsing and analysis library for Go!

## 🗃️Basic File Parsing

To analyze a file call the parse library, providing a struct for the file layout.
File layouts can be fixed length or delimitted

import basic libary, calling it grizzly for short

```go

import (
    os
    grizzly "github.com/egotch/go-grizzly"
)

```

Open a file and send it to the parsing engine to get a slice of
dataRows back

```go

   dataRows := grizzly.parse("./samples/UserData.csv", &grizzly.FieldMap{
        Username string,
        FirstName string,
        LastName string,
        Age int,
   }) 
```
