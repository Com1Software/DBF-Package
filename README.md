# DBF-Package
A Golang Package to Read and Write DBF files.

[DBF Links and Information](https://github.com/Com1Software/Test-DBF-Package/wiki/Test-DBF-Package)

[Testd DBF-Package](https://github.com/Com1Software/Test-DBF-Package)

[Go Playground for Programmers](https://github.com/Com1Software/Go-Playground-for-Programmers)

# Package functions:

### CheckForFile

Checks for full file name. Returns true if the file
is present and false if it is not.

### LoadFile

Loads a dbf file and returns a byte array.

### GetRecordCount

Returns the record count of the loaded dbf file.

### GetRecordLength

Returns the length of the records in the loaded dbf file.

### GetFirstRecordPosition

Returns the first records position in the loaded dbf file.

### GetRecord

Returns a specific record from the loaded dbf file.

### GetFieldCount

Returns the field count of the loaded dbf file.

### GetFieldName

Returns a specified fields name.

### GetFieldLength

Returns a specied fields lenghth.

### GetRecordField

Returns the Data in the specified Field and Record.
