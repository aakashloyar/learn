# S3
Some Terminologies
-> Bucket 
-> Object
-> Key

-> Bucket 
Folder or container where we store objects
-> Object
That we want to store
-> Key
Path inside container

-> how permission work in S3
Access managed using:
IAM
Bucket policies
ACLs
Pre-signed URLs

-> Pre-Signed URLs
-> Backend generates temporary secure upload/download link
-> User uploads directly to s3 without exposing AWS credentials



-> ACLs
-> ACL = Access Control List
-> it was used in old system
-> so there are two type of permissions 
-> resource level and object level
-> acl works at two level bucket, object
-> commmon acl permission
Read
Write
Full_Control
Read_ACP -> Read ACL itself
Write_ACP -> Write ACL itself

-> Common predefined acls
-> private
-> public-read
-> public-read-write
-> authenticated-read

-> how acl looks
-> ACL contains
-> Grantee -> who, Permission

-> Why now moving away from acl

-> Too complex
-> permission come from iam, bucket, acl
-> hard to debug

-> Security issues
-> People make files mistakenly public
aws s3 cp image.png s3://bucket --acl public-read

-> Object ownership problems
-> another aws account uploads file
-> who owns it
-> bucket owner
-> uploader
-> suppose uploader become owner
-> now bucket owner wants to delete object
-> cannot because he is not owner of this object
-> very frusating in production


-> Security team
-> prefer centeralized policy 
-> while acl is distributed per object

-> modern aws recommendation
-> disable acls
-> use iam, bucket, private-signed urls
-> basically why aws recommend to move away bcz to keep it simple


-> Bucket policy 
-> json attached to bucket
-> it defines who owns this bucket and what can they do

-> structure of bucket policy
| Field     | Meaning              |
| --------- | -------------------- |
| Effect    | Allow or Deny        |
| Principal | Who                  |
| Action    | What operation       |
| Resource  | Which bucket/object  |
| Condition | Optional extra rules |

