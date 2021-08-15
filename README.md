# Bread Gorris

고랭(go)로 작성된 디스코드 봇입니다.

[초대 (관리자 권한 O)](https://discord.com/api/oauth2/authorize?client_id=862671849863512064&permissions=8&scope=bot)

[초대 (관리자 권한 X)](https://discord.com/api/oauth2/authorize?client_id=862671849863512064&permissions=0&scope=bot)

## 실행

---
다음 명령어를 입력한후,

```shell
$ git clone https://github.com/Reffis/BreadGorris.git
$ cd BreadGorris
```
---
`token` 파일을 생성한후 (확장자 X), 아래의 텍스트를 입력합니다.

```
Bot [토큰]
```
`[토큰]` 에는 봇의 토큰을 입력해주세요.

---
```shell
$ go run main.go
```
---
> env 사용시, `main.go` 수정 바람. 


# 사용법

접두사: `gorris`

<br>

`gorris help`: 도움말을 표시합니다.

`gorris avatar [id or mention]`: 유저의 아바타를 가져옵니다.

`gorris choice [items]`: 랜덤으로 아이템을 고릅니다.

`gorris status [Text]`: 게임 상태를 업데이트합니다.

`gorris opensource`: 오픈소스의 정보를 표시합니다.

`gorris dev`: 개발자의 정보를 표시합니다.

> (더 추가될 예정)

<br><br>

---
> **호스팅은 안할건가요?**
> 
> * **호스팅은 생각 없습니다.** 
---
> **봇이 응답하지 않아요.**
> 
> * **명령어를 잘못 입력했거나, 권한 등이 부족할 경우 봇이 응답하지 않습니다.**
> 명령어를 제대로 입력하고, 권한도 있는데 봇이 응답하지 않을경우, 이슈를 넣어주시기 바랍니다.
---
> **소스의 일부분을 가져가실때는 반드시 원작자 표시를 해주시기 바랍니다.**
---