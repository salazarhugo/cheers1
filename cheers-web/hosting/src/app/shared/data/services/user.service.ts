import {Injectable} from '@angular/core';
import {map, Observable, ReplaySubject} from "rxjs";
import {ApiService} from "./api.service";
import {toUser, toUserFromUserItem, UserModel} from "../models/user.model";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {HttpClient} from "@angular/common/http";
import {Chat, toChat} from "../models/chat.model";
import {CreateRoomResponse} from "../../../../gen/ts/cheers/chat/v1/chat_service";
import {environment} from "../../../../environments/environment";
import {
    CreateUserRequest, CreateUserResponse,
    ListFollowersResponse,
    ListFollowingResponse
} from "../../../../gen/ts/cheers/user/v1/user_service";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";

@Injectable({
    providedIn: 'root'
})
export class UserService {

    private userSubject = new ReplaySubject<UserModel>(1);
    user$: Observable<UserModel> = this.userSubject.asObservable();

    constructor(
        private http: HttpClient,
        private api: ApiService,
    ) {
    }

    listFollowers(userId: string): Observable<UserModel[]> {
        return this.http.get<ListFollowersResponse>(`${environment.GATEWAY_URL}/v1/users/${userId}/followers`)
            .pipe(map(res => res.users.map(user => toUserFromUserItem(user))))
    }

    listFollowing(userId: string): Observable<UserModel[]> {
        return this.http.get<ListFollowingResponse>(`${environment.GATEWAY_URL}/v1/users/${userId}/following`)
            .pipe(map(res => res.users.map(user => toUserFromUserItem(user))))
    }

    getUserTickets() {
        return this.api.getUserTickets()
    }

    updateUser(user: UserModel) {
        return this.api.updateUser(user)
    }

    followUser(username: string) {
        return this.api.followUser(username)
    }

    unfollowUser(username: string) {
        return this.api.unfollowUser(username)
    }

    searchUser(query: string) {
        return this.api.searchUser(query)
    }

    refreshCurrentUser() {
        const userId = JSON.parse(localStorage.getItem('user')!).uid;
        // this.getUser(userId).subscribe(user => this.userSubject.next(user))
    }

    setUser(user: UserModel) {
        this.userSubject.next(user)
    }

    getUser(userIdOrUsername: string) {
        return this.api.getUser(userIdOrUsername)
    }

    usernameAvailable(username: string): Observable<boolean> {
        return this.api.usernameAvailability(username)
    }

    register(userReq: any): Observable<any> {
        return this.http.post<CreateUserResponse>(`${environment.GATEWAY_URL}/v1/users`, userReq)
    }
}
