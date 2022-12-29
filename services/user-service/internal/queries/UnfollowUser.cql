MATCH
  (u:User {id: $userID})-[r:FOLLOWS]->(u2:User)
WHERE
  u2.id = $otherUser OR u2.username = $otherUser
DELETE
  r