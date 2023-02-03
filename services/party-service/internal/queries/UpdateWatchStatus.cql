MATCH
  (u:User {id: $userID})
MATCH
  (party:Party {id: $partyID})
MERGE
  (u)-[:WATCH_STATUS { value: $watchStatus}]->(party)