MATCH
  (me:User {id: $userID})
MATCH
  (user:User)-[:POSTED]->(party:Party {id: $partyID})
OPTIONAL MATCH
  (me)-[:FRIENDS]->(friend:User)-[w:WATCH_STATUS]->(party) WHERE w.value = "GOING"
OPTIONAL MATCH
  (me)-[watchStatus:WATCH_STATUS]->(party)
WITH
  party,
  user,
  me.id = user.id AS is_creator,
  exists((me)-[:GOING]->(party)) AS is_going,
  exists((me)-[:INTERESTED]->(party)) AS is_interested,
  count(DISTINCT watchStatus) AS interested_count,
  count(DISTINCT watchStatus) AS going_count,
  count(DISTINCT watchStatus) AS invited_count,
  collect(friend.username) AS mutual_usernames,
  collect(friend.picture) AS mutual_pictures,
  watchStatus.value AS watch_status
RETURN
  {
    party:            properties(party),
    user:             properties(user),
    going_count:      going_count,
    interested_count: interested_count,
    invited_count:    invited_count,
    is_creator:       is_creator,
    viewer_watch_status: watch_status,
    mutual_usernames: mutual_usernames,
    mutual_pictures: mutual_pictures
  }
  ORDER BY party.createTime DESC
