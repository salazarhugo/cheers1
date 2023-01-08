MATCH
  (u:User {id: $userID})
MATCH
  (f:User {id: $otherUserID})
MATCH (u)-[r:FRIENDS]-(f)
DELETE r
