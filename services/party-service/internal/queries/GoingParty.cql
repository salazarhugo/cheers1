MATCH
  (u:User {id: $userID})
MATCH
  (party:Party {id: $partyID})
MERGE
  (u)-[:GOING]->(party)