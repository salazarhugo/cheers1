MATCH
  (me:User {id: $userID})
MATCH
  (user:User)-[:POSTED]->(party:Party {id: $partyID})
OPTIONAL MATCH
  (:User)-[interest:INTERESTED]->(party)
OPTIONAL MATCH
  (:User)-[going:GOING]->(party)
WITH
  party,
  user,
  me.id = user.id AS is_creator,
  exists((me)-[:GOING]->(party)) AS is_going,
  count(DISTINCT interest) AS interested_count,
  count(DISTINCT going) AS going_count
RETURN
  {
    party:         properties(party),
    user:          properties(user),
    user_response: 'GOING',
    is_creator:    is_creator
  }
  ORDER BY party.createTime DESC
