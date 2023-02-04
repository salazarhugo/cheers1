MATCH
  (u:User {id: $userID})
MATCH
  (party:Party {id: $partyID})
MERGE
  (u)-[r:WATCH_STATUS]->(party)
SET
  r.value = $watchStatus