## How to Write Go Code

- go 프로그램들 package들로 구성되어있습니다.
- 패키지는 같은 디렉토리에서 함께 컴파일된 소스코드의 묶음입니다. 한 소스파일에 정의된 Functions, types, variables, constants은 같은 패키지 내에 다른 소스 파일에 표시됩니다.
- 하나의 repository는 모듈들을 포함하며, 한 모듈은 함께 릴리즈된 서로 관련있는 Go 패키지들의 묶음입니다.
- Go 저장소는 일반적으로 저장소의 루트에 위치한 하나의 모듈만 포함한다.

- 아래와 같이 go env 명령을 통해 go 환경변수 기본 값을 설정할 수 있습니다.
    ```sh
    # 환경변수 값 설정
    $ go env -w GOBIN=/somewhere/else/bin
    # 환경변수 값 해제
    $ go env -u GOBIN
    ```

- 사용자가 작성하려는 go 프로젝트의 폴더 경로를 go.mod를 통해 등록한 후, import 시 패키지에 추가적으로 경로를 붙여 사용합니다.
