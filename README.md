# aws-ip-checker
This tool checks whether IP addresses are set for the following AWS resources.
- SecurityGroup Ingress/Egress
- WAF v2 IPSet
- WAF Classic IPSet
- ApplicationLoadBlancer ListenerRule

## Install
### Mac
```
$ AWS_IP_CHECKER_VERSION=0.0.1
$ curl -OL https://github.com/kishii4726/aws-ip-checker/releases/download/v${AWS_IP_CHECKER_VERSION}/aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_darwin_amd64.zip

$ unzip aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_darwin_amd64.zip aws-ip-checker

$ sudo cp aws-ip-checker /usr/local/bin
```

### Linux
```
$ AWS_IP_CHECKER_VERSION=0.0.1
$ curl -OL https://github.com/kishii4726/aws-ip-checker/releases/download/v${AWS_IP_CHECKER_VERSION}/aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_linux_amd64.zip

$ unzip aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_linux_amd64.zip aws-ip-checker

$ sudo cp aws-ip-checker /usr/local/bin
```

## Usage

### Specify IP address as argument
```
$ aws-ip-checker xxx.xxx.xxx.xxx/xx
```

```
$ aws-ip-checker 0.0.0.0/0 111.111.111.111/32 222.222.222.222/32

AccountId: xxxxxxxxxxxx
+---------------+---------------------+---------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
|    SERVICE    |       DETAIL        |      RESOURCE       |                                                      ID,ARN                                                      |         IP         |
+---------------+---------------------+---------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| SecurityGroup | Ingress             | test01              | sg-01234567890                                                                                                   | 0.0.0.0/0          |
| WAF v2        | IPSet, Regional     | test-ap-northeast-1 | xxxxxxx-yyyy-zzzz-0000-01234567890                                                                               | 222.222.222.222/32 |
| WAF Classic   | IPSet, CloudFront   | test-cloudfront     | aaaaaaa-bbbb-cccc-0000-01234567890                                                                               | 111.111.111.111/32 |
| ALB           | Listener, Port: 443 | alb-test            | arn:aws:elasticloadbalancing:ap-northeast-1:xxxxxxxxxxxx:listener/app/alb-test/xxxxxxxxxxxxxxxx/yyyyyyyyyyyyyyyy | 222.222.222.222/32 |
| ALB           | Listener, Port: 80  | alb-test            | arn:aws:elasticloadbalancing:ap-northeast-1:xxxxxxxxxxxx:listener/app/alb-test/zzzzzzzzzzzzzzzz/aaaaaaaaaaaaaaaa | 111.111.111.111/32 |
+---------------+---------------------+---------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
```

### Specify csv as argument
```
$ aws-ip-checker xxx.csv
```

```
$ cat sample.csv
0.0.0.0/0,111.111.111.111/32,222.222.222.222/32

$ aws-ip-checker sample.csv

AccountId: xxxxxxxxxxxx
+---------------+---------------------+---------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
|    SERVICE    |       DETAIL        |      RESOURCE       |                                                      ID,ARN                                                      |         IP         |
+---------------+---------------------+---------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| SecurityGroup | Ingress             | test01              | sg-01234567890                                                                                                   | 0.0.0.0/0          |
| WAF v2        | IPSet, Regional     | test-ap-northeast-1 | xxxxxxx-yyyy-zzzz-0000-01234567890                                                                               | 222.222.222.222/32 |
| WAF Classic   | IPSet, CloudFront   | test-cloudfront     | aaaaaaa-bbbb-cccc-0000-01234567890                                                                               | 111.111.111.111/32 |
| ALB           | Listener, Port: 443 | alb-test            | arn:aws:elasticloadbalancing:ap-northeast-1:xxxxxxxxxxxx:listener/app/alb-test/xxxxxxxxxxxxxxxx/yyyyyyyyyyyyyyyy | 222.222.222.222/32 |
| ALB           | Listener, Port: 80  | alb-test            | arn:aws:elasticloadbalancing:ap-northeast-1:xxxxxxxxxxxx:listener/app/alb-test/zzzzzzzzzzzzzzzz/aaaaaaaaaaaaaaaa | 111.111.111.111/32 |
+---------------+---------------------+---------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
```