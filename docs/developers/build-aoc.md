### Build Artifacts

aws-otel-collector is an open source project, we’re looking for contributions from the open source community to make it better. You can build your own executable binaries for your own customization. We provide the following commands for you to build your own executables.

#### Build RPM file
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
cd aws-otel-collector
make package-rpm
```

#### Build Deb file
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
cd aws-otel-collector
make package-deb
```

#### Build MSI file
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
cd aws-otel-collector
choco install make
make windows-build
.\tools\packaging\windows\create_msi.ps1 
```
