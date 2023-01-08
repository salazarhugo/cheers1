MATCH
  (u:User {id: $userID})
MATCH
  (f:User)
WHERE f.id = $otherUser OR f.username = $otherUser
MATCH (u)-[r:REQUESTED]->(f)
DELETE r
