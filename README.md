# aws-ip-checker
This tool checks whether IP addresses are set for the following AWS resources.
- SecurityGroup Ingress/Egress
- WAF v2 IPSet
- WAF Classic IPSet
- ApplicationLoadBlancer ListenerRule

## Install
### Mac(amd64)
```
$ AWS_IP_CHECKER_VERSION=0.0.3
$ curl -OL https://github.com/kishii4726/aws-ip-checker/releases/download/v${AWS_IP_CHECKER_VERSION}/aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_darwin_amd64.zip

$ unzip aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_darwin_amd64.zip aws-ip-checker

$ sudo cp aws-ip-checker /usr/local/bin
```

### Mac(arm64)
```
$ AWS_IP_CHECKER_VERSION=0.0.3
$ curl -OL https://github.com/kishii4726/aws-ip-checker/releases/download/v${AWS_IP_CHECKER_VERSION}/aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_darwin_arm64.zip

$ unzip aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_darwin_arm64.zip aws-ip-checker

$ sudo cp aws-ip-checker /usr/local/bin
```

### Linux
```
$ AWS_IP_CHECKER_VERSION=0.0.3
$ curl -OL https://github.com/kishii4726/aws-ip-checker/releases/download/v${AWS_IP_CHECKER_VERSION}/aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_linux_amd64.zip

$ unzip aws-ip-checker_v${AWS_IP_CHECKER_VERSION}_linux_amd64.zip aws-ip-checker

$ sudo cp aws-ip-checker /usr/local/bin
```

## Usage

### Specify IP address as argument
```
$ aws-ip-checker ip xxx.xxx.xxx.xxx/xx
```

```
$ aws-ip-checker ip 0.0.0.0/0 111.111.111.111/32 222.222.222.222/32

AccountId: xxxxxxxxxxxx
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
|    SERVICE    |           DETAIL           |         RESOURCE         |                                                      ID,ARN                                                      |         IP         |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| SecurityGroup | Ingress, Port: 443 - 443   | sg-test                  | sg-xxxxxxxxxxxxxxxxx                                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+                          +                                                                                                                  +--------------------+
| SecurityGroup | Ingress, Port: 3306 - 3306 |                          |                                                                                                                  | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF v2        | IPSet, Regional            | v2-ip-set-ap-northeast-1 | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF v2        | IPSet, CloudFront          | v2-ip-set-cloudfront     | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF Classic   | IPSet, Regional            | v1-ip-set-ap-northeast-1 | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 222.222.222.222/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF Classic   | IPSet, CloudFront          | v1-ip-set-cloudfront     | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| ALB           | Listener, Port: 80         | alb-test                 | arn:aws:elasticloadbalancing:ap-northeast-1:xxxxxxxxxxxx:listener/app/alb-test/xxxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxx | 0.0.0.0/0          |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
```

### Specify csv as argument
```
$ aws-ip-checker file xxx.csv
```

```
$ cat sample.csv
0.0.0.0/0,111.111.111.111/32,222.222.222.222/32

$ aws-ip-checker file sample.csv

AccountId: xxxxxxxxxxxx
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
|    SERVICE    |           DETAIL           |         RESOURCE         |                                                      ID,ARN                                                      |         IP         |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| SecurityGroup | Ingress, Port: 443 - 443   | sg-test                  | sg-xxxxxxxxxxxxxxxxx                                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+                          +                                                                                                                  +--------------------+
| SecurityGroup | Ingress, Port: 3306 - 3306 |                          |                                                                                                                  | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF v2        | IPSet, Regional            | v2-ip-set-ap-northeast-1 | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF v2        | IPSet, CloudFront          | v2-ip-set-cloudfront     | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF Classic   | IPSet, Regional            | v1-ip-set-ap-northeast-1 | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 222.222.222.222/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| WAF Classic   | IPSet, CloudFront          | v1-ip-set-cloudfront     | xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx                                                                             | 111.111.111.111/32 |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
| ALB           | Listener, Port: 80         | alb-test                 | arn:aws:elasticloadbalancing:ap-northeast-1:xxxxxxxxxxxx:listener/app/alb-test/xxxxxxxxxxxxxxxx/xxxxxxxxxxxxxxxx | 0.0.0.0/0          |
+---------------+----------------------------+--------------------------+------------------------------------------------------------------------------------------------------------------+--------------------+
```