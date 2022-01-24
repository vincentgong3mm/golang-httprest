### mongodb 기본 명령어 입니다.

[참고링크](https://velopert.com/mongodb-tutorial-list)

| 명령어 | 설명 |
| ------ | ----------- |
| show dbs | db 목록 조회 |
| use `DB_NAME` | `DB_NAME`사용하도록 변경,  `DB_NAME`이 없으면 생성함 |
| db.`COLLECTION_NAME`.find()  | `COLLECTION_NAME`에 있는 데이터(Document)모두 조회|
| db.`COLLECTION_NAME`.insert({`name`:`my name`}) | `COLLECTION_NAME`에 `name` field에 `my name` 데이터를 추가 |
| db.`COLLECTION_NAME`.update({`name`:`my name`}, {`name`:`my name`, `email`: `my email`}) | Document 전체 변경, `COLLECTION_NAME`에 `name` field가 `my name` 데이터의 Document를 변경 |
| db.`COLLECTION_NAME`.update({`name`:`my name`}, {$set : {`email`: `my email2`}}) | Document의 특정 field변경, `COLLECTION_NAME`에 `name` field가 `my name` 데이터의 `email` field를 `my email2`로 변경 |
| db.`COLLECTION_NAME`.remove({})| 모든 Document 삭제|
| db.`COLLECTION_NAME`.find({}, {`email`:true}) | 모든 Document 검색 후 `email` field만 표시|
| db.`COLLECTION_NAME`.createIndex( { email: 1 }, { unique: true } )|`email` field로 unique index 생성|
| db.`COLLECTION_NAME`.createIndex( { firstName: 1, lastName: 1 }, { unique: true } )|`firstName` + `lastName` field로 unique index 생성|

### Collection Schema 확인하기
- 아래는 `avatar` Collection의 Schema확인 하는 방법입니다.
```
var obj = db.avatar.findOne()
for (var key in obj) {print(key, typeof obj[key]);}
_id object
user_name string
email string
```
참고 : [How to find Schema of a Collection in MongoDB](https://medium.com/@ahsan.ayaz/how-to-find-schema-of-a-collection-in-mongodb-d9a91839d992)


### Seaside DB설명
#### seaside_user DB
> user Collection
> ```
> _id object
> user_name string
> email string
> ```

#### seaside_avatar DB
> avatar Collection
> ```
> _id object
> avatar_name string
> avatar_level string
> equipment string
> user_name string
> ````

> chat Collection
> ```
> _id object
> room_name string
> max_user number
> attendant object
> ```

> dict_item Collection
> ```
> _id object
> item_type string
> item_name string
> ```

> inventory Collection
> ```
> _id object
> item_type string
> item_name string
> avatar_name string
> ```
