# google calendar의 정보를 저장할 mongodb의 db, collection, document입니다.

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


