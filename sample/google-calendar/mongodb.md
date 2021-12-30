# google calendar의 정보를 저장할 mongodb의 db, collection, document입니다.

# db
## gcal

# collection
## calendar
유저별 calendar 정보를 저장합니다. 
user_id에 google email id를 저장하고 조회하는 조건으로 사용합니다.
response 할 때는 user_id를 제외하고 전달합니다.
```json
{
	"user_id" : "my field", 
	"kind": "google",
	"etag": "\"google\"",
	"id": "google",
	"summary": "google",
	"timeZone": "google",
	"conferenceProperties": {
	 "allowedConferenceSolutionTypes": [
	  "google"
	 ]
	}
}
```
### 처리하는 google calendar api
- https://developers.google.com/calendar/api/v3/reference#CalendarList, /users/me/calendarList/ 로 시작하는 모든 api
- https://developers.google.com/calendar/api/v3/reference#Calendars, /calendars/:calendarId로 시작하는 모든 api


## event
유저의 calendar id별 events(일정)을 저장합니다.
```json
{
	"kind": "",
	"etag": "",
	"id": "",
	"status": "",
	"htmlLink": "",
	"created": "",
	"updated": "",
	"summary": "",
	"creator": {
	"email": "",
	"self": 
	},
	"organizer": {
	"email": "",
	"self": 
	},
	"start": {
	"date": ""
	},
	"end": {
	"date": ""
	},
	"transparency": "",
	"iCalUID": "",
	"sequence": ,
	"reminders": {
	"useDefault": 
	},
	"eventType": ""
}
```

### 처리하는 google calendar api
- https://developers.google.com/calendar/api/v3/reference#Events
- paging : https://developers.google.com/calendar/api/guides/pagination

# mongodb 기본 명령어 입니다.
### $ show dbs
### $ user gcalendar 
### $ show collections
### $ db.calendar.find()
모든 document 검색
### $ db.calendar.insert({kkk:"bbb"})
document 하나 추가
### $ db.calendar.remove({})
document 모두 삭제
### $ db.calendar.find({}, {summary:true})
모든 document 검색 후 summary field만 표시 
### $ db.userinfo.createIndex( { email: 1 }, { unique: true } )
email field를 unique index로 생성
### $ db.userinfo.createIndex( { firstName: 1, lastName: 1 }, { unique: true } )
firstName, lastName을 복합인덱스, 복합 unique index로 생성

### $ var obj = db.avatar.findOne()
### $ for (var key in obj) {print(key, typeof obj[key]);}
#### _id object
#### avatar_name string
#### avatar_level string
#### equipment string
#### user_name string
collection의 schema 확인하기 https://medium.com/@ahsan.ayaz/how-to-find-schema-of-a-collection-in-mongodb-d9a91839d992


