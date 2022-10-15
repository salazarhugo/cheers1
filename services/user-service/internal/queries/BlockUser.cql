MATCH
  (u:User)-[r:FOLLOWS]-(u2:User)
WHERE
  u.id = $userID AND u2.id = $otherUserID
DELETE
  r
MERGE
  (u)-[:BLOCKED]->(u2)
