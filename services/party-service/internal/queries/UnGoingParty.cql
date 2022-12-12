MATCH
  (u:User {id: $userID})-[r:GOING]->(party:Party {id: $partyID})
DELETE r