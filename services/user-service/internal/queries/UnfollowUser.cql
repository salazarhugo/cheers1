MATCH
  (u:User)-[r:FOLLOWS]->(u2:User)
WHERE
  u.id = $userID AND u2.username = $otherUserID
DELETE
  r