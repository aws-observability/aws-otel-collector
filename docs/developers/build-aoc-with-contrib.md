## Example: How to build OTEL Collector with a Contrib Module 

In this example you will build OTEL Collector with the SQL Query Receiver on a Windows host.  


### Prerequisites and clone code
```
# Install Chocolatey if not already installed 
choco install make git go 
git clone https://github.com/aws-observability/aws-otel-collector.git  
cd aws-otel-collector
```

### Modify the code to include contrib receiver code

**Add module to "aws-otel-collector\go.mod"**

You can search for the word "receiver" and add a line with the SQL Query Receiver after the last receiver line: 
```
github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver vX.XX.X
```
Make sure you substitue the version number (vX.XX.X) with the correct version 

**Add the new receiver information to the default components file**

Edit the file "aws-otel-collector\pkg\defaultcomponents\defaults.go"

Add line to include the SQL Query Receiver module inside the "import" section: 
```
"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"
```
**Add the new receiver to the "receiverList"**

Search for "receiverList" and add ath the end of the list: 
```
sqlqueryreceiver.NewFactory(),
```
**Add reference to the testbed file**

Edit the file "aws-otel-collector\testbed\go.mod"
Search for "receiver" and add an entry for the SQL Query Receiver code
```
github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver v0.80.0 //indirect
```

### Refresh the modules
**In main folder run**
```
go clean -modcache
go mod tidy
go mod download
go mod vendor
```
**Inside the testbed folder run**
```
go mod tidy
go mod download
go mod vendor
```

### Build the MSI

```
cd aws-otel-collector
make windows-build
.\tools\packaging\windows\create_msi.ps1 
```

