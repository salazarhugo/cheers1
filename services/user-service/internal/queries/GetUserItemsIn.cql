MATCH
  (me:User {id: $userID})
MATCH
  (u:User) WHERE u.id IN $userIDs
OPTIONAL MATCH
  (u)-[:POSTED]->(s:Story) WHERE s.created > datetime().epochMillis-24*60*60*1000
OPTIONAL MATCH
  (me)-[seen:SEEN]->(s)
WITH
  u,
  EXISTS((me)-[:FOLLOWS]->(u)) as has_followed,
  EXISTS((me)-[:FRIENDS]->(u)) as friend,
  EXISTS((me)-[:REQUESTED]->(u)) as requested,
  COUNT(s) > 0 as has_story,
  COUNT(seen) = COUNT(s) as seen_story
  LIMIT 20
RETURN
  u {
    .id,
    .name,
    .username,
    .verified,
    .picture,
    has_followed: has_followed,
    friend: friend,
    requested: requested,
    storyState:
    CASE
      WHEN has_story AND NOT seen_story THEN "NOT_SEEN"
      WHEN has_story AND seen_story THEN "SEEN"
      ELSE "EMPTY"
    END
  }
