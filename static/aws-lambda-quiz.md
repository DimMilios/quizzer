## AWS Lambda

#### Q1. How can you increase the CPU resources for your Lambda?

- [ ] Increase the configured CPU value
- [ ] Increase the configured timeout value
- [x] Increase the configured memory value
- [ ] Increase the configured concurrency value

#### Q2. How can additional code or content be provided for your Lambda?

- [ ] blocks
- [x] layers
- [ ] aliases
- [ ] handlers

#### Q3. How can Step Functions call Lambdas?

- [ ] in sequence
- [x] both of these answers
- [ ] neither of these answers
- [ ] in parallel

#### Q4. Which AWS CLI command invokes a function?

- [ ] aws lambda invoke --function ReturnBucketName outputfile.txt
- [ ] aws lambda execute --function-name ReturnBucketName outputfile.txt
- [x] aws lambda invoke --function-name ReturnBucketName outputfile.txt
- [ ] aws lambda execute --function ReturnBucketName outputfile.txt

#### Q5. What adds tracing capabilities to a Lambda?

- [ ] AWS Trace
- [ ] CloudStack
- [ ] CloudTrail
- [x] AWS X-Ray

#### Q6. You need to build a continuous integration/deployment pipeline for a set of Lambdas. What should you do?

- [ ] Create configuration files and deploy them using AWS CodePipeline.
- [ ] Create CloudFormation templates and deploy them using AWS CodeBuild
- [ ] Create configuration file and deploy using AWS CodeBuild
- [x] Create CloudFormation templates and deploy them using AWS CodePipeline.

#### Q7. What can you use to monitor function invocations?

- [ ] API Gateway
- [ ] S3
- [ ] SAS
- [x] CloudTrail

#### Q8. It is AWS best practice to enable Lambda logging by which of these methods.

- [ ] Use S3 metrics and CloudWatch alarms
- [ ] Create custom metrics within your Lambda code.
- [ ] Create custom metrics within your CloudWatch code.
- [x] Use Lambda metrics and CloudWatch alarms.

#### Q9. What may be provided for environment variables?

- [ ] an SSL certificate
- [ ] a bitmask
- [x] an AWS KMS key
- [ ] an HTTP protocol

#### Q10. Lambdas allow for running of what other things?

- [ ] binaries.
- [x] all of these answers
- [ ] executables
- [ ] Shell scripts