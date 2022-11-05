import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {StoryState, User, UserResponse} from "../models/user.model";
import {Party} from "../models/party.model";
import {Post} from "../models/post.model";
import {Story} from "../models/story.model";
import {Ticket} from "../models/ticket.model";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";
import {user} from "@angular/fire/auth";
import {orderByKey, orderByPriority} from "@angular/fire/database";
import {FeedPostResponse, PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {Empty} from "../../../../gen/ts/google/protobuf/empty";

@Injectable({
    providedIn: 'root'
})
export class ApiService {

    BASE_URL = "https://rest-api-r3a2dr4u4a-nw.a.run.app"
    // GATEWAY_URL = "https://web-gateway-r3a2dr4u4a-nw.a.run.app"
    GATEWAY_URL = "http://localhost:8080"

    constructor(
        private http: HttpClient,
    ) {
    }

    getUserTickets(): Observable<Ticket[]> {
        return this.http.get<Ticket[]>(`${this.BASE_URL}/users/tickets`)
    }

    usernameAvailability(username: string): Observable<boolean> {
        return this.http.get<boolean>(`${this.BASE_URL}/users/available/${username}`)
    }

    register(userReq: any): Observable<User> {
        return this.http.post<User>(`${this.BASE_URL}/users/create`, userReq)
    }

    likePost(postId: string) {
        return this.http.post(`${this.BASE_URL}/posts/${postId}/like`, {})
    }

    unlikePost(postId: string) {
        return this.http.post(`${this.BASE_URL}/posts/${postId}/unlike`, {})
    }

    updateUser(user: User) {
        return this.http.post(`${this.GATEWAY_URL}/cheers.user.v1.UserService/UpdateUser`, user)
    }

    unfollowUser(username: string) {
        this.http.post(`${this.BASE_URL}/unfollow?username=${username}`, {})
    }

    followUser(username: string) {
        return this.http.post(`${this.BASE_URL}/follow?username=${username}`, {})
            .subscribe()
    }

    searchUser(query: string): Observable<User[]> {
        return this.http.get<User[]>(`${this.BASE_URL}/users/search/${query}`)
    }

    getUser(userIdOrUsername: string): Observable<User> {
        return this.http.post<UserResponse>(`${this.GATEWAY_URL}/cheers.user.v1.UserService/GetUser`, {
            id: userIdOrUsername
        }).pipe(
            map(res => {
                const user = res.user
                user.postCount = res.postCount
                user.following = res.followingCount
                user.followers = res.followersCount
                const strEnum = res.storyState as unknown as StoryState;
                user.storyState = strEnum
                return user
            })
        )
    }

    deletePost(id: string): Observable<Empty> {
        return this.http.delete(`${this.GATEWAY_URL}/v1/posts?id=${id}`)
    }

    deleteParty(id: string): Observable<any> {
        return this.http.delete(`${this.BASE_URL}/party/${id}`)
    }

    createPost(post: Post): Observable<PostResponse> {
        return this.http.post<PostResponse>(`${this.GATEWAY_URL}/v1/posts`, {
            post: {
                caption: post.caption
            }
        })
    }

    createParty(party: Party): Observable<Party> {
        party.privacy = "FRIENDS"
        return this.http.post<Party>(`${this.GATEWAY_URL}/v1/parties`, {
            party: {
                name: party.name,
                description: party.description,
                banner_url: party.bannerUrl,
                location_name: party.locationName,
                privacy: party.privacy,
            }
        })
    }

    interestParty(partyId: string): Observable<any> {
        return this.http.post(`${this.BASE_URL}/party/interest?partyId=${partyId}`, {})
    }

    uninterestParty(partyId: string): Observable<any> {
        return this.http.post(`${this.BASE_URL}/party/uninterest?partyId=${partyId}`, {})
    }

    getParty(partyId: string): Observable<Party> {
        return this.http.get<Party>(`${this.BASE_URL}/party/${partyId}`)
    }

    getMyParties(): Observable<Party[]> {
        return this.http.get<Party[]>(`${this.BASE_URL}/party/my?pageSize=10&page=0`)
    }

    getPartyFeed(): Observable<Party[]> {
        return this.http.get<Party[]>(`${this.BASE_URL}/party/feed?pageSize=10&page=0`)
    }

    getUserPosts(userId: string): Observable<PostResponse[]> {
        return this.http.get<PostResponse[]>(`${this.GATEWAY_URL}/v1/posts?userId=${userId}&pageSize=10&page=0`)
    }

    getPostFeed(): Observable<PostResponse[]> {
        return this.http.get<FeedPostResponse>(`${this.GATEWAY_URL}/v1/posts/feed?pageSize=10&page=0`).pipe(
            map(r => r.posts)
        )
    }

    getStoryFeed(): Observable<Story[]> {
        return this.http.get<Story[]>(`${this.BASE_URL}/stories/feed?pageSize=10&page=0`)
    }

    seenStory(storyId: string) {
        return this.http.post(`${this.BASE_URL}/stories/${storyId}/seen`, {})
    }
}
