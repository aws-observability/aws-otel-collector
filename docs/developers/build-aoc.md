### Build Artifacts

aws-otel-collector is an open source project, we’re looking for the contributions from the open source community to make it better. You can build your own executable binaries for your own customization. We provide the following command for you the build your own executables.

#### Build RPM file
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
make package-rpm
```

#### Build Deb file
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
.tools/packaging/debian/create_deb.sh
```

#### Build MSI file
```
git clone https://github.com/aws-observability/aws-otel-collector.git  
.\tools\packaging\windows\create_msi.ps1 
```
