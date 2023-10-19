MATCH
  (receiver:User {id: $userID})
MATCH
  (party: Party {id: $partyID})
OPTIONAL MATCH
  (creator:User)-[r:POSTED]->(party)
DELETE
  r
CREATE
  (receiver)-[:POSTED]->(party)
SET party.hostId = receiver.id
