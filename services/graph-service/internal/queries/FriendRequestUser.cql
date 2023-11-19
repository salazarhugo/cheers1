MATCH
  (u:User {id: $userID})
MATCH
  (f:User {id: $otherUserID})
MERGE
  (u)-[:REQUESTED]->(f)
