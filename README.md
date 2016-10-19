# aws-auth-proxy

## How To Use

### on Your Machine

```
$ AWS_ACCESS_KEY_ID=XXXXX AWS_SECRET_ACCESS_KEY= XXXXX ./aws-auth-proxy -dest "https://search.....es.amazonaws.com"
```

### on EC2 with IAM Role

```
$ ./aws-auth-proxy -dest "https://search.....es.amazonaws.com"
```

example IAM Role Policy

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "es:ESHttp*"
            ],
            "Resource": [
                "arn:aws:es:<region>:<account>:domain/<domain>"
            ]
        }
    ]
}
```