-> simple example
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::my-bucket/*"
    }
  ]
}

-> bucket itself || arn:aws:s3:::my-bucket
-> object itself || arn:aws:s3:::my-bucket/*


-> Common bucket use cases

-> Public website hosting 
-> allow everyone to read files

-> Restrict access by IP
-> only office network allowed

-> allow only cloudfront 
-> secure cdn setup


-> Enforce https
"aws:SecureTransport": "false"

-> Conditions
-> add extra rules
Allow only from specific IP
Allow only during HTTPS
Allow only from CloudFront


{
  "Effect": "Deny",
  "Principal": "*",
  "Action": "s3:*",
  "Resource": [
    "arn:aws:s3:::my-bucket",
    "arn:aws:s3:::my-bucket/*"
  ],
  "Condition": {
    "Bool": {
      "aws:SecureTransport": "false"
    }
  }
}


Reject non-HTTPS requests

-> allow only cloudfront
-> suppose if bucket is public 
-> then user bypass cdn
-> this is bad -> because user bypass cdn, no caching optmization, no protection, bandwidth cost increase, harder to control access
-> Cloudfront is aws cdn
-> it's job is cache files near user
-> faster delivery 
-> reduce latency
-> Security benefits
-> instead of 
-> User -> Cloudfront
-> User -> Cloudfront -> S3

-> Cloudfront helps
-> cache files globally
-> serves faster 
-> protects s3
-> can block countries/ips
-> support signed url
-> reduced s3 load

-> Cross account access
-> Let us discuss with example
| Team       | AWS Account |
| ---------- | ----------- |
| Production | Account A   |
| Analytics  | Account B   |
| Security   | Account C   |

-> Now production owns logs
-> suppose analytics require logs
-> they donot have access
-> Cross-account access
{
  "Effect": "Allow",
  "Principal": {
    "AWS": "arn:aws:iam::222222222222:root"
  },
  "Action": "s3:GetObject",
  "Resource": "arn:aws:s3:::prod-bucket/*"
}
Allow another AWS account to read files
-> why company use different aws account 
-> security isolation, billing separation, team separation, production safety


-> Most important rule
-> explicit deny always win

-> so bucket policy is like rules written at building entrance


-> IAM 
-> identity and access management
-> what can do what on which aws resource
-> so basically it is security system of aws


| AWS IAM Concept | Real World               |
| --------------- | ------------------------ |
| IAM User        | Employee                 |
| Policy          | Permission rules         |
| Role            | Temporary responsibility |
| Resource        | Company room/files       |
-> example 
Developer can enter server room
Intern cannot
HR can access payroll

Core IAM Components

There are mainly 4 important things:

IAM Users
IAM Policies
IAM Roles
IAM Groups

-> IAM User
-> a person or application
-> each user can have passowrd, access key, permission 

-> access key 
-> programatic access -> access key, secret key id
-> used by backend, cli, sdk
-> IAM policy
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::my-bucket/*"
    }
  ]
}


-> main policy parts
| Field     | Meaning            |
| --------- | ------------------ |
| Effect    | Allow/Deny         |
| Action    | What operation     |
| Resource  | Which AWS resource |
| Condition | Extra rules        |

-> Wildcards
-> "s3:*"
-> all s3 

-> resource arn
-> aws uniquely identifies resource using resource arn
arn:aws:s3:::my-bucket/*



-> IAM roles
-> temporary permission assumed by services/users
-> instead of hardcoding aws keys 
-> aws services temporarily gets permission 
-> example ec2 needs s3
-> bad way 
-> store secret key in ec2
-> good way
-> attach IAM role to ec2

-> why these are better
-> no hardcoding, temporary

-> IAM Groups
-> collection of users
-> attach policy to group instead of each user
-> these are the permission
1. Explicit DENY
2. Explicit ALLOW
3. Default DENY



IAM vs Bucket Policy
| Real World     | AWS             |
| -------------- | --------------- |
| Ticket         | IAM Policy      |
| Security guard | Bucket Policy   |
| Stadium        | S3 Bucket       |
| Customer       | User/Role       |
| Rules at gate  | Resource policy |

-> iam attached to identity
-> bucket policy is attached to resource
-> both work together

-> Bucket Versioning
-> it actually give us more safety with data
-> suppose you uploaded resume.pdf
-> then again you upload resume.pdf
-> now it will store like this
-> resume.pdf
-> resume.pdf
-> each will be stored with version
-> if not versioning then old one will get deleted
-> for delete
-> delete-marker
-> it will not delete just add a marker because of which objects will be hidden not deleted
-> now versioning can be
-> disabled
-> enabled
-> suspended -> it is suppose you enabled now the rule is you cannot diable only suspend 
-> in suspend it will have versioning for old object but new behaviour limited
hello.txt
 ├── V1
 └── V2

hello.txt
 ├── V1
 ├── V2
 └── Delete Marker


Permanent Delete
-> you must delete individual versions  

Delete Version V1
Delete Version V2
Delete Delete-Marker

object fully gone

-> without version id it will temporary delete
-> with version id it will permanent delete

aws s3api delete-object \
  --bucket my-bucket \
  --key hello.txt \
  --version-id VERSION_ID



Encryption type
Info
Secure your objects with two separate layers of encryption. For details on pricing, see DSSE-KMS pricing on the Storage tab of the Amazon S3 pricing page. 
Server-side encryption with Amazon S3 managed keys (SSE-S3)
-> aws just do it we donot have any control
Server-side encryption with AWS Key Management Service keys (SSE-KMS)
-> encryption keys are taken from KMS
Dual-layer server-side encryption with AWS Key Management Service keys (DSSE-KMS)  
-> dual layer of encryption 
-> it happens 2 ties

| Type     | Key Managed By   | Complexity | Typical Use           |
| -------- | ---------------- | ---------- | --------------------- |
| SSE-S3   | S3               | Simple     | Normal apps           |
| SSE-KMS  | KMS              | Medium     | Production enterprise |
| DSSE-KMS | KMS (dual layer) | Advanced   | High compliance       |


-> SSE-KMS is very common


Bucket Key
Using an S3 Bucket Key for SSE-KMS reduces encryption costs by lowering calls to AWS KMS. S3 Bucket Keys aren't supported for DSSE-KMS. Learn more 
Disable
Enable
-> for SSE-KMS enable helps to lower the cost
-> this will help lower the cost for KMS(Key managemen system) 


Advanced settings
Object Lock
Store objects using a write-once-read-many (WORM) model to help you prevent objects from being deleted or overwritten for a fixed amount of time or indefinitely. Object Lock works only in versioned buckets. Learn more 
Disable
Enable
Permanently allows objects in this bucket to be locked. Additional Object Lock configuration is required in bucket details after bucket creation to protect objects in this bucket from being deleted or overwritten.
-> this will make restrict to delete features


"Version": "2012-10-17",
-> it is aws policy version


Bucket-level permissions
| Permission                      | What it allows           |
| ------------------------------- | ------------------------ |
| `s3:ListBucket`                 | View files inside bucket |
| `s3:GetBucketLocation`          | Know region              |
| `s3:ListBucketMultipartUploads` | View multipart uploads   |
| `s3:GetBucketPolicy`            | Read bucket policy       |
| `s3:PutBucketPolicy`            | Modify bucket policy     |

Object-level permissions
| Permission        | Meaning         |
| ----------------- | --------------- |
| `s3:GetObject`    | Download object |
| `s3:PutObject`    | Upload object   |
| `s3:DeleteObject` | Delete object   |

{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::amzn-s3-beats-ingestion-bucket"
    },
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::amzn-s3-beats-ingestion-bucket/*"
    }
  ]
}

-> together

->cloudshell

[default]
aws_access_key_id=AKIA....
aws_secret_access_key=....


-> how to connect your server to s3

1. Install aws cli in your machine
2. use .env file and store credentails
3. specify credentials in ec2 itself


-> delete bucket 
aws s3 rb s3://amzn-s3-beats-ingestion-svc-bucket --force


-> administer access to iam user
aws iam attach-user-policy \
    --user-name aakashloyar \
    --policy-arn arn:aws:iam::aws:policy/AdministratorAccess



-> How AWS SDK Gets Credentials   
you donot need this 
config.WithCredentialsProvider(...)
because the AWS SDK automatically reads variables from .env file
awsCfg, err := config.LoadDefaultConfig(
	ctx,
	config.WithRegion(cfg.Region),
)

Environment Variables
      ↓
AWS SDK
      ↓
Authenticated Client


-> Postgresql on aws
-> aws aurora and RDS
-> so aurora(Postgresql compatibility) vs PostgreSQL 
-> use PostgreSQL 
-> if scaling to millions then use aurora 
-> aurora provides more features 
