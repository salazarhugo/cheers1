MATCH
  (me:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(party:Party)
WHERE
  toFloat(party.startDate) > timestamp() / 1000
OPTIONAL MATCH
  (me)-[watchStatus:WATCH_STATUS]->(party)
OPTIONAL MATCH
  (:User)-[interest:INTERESTED]->(party)
OPTIONAL MATCH
  (:User)-[going:GOING]->(party)
OPTIONAL MATCH
  (:User)-[invite:INVITED]->(party)
WITH
  creator,
  party,
  me.id = creator.id AS is_creator,
  count(DISTINCT interest) AS interested_count,
  count(DISTINCT going) AS going_count,
  count(DISTINCT invite) AS invited_count
  SKIP
  $skip
  LIMIT
  $pageSize
RETURN
  {
    party:            properties(party),
    user:             properties(creator),
    going_count:      going_count,
    interested_count: interested_count,
    invited_count:    invited_count,
    is_creator:       is_creator,
    viewer_watch_status: watchStatus.value
  }
  ORDER BY party.createTime DESC
