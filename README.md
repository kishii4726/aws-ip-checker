# aws-ip-checker
This tool checks whether IP addresses are set for the following AWS resources
- SecurityGroup Ingress/Egress
- WAF v2 IPSet
- WAF Classic IPSet
- ApplicationLoadBlancer ListenerRule

## Install

## Usage

### Specify IP address as argument
```
$ aws-ip-checker xxx.xxx.xxx.xxx/xx
```

### Specify csv as argument
```
$ aws-ip-checker xxx.csv
```