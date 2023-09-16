OPTIONAL MATCH
  (me:User {id: $userID})
MATCH
  (party:Party)
WHERE
  toFloat(party.startDate) > timestamp() / 1000
OPTIONAL MATCH
  (creator:User)-[:POSTED]->(party)
OPTIONAL MATCH
  (me)-[:FRIENDS]->(friend:User)-[w:WATCH_STATUS]->(party) WHERE w.value = "GOING"
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
  collect(friend.username) AS mutual_usernames,
  collect(friend.picture) AS mutual_pictures
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
    mutual_usernames: mutual_usernames,
    mutual_pictures: mutual_pictures
  }
