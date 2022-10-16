MATCH
  (me:User {id: $userId})
MATCH
  (u:User { id: $otherUserID})-[:FOLLOWS]->(f:User)
OPTIONAL MATCH
  (f)-[:POSTED]->(s:Story) WHERE s.created > datetime().epochMillis-24*60*60*1000
OPTIONAL MATCH
  (me)-[seen:SEEN]->(s)
WITH
  f,
  EXISTS((u)-[:FOLLOWS]->(f)) as has_followed,
  COUNT(s) > 0 as has_story,
  COUNT(seen) = COUNT(s) as seen_story
  LIMIT 20
RETURN
  f {
    .id,
    .name,
    .username,
    .verified,
    .picture,
    has_followed: has_followed,
    storyState:
      CASE
        WHEN has_story AND NOT seen_story THEN "NOT_SEEN"
        WHEN has_story AND seen_story THEN "SEEN"
        ELSE "EMPTY"
      END
  }
