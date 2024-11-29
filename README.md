# mercury

## Tools to install
go install \  
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \  
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \  
    google.golang.org/protobuf/cmd/protoc-gen-go \  
    google.golang.org/grpc/cmd/protoc-gen-go-grpc  

## Services
- bff 8080  
- user 8091  
- article 8092  
- interactive 8093  
- comment 8094  
- captcha 8095  
- sms 8096  
- oauth2 8097  
- ranking 8098  
- crontask 8099  
- follow 9000  
- payment 9001  
- account 9002  

## Todo
1. [article] userrpc improve

## Startup sequence
1. user, sms, interactive, comment, ~~oauth2~~, ~~crontask~~  
2. article, captcha
3. ranking
4. bff