MATCH
  (me:User {id: $userID})
MATCH (u:User)
   WHERE u.id = $otherUserID OR u.username = $otherUserID
OPTIONAL MATCH
  (u)-[authorPosts:POSTED]->(:Post)
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
  COUNT(DISTINCT authorPosts) as post_count,
  COUNT(DISTINCT followers) as followers_count,
  COUNT(DISTINCT following) as following_count,
  COUNT(DISTINCT friends) as friends_count,
  COUNT(s) > 0 as hasStory,
  COUNT(seen) = count(s) as seenStory,
  EXISTS((me)-[:FOLLOWS]->(u)) as has_followed,
  EXISTS((me)-[:FRIENDS]->(u)) as friend,
  EXISTS((me)-[:REQUESTED]->(u)) as requested,
  EXISTS((u)-[:REQUESTED]->(me)) as has_requested_viewer
RETURN
   {
      user: PROPERTIES(u),
      post_count: post_count,
      followers_count: followers_count,
      following_count: following_count,
      friends_count: friends_count,
      has_followed : has_followed,
      friend : friend,
      requested : requested,
      has_requested_viewer : has_requested_viewer,
      story_state:
       CASE
         WHEN hasStory AND NOT seenStory THEN "NOT_SEEN"
         WHEN hasStory AND seenStory THEN "SEEN"
         ELSE "EMPTY"
       END
   }