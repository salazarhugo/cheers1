MATCH
  (u:User)
WHERE
  u.username CONTAINS $query OR toLower(u.name) CONTAINS $query AND u.id <> $userID
WITH
  u
LIMIT 10
OPTIONAL MATCH
  (u)-[posts:POSTED]->(:Post)
OPTIONAL MATCH
  (u)-[following:FOLLOWS]->(:User)
OPTIONAL MATCH
  (:User)-[followers:FOLLOWS]->(u)
WITH
  u,
  EXISTS((:User {id: $userID})-[:FOLLOWS]->(u)) as followBack
RETURN
  PROPERTIES(u)
