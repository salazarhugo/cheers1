import {Injectable} from '@angular/core';
import {HttpClient, HttpErrorResponse} from "@angular/common/http";
import {catchError, map, Observable, of, throwError} from "rxjs";
import {StoryState, toUser, UserModel, UserResponse} from "../models/user.model";
import {PostModel} from "../models/post.model";
import {Story, toStory} from "../models/story.model";
import {Ticket, toTicket} from "../models/ticket.model";
import {
    FeedPostResponse,
    ListMapPostResponse,
    ListPostResponse,
    PostResponse
} from "../../../../gen/ts/cheers/post/v1/post_service";
import {Empty} from "../../../../gen/ts/google/protobuf/empty";
import {environment} from "../../../../environments/environment";
import {CreateRegistrationTokenResponse} from "../../../../gen/ts/cheers/notification/v1/notification_service";
import {Account, GetAccountResponse} from "../../../../gen/ts/cheers/account/v1/account_service";
import {ListUserOrdersResponse, Order} from "../../../../gen/ts/cheers/order/v1/order_service";
import {ListTicketResponse} from "../../../../gen/ts/cheers/ticket/v1/ticket";
import {FeedStoryResponse} from "../../../../gen/ts/cheers/story/v1/story_service";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";
import {SearchUserResponse} from "../response/search-user.response";
import {
    BeginLoginResponse,
    FinishLoginRequest,
    FinishLoginResponse,
    PublicKeyCredentialOptions
} from "../models/webauthn.models";

@Injectable({
    providedIn: 'root'
})
export class ApiService {

    BASE_URL = "https://rest-api-r3a2dr4u4a-nw.a.run.app"

    constructor(
        private http: HttpClient,
    ) {
    }

    beginLogin(username: string): Observable<BeginLoginResponse | null> {
        return this.http.get<BeginLoginResponse>(`${environment.GATEWAY_URL}/v1/auth/login/begin/${username}`)
            .pipe(catchError(this.handleError))
    }

    finishLogin(request: FinishLoginRequest): Observable<FinishLoginResponse> {
        return this.http.post<FinishLoginResponse>(`${environment.GATEWAY_URL}/v1/auth/login/finish`, request)
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

    register(userReq: any): Observable<UserModel> {
        return this.http.post<UserModel>(`${environment.GATEWAY_URL}/users/create`, userReq)
    }

    likePost(postId: string) {
        return this.http.post(`${environment.GATEWAY_URL}/v1/posts/${postId}/like`, {})
    }

    unlikePost(postId: string) {
        return this.http.post(`${environment.GATEWAY_URL}/v1/posts/${postId}/unlike`, {})
    }

    updateUser(user: UserModel) {
        return this.http.patch(`${environment.GATEWAY_URL}/v1/users`, user)
    }

    followUser(username: string) {
        return this.http.post(`${environment.GATEWAY_URL}/v1/users/${username}/follow`, {})
    }

    unfollowUser(username: string) {
        return this.http.delete(`${environment.GATEWAY_URL}/v1/users/${username}/unfollow`)
    }

    searchUser(query: string): Observable<UserModel[]> {
        return this.http.get<SearchUserResponse>(`${environment.GATEWAY_URL}/v1/users/search/${query}`)
            .pipe(map(response => response.users))
    }

    getAccount(accountId: string): Observable<Account | undefined> {
        return this.http.get<GetAccountResponse>(`${environment.GATEWAY_URL}/v1/accounts/${accountId}`)
            .pipe(map(res => res.account))
}

    getUser(userIdOrUsername: string): Observable<UserModel | null> {
        return this.http.get<UserResponse>(`${environment.GATEWAY_URL}/v1/users/${userIdOrUsername}`)
            .pipe(
                map(res => toUser(res)),
                catchError(this.handleError),
            )
    }

    private handleError(error: HttpErrorResponse) {
        if (error.status === 404) {
            return of(null);
        }
        // Return an observable with a user-facing error message.
        return throwError(() => new Error('Something bad happened; please try again later.'));
    }

    verifyUser(uid: string): Observable<Empty> {
        return this.http.post(`${environment.GATEWAY_URL}/v1/auths/verify`, {
            user_id: uid
        })
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

    createPost(post: PostModel): Observable<PostResponse> {
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

    getUserPosts(username: string): Observable<PostResponse[]> {
        return this.http.get<ListPostResponse>(`${environment.GATEWAY_URL}/v1/posts/list?username=${username}&pageSize=10&page=0`)
            .pipe(map(r => r.posts))
    }

    listMapPost(): Observable<PostResponse[]> {
        return this.http.get<ListMapPostResponse>(`${environment.GATEWAY_URL}/v1/posts/map?pageSize=10&page=0`)
            .pipe(map(r => r.posts))
    }

    getPostFeed(): Observable<PostResponse[]> {
        return this.http.get<FeedPostResponse>(`${environment.GATEWAY_URL}/v1/posts/feed?pageSize=10&page=1`)
            .pipe(
                map(r => r.posts.sort((a, b) => b.post?.createTime! - a.post?.createTime!)),
            )
    }

    getStoryFeed(): Observable<Story[]> {
        return this.http.get<FeedStoryResponse>(`${environment.GATEWAY_URL}/v1/stories/feed?pageSize=10&page=0`)
            .pipe(
                map(r => r.items.flatMap(uws => toStory(uws)))
            )
    }

    seenStory(storyId: string) {
        return this.http.post(`${this.BASE_URL}/stories/${storyId}/seen`, {})
    }
}
