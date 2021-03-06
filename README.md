<br/>

# Highlighting

<br/>

### 팀 프로젝트는 끝났지만 아쉬운 부분들이 남는다

  * 시간과 역량의 부족으로 기술/언어 선택, 구현 방향 등에 타협했던 점
  * 나와 팀원들의 성장과 경험을 목적으로 하여 실제 프로덕트로써의 효용성과는 거리가 멀어졌던 점
  * 좋은 아이템이고 확장할 아이디어가 많은데 이른 시기에 마무리된 점 등등

<br/>

### 편집자가 아닌 일반인(팬)을 위한 서비스로 다시 만들어보자
  * 쏟아져나오는 수많은 크리에이터들의 방송영상 중 특별한 장면만 골라 볼 수 있도록
  * 원하는 장면은 클립으로 모아 유튜브, 트위치 등 플랫폼에 구분되지 않고 한데 모아 볼 수 있도록
  * 유저 백그라운드의 데이터를 서버에 모아 빠른 호출을 지원하고 다른 사람들과의 클립 공유가 가능하도록

<br/>

### 바닥부터 다시 만들면서 공부
  * 다른 기술 스택을 경험하고, 다른 구조로 디자인하며 만들어보기
  * 확장성을 고려한 선택을 해보자

<br/>

### 현재 구상
  * Desktop App : Electron, JS
  * Client View : React, Zustand
  * Client Background : JS, Python
  * API Server : Go, Nginx
  * Database : MySQL

<br/><br/>

# 개발 일지

<br/>

## 04. 13 (수)
### 02:00 AM ~ 04:00 AM
* Create React App, Electron 초기 세팅
* 초기 디자인은 메인 홈페이지, 방송 아카이브 뷰어, 클립 모음의 3 page 구성
* HashRouter - Switch - Link를 이용한 3 page 간의 link 이동 구현

<br/>

## 04. 14 (목)
### 00:00 AM ~ 03:00 AM
* 아카이브 뷰어 page 와이어프레임 디자인, 컴포넌트 설계
* sass install
* 와이어프레임에 따라 컴포넌트를 배치할 cover div style scss 작성

<br/>

## 04. 15 (금)
### 00:20 AM ~ 01:50 AM
* Server 기술 선택을 위한 언어 조사
* Go 환경 세팅 및 기초 문법 예제
* http 모듈을 이용한 localhost server 구현
* GET method

<br/>

## 04. 16 (토)
### 02:00 AM ~ 05:30 AM
* Go Modlue 공부, go mod 세팅
* route 모듈로 분리
* input data를 JSON화하는 모듈 구현
* 서버 response를 JSON으로 변경
* MySQL 설치, 예제 DB 생성

<br/>

## 04. 17 (일)
### 02:00 AM ~ 05:00 AM
* Go - MySQL 연동
* Go - MySQL 조작 연습 위한 test table 생성
* database/sql 모듈 사용법 연습
* server querystring으로 들어온 data를 insert, select하는 기능 구현
* 아카이브 뷰어 page의 data view 비즈니스 로직 디자인
* 이에 따른 DB 구조, API Server 설계
  * [txt](./archiveviewerlogic.txt)   

<br/>

