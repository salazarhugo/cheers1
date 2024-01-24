import {Injectable} from '@angular/core';
import {ActivatedRouteSnapshot, Resolve, Router} from "@angular/router";
import {UserService} from '../../shared/data/services/user.service';
import {first} from "rxjs/operators";
import {UserModel} from "../../shared/data/models/user.model";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {ApiService} from "../../shared/data/services/api.service";
import {EMPTY, firstValueFrom, throwError} from "rxjs";
import {Chat} from "../../shared/data/models/chat.model";
import {ChatService} from "./chat.service";

@Injectable({
    providedIn: 'root'
})
export class ChatResolver implements Resolve<Chat> {

    constructor(
        public chatService: ChatService,
        private router: Router,
        private afAuth: AngularFireAuth,
        private api: ApiService,
    ) {
    }

    resolve(route: ActivatedRouteSnapshot): Promise<Chat> {
        return new Promise(async (resolve, reject) => {

            const authUser = await firstValueFrom(this.afAuth.authState)

            // Check if user is signed in
            if (!authUser) {
                await this.router.navigate(['sign-in']);
                return reject();
            }

            const user = await firstValueFrom(this.api.getUser(authUser.uid))

            // Check if user document exists
            if (user) {
                this.chatService.joinRoom("room01")
                resolve(user)
            } else {
                console.log("Chat doesn't exist")
                await this.router.navigate(['finish-sign-up']);
                return reject();
            }
        })
    }
}
