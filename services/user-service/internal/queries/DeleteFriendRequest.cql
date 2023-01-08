MATCH
  (u:User {id: $userID})
MATCH
  (f:User {id: $otherUserID})
MATCH (u)-[r:REQUESTED]->(f)
DELETE r
