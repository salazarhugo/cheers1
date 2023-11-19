MATCH
  (u:User {id: $userID})
MATCH
  (f:User {id: $otherUserID})
MERGE
  (u)-[:FRIENDS]->(f)
MERGE
  (u)<-[:FRIENDS]-(f)
