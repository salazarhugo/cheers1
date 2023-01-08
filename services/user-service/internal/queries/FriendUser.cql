MATCH
  (u:User {id: $userID})
MATCH
  (f:User)
WHERE f.id = $otherUser OR f.username = $otherUser
MERGE
  (u)-[:FRIENDS]->(f)
MERGE
  (u)<-[:FRIENDS]-(f)
