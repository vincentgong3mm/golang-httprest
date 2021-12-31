
-- 테스트의 다양화 하기 위해서 Database를 두 개 이상 구성합니다.

DataBase : seaside_user
	* 유저의 정보
	* 테스트를 위해서 분리했고, Collection을 하나만 만듭니다.
	Collection : user 
		user_name
		email
		cur_avatar

DataBase : seaside_avatar
	* 유저는 n개의 avatar를 가질 수 있습니다.
	* avatar는 n개의 인벤토리를 가질 수 있습니다.

	Collection : avatar
		user_name
		avatar_name
		avatar_level
		equipment
			hat
			coat
			pants
			shoes		

	Collection : inventory
		avatar_name
		item_type {hat | coat | pants | shoes}
		item_name

	Collection : dict_item
		item_type
		item_name
	
	Collection : chat	- 채팅방 정보
		room_name
		max_user
		attendant {user_name, avatar_name +}
	Collection : history_chat - 채팅 히스토리
		room_name
		seq
		message {"aaaa" | "bbb"}





