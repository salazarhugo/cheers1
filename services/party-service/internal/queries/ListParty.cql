MATCH
  (viewer:User {id: $viewerID})
MATCH
  (creator:User {id: $userID})-[:POSTED]->(party:Party)
OPTIONAL MATCH
  (viewer)-[watchStatus:WATCH_STATUS]->(party)
OPTIONAL MATCH
  (:User)-[interest:INTERESTED]->(party)
OPTIONAL MATCH
  (:User)-[going:GOING]->(party)
OPTIONAL MATCH
  (:User)-[invite:INVITED]->(party)
WITH
  creator,
  party,
  viewer.id = creator.id AS is_creator,
  count(DISTINCT interest) AS interested_count,
  count(DISTINCT going) AS going_count,
  count(DISTINCT invite) AS invited_count,
  watchStatus.value AS watch_status
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
    viewer_watch_status: watch_status
  }
  ORDER BY party.createTime DESC
