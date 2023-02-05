MATCH
  (receiver:User {id: $userID})
MATCH
  (creator:User)-[r:POSTED]->(party: Party {id: $partyID})
DELETE
  r
CREATE
  (receiver)-[:POSTED]->(party)
SET party.hostId = receiver.id
