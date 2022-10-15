MATCH
  (me:User {id: $userID})
MATCH (u:User)
   WHERE u.id = $otherUserID OR u.username = $otherUserID
OPTIONAL MATCH
  (u)-[authorPosts:POSTED]->(:Post)
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
  COUNT(s) > 0 as hasStory,
  COUNT(seen) = count(s) as seenStory,
  EXISTS((me)-[:FOLLOWS]->(u)) as has_followed
RETURN
   {
      user: u,
      post_count: post_count,
      followers_count: followers_count,
      following_count: following_count,
      has_followed : has_followed,
      story_state:
       CASE
         WHEN hasStory AND NOT seenStory THEN "NOT_SEEN"
         WHEN hasStory AND seenStory THEN "SEEN"
         ELSE "EMPTY"
         END
   }