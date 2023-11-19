MATCH
  (u:User {id: $userID})
MATCH
  (f:User {id: $otherUserID})
MATCH (u)-[r:FRIENDS]->(f)
MATCH (f)-[r2:FRIENDS]->(u)
DELETE r
DELETE r2
