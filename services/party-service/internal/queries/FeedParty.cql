MATCH
  (me:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(party:Party)
OPTIONAL MATCH
  (:User)-[interest:INTERESTED]->(party)
OPTIONAL MATCH
  (:User)-[going:GOING]->(party)
OPTIONAL MATCH
  (:User)-[invite:INVITED]->(party)
WITH
  me,
  creator,
  party,
  count(DISTINCT interest) AS interested_count,
  count(DISTINCT going) AS going_count,
  count(DISTINCT invite) AS invited_count
SKIP
  $skip
LIMIT
  $pageSize
RETURN
  {
    party:            PROPERTIES(party),
    user:             PROPERTIES(creator),
    going_count:      going_count,
    interested_count: interested_count,
    invited_count:    invited_count,
    is_creator:       creator.id = me.id
  }
  ORDER BY party.createTime DESC
