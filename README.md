# Search Jobs
 
[Search Jobs](https://jobs.yjdev.world)는 Job 웹 스크래핑을 해줍니다. [Search Jobs](https://jobs.yjdev.world)에 접속하여 찾고 싶은 job을 입력 후 검색하면 결과를 svc 파일로 다운로드 받을 수 있습니다. 현재는 jobkorea에서만 스크래핑이 가능합니다.

## Download and Install
### Install From Source
Git clone:   
```git clone https://github.com/jonggulee/search-job.git```

Docker build:   
```cd search-job```
```docker build -t searchjob:v0.1 .```

Docker run:   
```docker run -d -it -p 8080:8080 --name searchjob searchjob:v0.1```

## Install From DockerHub
Docker run:   
```docker run -d -it -p 8080:8080 --name searchjob jgjakelee/search-job:v0.1```
