MATCH
  (u:User)
WHERE
  u.username CONTAINS $query OR toLower(u.name) CONTAINS $query AND u.id <> $userID
WITH
  u
LIMIT 10
OPTIONAL MATCH
  (u)-[posts:POSTED]->(:Post)
OPTIONAL MATCH
  (u)-[friends:FRIENDS]->(:User)
OPTIONAL MATCH
  (u)-[following:FOLLOWS]->(:User)
OPTIONAL MATCH
  (:User)-[followers:FOLLOWS]->(u)
OPTIONAL MATCH
  (u)-[:POSTED]->(s:Story) WHERE s.created > datetime().epochMillis-24*60*60*1000
OPTIONAL MATCH
  (me)-[seen:SEEN]->(s)
WITH
  u,
  EXISTS((:User {id: $userID})-[:FOLLOWS]->(u)) as has_followed,
  COUNT(s) > 0 as has_story,
  COUNT(DISTINCT friends) as friends_count,
  EXISTS((me)-[:FRIENDS]->(u)) as friend,
  EXISTS((me)-[:REQUESTED]->(u)) as requested,
  EXISTS((u)-[:REQUESTED]->(me)) as has_requested_viewer,
  COUNT(seen) = count(s) as seen_story
RETURN DISTINCT
  {
    id: u.id,
    name: u.name,
    username: u.username,
    verified: u.verified,
    picture: u.picture,
    has_followed: has_followed,
    friend: friend,
    requested: requested,
    has_requested_viewer: has_requested_viewer,
    friends_count: friends_count,
    story_state:
      CASE
        WHEN has_story AND NOT seen_story THEN "NOT_SEEN"
        WHEN has_story AND seen_story THEN "SEEN"
        ELSE "EMPTY"
      END
  }
