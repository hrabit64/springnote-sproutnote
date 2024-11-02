# Sproutnote

`Springnote` 블로그 프로젝트를 위한 간단한 도커 기반의 백업 애플리케이션입니다.

파일 및 Mysql(or Mariadb) 데이터베이스에 대한 간단한 백업을 수행할 수 있습니다.



## 0. 사용법

``` bash
docker build -t your-img-name:tag .
```



#### docker-compost.yml

``` yaml
sproutnote:
    image: sproutnote:latest
    restart: always
    container_name: sproutnote
    volumes:
      - sproutnote_db:/app/data.db
      - sproutnote_backup:/app/backup
```



## 1. 스케줄러

`Sproutnote` 는 스케줄러를 통해 정해진 시간에 파일 아이템 및 데이터베이스 아이템 전체를 백업합니다.

백업 수행 후,  `/app/backup`  경로에 ` [db or file]_[item display_name]_[run date]_[uuid]` 형식의 디렉토리를 생성한 후, 해당 디렉토리 내부에 백업 파일이 생성되게 됩니다.

### 파일 백업

디렉토리나 파일을 지정할 수 있으며, 백업 대상 디렉토리 내부에 파일을 복사합니다.



### 데이터베이스 백업

mysqldump를 사용하여 백업을 실행하며, 백업 디렉토리에 backup.sql 파일 형식으로 dump를 실행합니다.

이때 아래 플래그를 사용합니다.

``` bash
--single-transaction
--skip-opt
--extended-insert
--add-drop-database
--add-drop-table
--no-create-db
--no-create-info
--default-character-set=utf8
--quick
```



## 2. Sproutnote CLI

`Sproutnote cli` 를 통해 데이터베이스 아이템을 관리하고, 파일 아이템을 관리할 수 있습니다.

```shell
docker exec -it [container_name] /bin/sproutnote
```

### Commands

#### 1. 데이터베이스 아이템 관리

`db add` : 데이터베이스에 아이템 추가

**!주의! 커넥션 정보가 평문으로 저장됩니다. 유출에 주의하세요.**

``` bash
>>> db add
1. Enter Display Database Name. (1 to 50 characters, for only display. not connection url) : 
Display Name: test
2. Enter Database URL (do not include port) : 
Database URL: testurl
3. Enter Database Port : 
Database Port: 1234
4. Enter Database Account ID : 
Account ID: some
5. Enter Database Account Password : 
Password: ****
6. Enter Target Database Name : 
Target Database Name: some-database
7. Confirm the entered information.
Database Name :  test
Database URL :  testurl
Database Port :  1234
Database Account ID :  some
Database Target DB :  some-database
Is the information correct? (y/n)
✔ Yes
```

`db show` : 데이터베이스 아이템 조회

``` bash
>>> db show
<Database items> page1/1
│────│────────────│────────────────────│────────────│───────────────────────│
│ ID │ NAME       │ URL                │ ACCOUNT ID │ TARGET DB             │
│────│────────────│────────────────────│────────────│───────────────────────│
│ 1  │ Sproutnote │ some-conn-url:3306 │  some-acc  │    some-target-db     │
│────│────────────│────────────────────│────────────│───────────────────────│
```

`db remove [id]` : 데이터베이스 아이템 삭제

`db dump [id:optional]` : 데이터베이스 백업 실행, id 가 없으면 전체 백업



#### 2. 파일 아이템 관리

`file add` : 파일 아이템 추가

``` bash
>>> file add
1. Enter the file item display name. (1 to 50 characters) : 
Display Name: some
2. Enter the file item path : 
Path: /test/path
3. Confirm the file item information.
Name : some
Path : /test/path
Do you want to add this file item? (y/n)
✔ Yes

```

`file show` : 파일 아이템 조회

``` bash
>>> file show
<File items> page1/1
│────│──────│────────────│
│ ID │ NAME │ PATH       │
│────│──────│────────────│
│ 1  │ loki │ /data/loki │
│────│──────│────────────│
```

`file remove [id]` : 파일 아이템 삭제

`file backup [id:optional]` : 파일 백업 실행, id 가 없으면 전체 백업



#### 3. 히스토리 확인

`history [type:db or file] [item id]` : 해당 아이템의 히스토리 조회



#### 4. 설정 관리

`config show` : 설정 조회

```bash
>>> config show
│─────────────────────────│──────────│
│ KEY (5)                 │ VALUE    │
│─────────────────────────│──────────│
│ BACKUP_PATH             │ ./backup │
│ FILE_BACKUP_TIME        │ 01:00    │
│ MAX_FILE_BACKUP_HISTORY │ 14       │
│ DB_BACKUP_TIME          │ 01:30    │
│ MAX_DB_BACKUP_HISTORY   │ 14       │
│─────────────────────────│──────────│

```

`config edit [key] [value]` : 설정 수정

| KEY                     |                                                              |
| ----------------------- | ------------------------------------------------------------ |
| db_backup_time          | 데이터베이스 아이템 백업을 실행하는 시간입니다. HH:MM 형식입니다. |
| max_db_backup_history   | 저장할 데이터베이스의 최대 히스토리 갯수입니다. 각 아이템별로 개별로 관리되며, 해당 갯수를 넘어가는 백업본은 자동으로 삭제 됩니다. |
| file_backup_time        | 파일 아이템 백업을 실행하는 시간입니다. HH:MM 형식입니다.    |
| max_file_backup_history | 저장할 파일의 최대 히스토리 갯수입니다. 각 아이템별로 개별로 관리되며, 해당 갯수를 넘어가는 백업본은 자동으로 삭제 됩니다. |



