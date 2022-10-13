MATCH
  (u: User{id:$userId})
OPTIONAL MATCH
  (e1:Party {privacy: "PUBLIC"})
OPTIONAL MATCH
  (u)-[:FOLLOWS]->(author:User)-[:POSTED]->(e2:Party {privacy: "FRIENDS"})
OPTIONAL MATCH
  (e3:Party)-[:INVITED]->(u)
UNWIND
  [val in [e1, e2, e3] WHERE val is not null] as e
OPTIONAL MATCH
  (:User)-[interest:INTERESTED]->(e)
OPTIONAL MATCH
  (:User)-[going:GOING]->(e)
WITH
  u,
  e,
  author,
  exists((u)-[:INTERESTED]->(e)) as interested,
  count(DISTINCT interest) as interestedCount,
  exists((u)-[:GOING]->(e)) as going,
  count(DISTINCT going) as goingCount
  SKIP
  $skip
  LIMIT
  $pageSize
CALL {
  WITH u, e, author
  OPTIONAL MATCH (u)-[:FOLLOWS]->(mutual:User)-[:GOING]->(e)
  WITH mutual, count(mutual) as mutualCount LIMIT 2
  RETURN collect(mutual.picture) as mutualpictures,
         collect(mutual.username) as mutualUsernames, mutualCount
}
RETURN e {
       .*,
         hostId: author.id,
         hostName: author.name,
         interested: interested,
         interestedCount: interestedCount,
         going: going,
         goingCount: goingCount,
         mutualCount: mutualCount,
         mutualUsernames: mutualUsernames,
         mutualpictures: mutualpictures
       }
ORDER BY e.startDate DESC
