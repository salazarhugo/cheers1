import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {StoryState, toUser, User, UserResponse} from "../models/user.model";
import {Post} from "../models/post.model";
import {Story} from "../models/story.model";
import {Ticket, toTicket} from "../models/ticket.model";
import {FeedPostResponse, ListPostResponse, PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {Empty} from "../../../../gen/ts/google/protobuf/empty";
import {environment} from "../../../../environments/environment";
import {CreateRegistrationTokenResponse} from "../../../../gen/ts/cheers/notification/v1/notification_service";
import {Account, GetAccountResponse} from "../../../../gen/ts/cheers/account/v1/account_service";
import {ListUserOrdersResponse, Order} from "../../../../gen/ts/cheers/order/v1/order_service";
import {ListTicketResponse} from "../../../../gen/ts/cheers/ticket/v1/ticket";

@Injectable({
    providedIn: 'root'
})
export class ApiService {

    BASE_URL = "https://rest-api-r3a2dr4u4a-nw.a.run.app"

    constructor(
        private http: HttpClient,
    ) {
    }

    listUserOrders(userId: string): Observable<Order[]> {
        return this.http.get<ListUserOrdersResponse>(`${environment.GATEWAY_URL}/v1/orders/user`)
            .pipe(map(r => r.orders))
    }

    getUserTickets(): Observable<Ticket[]> {
        return this.http.get<ListTicketResponse>(`${environment.GATEWAY_URL}/v1/tickets`)
            .pipe(map(res => res.tickets.map(ticket => toTicket(ticket))))
    }

    usernameAvailability(username: string): Observable<boolean> {
        return this.http.get<boolean>(`${this.BASE_URL}/users/available/${username}`)
    }

    register(userReq: any): Observable<User> {
        return this.http.post<User>(`${this.BASE_URL}/users/create`, userReq)
    }

    likePost(postId: string) {
        return this.http.post(`${environment.GATEWAY_URL}/v1/posts/${postId}/like`, {})
    }

    unlikePost(postId: string) {
        return this.http.post(`${environment.GATEWAY_URL}/v1/posts/${postId}/unlike`, {})
    }

    updateUser(user: User) {
        return this.http.patch(`${environment.GATEWAY_URL}/v1/users`, user)
    }

    followUser(username: string) {
        return this.http.post(`${this.BASE_URL}/follow?username=${username}`, {})
    }

    unfollowUser(username: string) {
        return this.http.post(`${this.BASE_URL}/unfollow?username=${username}`, {})
    }

    searchUser(query: string): Observable<User[]> {
        return this.http.get<User[]>(`${this.BASE_URL}/users/search/${query}`)
    }

    getAccount(accountId: string): Observable<Account | undefined> {
        return this.http.get<GetAccountResponse>(`${environment.GATEWAY_URL}/v1/accounts/${accountId}`)
            .pipe(map(res => res.account))
}

    getUser(userIdOrUsername: string): Observable<User | null> {
        return this.http.get<UserResponse>(`${environment.GATEWAY_URL}/v1/users/${userIdOrUsername}`)
            .pipe(map(res => toUser(res)))
    }

    promoteToBusiness(uid: string): Observable<Empty> {
        return this.http.post(`${environment.GATEWAY_URL}/v1/auths/business`, {
            user_id: uid
        })
    }

    deletePost(id: string): Observable<Empty> {
        return this.http.delete(`${environment.GATEWAY_URL}/v1/posts?id=${id}`)
    }

    deleteParty(id: string): Observable<any> {
        return this.http.delete(`${this.BASE_URL}/party/${id}`)
    }

    createPost(post: Post): Observable<PostResponse> {
        return this.http.post<PostResponse>(`${environment.GATEWAY_URL}/v1/posts`, {
            post: {
                caption: post.caption
            }
        })
    }

    createToken(token: string): Observable<CreateRegistrationTokenResponse> {
        return this.http.post<CreateRegistrationTokenResponse>(`${environment.GATEWAY_URL}/v1/notifications/token`, {
            token: token,
        })
    }

    interestParty(partyId: string): Observable<any> {
        return this.http.post(`${this.BASE_URL}/party/interest?partyId=${partyId}`, {})
    }

    uninterestParty(partyId: string): Observable<any> {
        return this.http.post(`${this.BASE_URL}/party/uninterest?partyId=${partyId}`, {})
    }

    getUserPosts(username: string): Observable<PostResponse[]> {
        return this.http.get<ListPostResponse>(`${environment.GATEWAY_URL}/v1/posts/list?username=${username}&pageSize=10&page=0`)
            .pipe(map(r => r.posts))
    }

    getPostFeed(): Observable<PostResponse[]> {
        return this.http.get<FeedPostResponse>(`${environment.GATEWAY_URL}/v1/posts/feed?pageSize=10&page=0`)
            .pipe(map(r => r.posts))
    }

    getStoryFeed(): Observable<Story[]> {
        return this.http.get<Story[]>(`${this.BASE_URL}/stories/feed?pageSize=10&page=0`)
    }

    seenStory(storyId: string) {
        return this.http.post(`${this.BASE_URL}/stories/${storyId}/seen`, {})
    }
}
