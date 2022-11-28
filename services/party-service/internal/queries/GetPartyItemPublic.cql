MATCH
  (user:User)-[:POSTED]->(party:Party {id: $partyID})
OPTIONAL MATCH
  (:User)-[interest:INTERESTED]->(party)
OPTIONAL MATCH
  (:User)-[going:GOING]->(party)
OPTIONAL MATCH
  (:User)-[invited:INVITED]->(party)
WITH
  party,
  user,
  count(DISTINCT going) AS going_count,
  count(DISTINCT interest) AS interested_count,
  count(DISTINCT invited) AS invited_count
RETURN
  {
    party:            properties(party),
    user:             properties(user),
    going_count:      going_count,
    interested_count: interested_count,
    invited_count:    invited_count
  }
  ORDER BY party.createTime DESC
