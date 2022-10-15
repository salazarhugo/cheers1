MATCH
  (u:User)-[r:BLOCKED]->(u2:User)
WHERE
  u.id = $userID AND u2.username = $otherUserID
DELETE
  r