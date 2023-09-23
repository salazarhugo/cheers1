OPTIONAL MATCH
  (me:User {id: $userID})
MATCH
  (party:Party)
WHERE
  toFloat(party.startDate) > timestamp() / 1000
OPTIONAL MATCH
  (creator:User)-[:POSTED]->(party)
OPTIONAL MATCH
  (me)-[:FRIENDS]->(friend_going:User)-[w:WATCH_STATUS]->(party) WHERE w.value = "GOING"
OPTIONAL MATCH
  (me)-[:FRIENDS]->(friend_interested:User)-[w:WATCH_STATUS]->(party) WHERE w.value = "INTERESTED"
OPTIONAL MATCH
  (me)-[watchStatus:WATCH_STATUS]->(party)
WITH
  creator,
  party,
  me.id IS NOT NULL AND me.id = creator.id AS is_creator,
  count(DISTINCT watchStatus) AS interested_count,
  count(DISTINCT watchStatus) AS going_count,
  count(DISTINCT watchStatus) AS invited_count,
  watchStatus.value AS watch_status,
  collect(DISTINCT properties(friend_going)) AS mutual_going,
  collect(DISTINCT properties(friend_interested)) AS mutual_interested
ORDER BY party.startDate
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
    viewer_watch_status: watch_status,
    mutual_going: mutual_going,
    mutual_interested: mutual_interested
  }
