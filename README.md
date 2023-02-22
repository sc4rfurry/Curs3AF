<h1 align="center">
  Curs3AF
</h1>

<h4 align="center">Web Application Firewall fingerprinting tool.</h4>
<div style="text-align:center">
    <div style="align:center">
    <img src="https://img.shields.io/badge/Author-sc4rfurry-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Author">
    <img src="https://img.shields.io/badge/Version-1.0.1-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Version">
    <img src="https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17" alt="Go Version">
    <img src="https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17" alt="OS">
    <img src="https://img.shields.io/badge/Go_Library-wafme0w-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17" alt="Go Library">
    </div>
This tool uses <span style="color:cyan">Go Library</span> <a href ="https://github.com/Lu1sDV/wafme0w">wafme0w</a> a fork of <a href ="https://github.com/EnableSecurity/wafw00f/">Wafw00f</a> to detect if a website is protected by a WAF.
</div>

#

## Table of Contents

- [Installation](#installation)
- [Running Curs3AF](#running-curs3af)
- [Example](#example)
- [Building Binaries](#building-curs3af)
- [References](#references)
- [Contributing](#contributing)
- [License](#license)


#

### 🔧 Technologies & Tools

![](https://img.shields.io/badge/Editor-VS_Code-informational?style=flat-square&logo=visual-studio&logoColor=blue&color=5194f0)
![](https://img.shields.io/badge/Language-Go-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)
![](https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)

#

### 📚 Requirements
> - Go 18.1 linux/amd64

#
### Installation

- sudo apt-get update && sudo apt-get golang
- git clone https://github.com/sc4rfurry/Curs3AF.git
- cd Curs3AF
- go get .
- go build main.go
    - or use the `builder.sh` script to build the tool.


### Features

- Can detect **153** different Firewalls
- Concurrent fingerprinting
- Scan Multiple Domains from a file
- Fast detection mode (only checks for the most common WAFs - **Optional**)

#

## Running Curs3AF
```sh
go run main.go --help
go run main.go -u https://www.google.com
go run main.go -f domains.txt -g
```


### Example

To run the tool on a target, just use the following command.

```console
go run main.go --url asgoogle.com

   ____                        _____      _      _____
  / ___|  _   _   _ __   ___  |___ /     / \    |  ___|
 | |     | | | | | '__| / __|   |_ \    / _ \   | |_
 | |___  | |_| | | |    \__ \  ___) |  / ___ \  |  _|
  \____|  \__,_| |_|    |___/ |____/  /_/   \_\ |_|


  Description: Tool to check if a website is protected by a WAF(HTTP/HTTPS).


	Author: 	 sc4rfurry
	Version: 	 1.0.1
	Go Version: 	 1.18.1 or higher
	Github: 	 https://github.com/sc4rfurry
=================================================================================================


[info] Starting WAF Detection on asgoogle.com
[info] Running in Normal Mode (Scan for all 153 Wafs)- Could take time to scan



[!] http://asgoogle.com is protected by [{AWS Elastic Load Balancer (Amazon)}]
[!] https://asgoogle.com is protected by [{AWS Elastic Load Balancer (Amazon)}]
```

#

## Building Curs3AF
> To build the tool, you can use the following command.
```sh
env GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static"' -o Curs3AF main.go
```

> You can also use the bultin Bash script to build the tool.

- Before running the script, make sure to give it execution permissions.
- The bash script can build both Linux and Windows binaries.
- Binaries will be Stripped and Compressed. (lolcat, strip and upx are required)
```sh
chmod +x builder.sh
./builder.sh main.go
```
#
## Pre-Compiled Binaries
<div>
<div style="text-align:center">
    <a href="https://github.com/sc4rfurry/Curs3AF/releases/tag/v1.0.1">
    <img src="https://img.shields.io/badge/Download-v1.0.1-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Download">
    </a>
<div style="text-align:center">
    <img src="https://img.shields.io/badge/Status-Active-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Status">
</div>
</div>
</div>

#

## References
* Special thanks to the following projects (they are the base of this tool):
    - [Wafw00f](https://github.com/EnableSecurity/wafw00f)
    - [Wafme0w](https://github.com/Lu1sDV/wafme0w)

#

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

[def]: https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17