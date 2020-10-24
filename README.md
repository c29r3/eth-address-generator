# eth-address-generator  
Generates a private key that matches the suffix  

## Install
```
git clone https://github.com/c29r3/eth-address-generator.git \
&& cd eth-address-generator \
&& go get github.com/ethereum/go-ethereum/crypto \
&& go build suffix-generator.go
```  

Or download binary file for linux x64
```
wget https://github.com/c29r3/eth-address-generator/releases/download/v0.1/suffix-generator-linux \
&& chmod u+x suffix-generator-linux
```

## Usage  
```
./suffix-generator --help
Usage of ./suffix-generator:
  -suffix string
    	Address suffix, for example 0x12345 or 0xfffff (default "12345")
  -threads int
    	The number of threads the script will work with (default 5)
```
Example:
```
./suffix-generator -suffix 0x123 -threads 10
2020/10/24 18:39:02 Number of threads 10
Generating address with suffix 0x123
Found address with suffix 0x123 --> 0x123768690183ae46fd765f747a62c1c4fd629eea
```
