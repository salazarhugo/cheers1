MATCH
  (me:User {id: $userID})
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
  me.id = user.id AS is_creator,
  exists((me)-[:GOING]->(party)) AS is_going,
  count(DISTINCT going) AS going_count,
  count(DISTINCT interest) AS interested_count,
  count(DISTINCT invited) AS invited_count
RETURN
  {
    party:            properties(party),
    user:             properties(user),
    going_count:      going_count,
    interested_count: interested_count,
    invited_count:    invited_count,
    is_creator:       is_creator,
    user_response:    'GOING'
  }
  ORDER BY party.createTime DESC
