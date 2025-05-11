# 🐻 grizzly 🐻

A file parsing and analysis library for Go!

## 🗃️Basic File Extraction

To analyze a file call the parse library, providing a struct for the file layout.
File layouts can be fixed length or delimitted

import basic libary, calling it grizzly for short

```go

import (
    os
    grizzly "github.com/egotch/go-grizzly"
)

```

###  Create the field map
Create a field map for use in extracting data.
The field map consists of a slice of fields that define how
each line should be parsed

```go


```

### Extract data
Send a file to the extraction engine to get a data matrix back.
A data matrix consists of several details you may care about a file:

- file name and directory
- file size and number of lines
- map used to extract data (a FieldMap)
- slice of data rows

** Extracting a CSV file **
```go

    // create the field map
    fieldMap = grizzly.NewFieldMap(
            
        )

    // extracting a simple comma delimited csv file
    dm = grizzly.extract("./example/UserData.csv",
        AsCSV(","),
        WithFieldMap(fieldMap),
        )
```

**
