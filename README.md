# ec2
Create/configure Amazon VPCs from command line

# step 1
Generate API keys and set them as environment variables
```
set AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
set AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

# step 2
Run the program with the name, network block, and subnets

```ec2.exe  -name vpc1 -cidr 10.0.0.0/16 -subnets 10.0.0.0/24@us-west-2a,10.0.1.0/24@us-west-2b```

# step 3
Go to AWS
