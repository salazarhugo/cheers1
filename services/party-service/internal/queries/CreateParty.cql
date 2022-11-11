MATCH
  (u:User { id: $userID})
CREATE
  (party: Party $party)
CREATE
  (u)-[:POSTED]->(party)