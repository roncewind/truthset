# truthset

Generates a golang snippet file (snippet.go) from the specified Senzing truth set.
Used for testing the Senzing golang SDK.

## snippet.go

Note this isn't a completely valid golang file.  It's meant as a snippet to be
copied into a test when the truth set changes.

```
// A list of test records.
var ReferenceRecords = map[string]struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
"1001": {
	DataSource: "CUSTOMERS",
	Id:         "1001",
	Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robert", "DATE_OF_BIRTH": "12/11/1978", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 Main Street, Las Vegas NV 89132", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "702-919-1300", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "100"}`,
	LoadId:     "TRUTHSET_REFERENCE_LOAD",
},
"1002": {
	DataSource: "CUSTOMERS",
	Id:         "1002",
	Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1002", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith II", "PRIMARY_NAME_FIRST": "Bob", "DATE_OF_BIRTH": "11/12/1978", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Lane", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89111", "PHONE_TYPE": "MOBILE", "PHONE_NUMBER": "702-919-1300", "DATE": "3/10/17", "STATUS": "Inactive", "AMOUNT": "200"}`,
	LoadId:     "TRUTHSET_REFERENCE_LOAD",
},
...
}
```
