import {Injectable} from '@angular/core';
import {Observable, ReplaySubject} from "rxjs";
import {ApiService} from "./api.service";
import {User} from "../models/user.model";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";

@Injectable({
    providedIn: 'root'
})
export class UserService {

    private userSubject = new ReplaySubject<User>(1);
    user$: Observable<User> = this.userSubject.asObservable();

    constructor(
        private api: ApiService,
    ) {
    }

    getUserTickets() {
        return this.api.getUserTickets()
    }

    unfollowUser(username: string) {
        return this.api.unfollowUser(username)
    }

    updateUser(user: User) {
        return this.api.updateUser(user)
    }

    followUser(username: string) {
        return this.api.followUser(username)
    }

    searchUser(query: string) {
        return this.api.searchUser(query)
    }

    refreshCurrentUser() {
        const userId = JSON.parse(localStorage.getItem('user')!).uid;
        this.getUser(userId).subscribe(user => this.userSubject.next(user))
    }

    setUser(user: User) {
        this.userSubject.next(user)
    }

    getPostFeed(): Observable<PostResponse[]> {
        return this.api.getPostFeed()
    }

    getUser(userIdOrUsername: string) {
        return this.api.getUser(userIdOrUsername)
    }

    usernameAvailable(username: string): Observable<boolean> {
        return this.api.usernameAvailability(username)
    }

    register(userReq: any): Observable<User> {
        return this.api.register(userReq)
    }
}
