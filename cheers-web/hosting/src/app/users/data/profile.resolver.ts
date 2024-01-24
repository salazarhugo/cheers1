import { Injectable } from '@angular/core';
import { Router, RouterStateSnapshot, ActivatedRouteSnapshot } from '@angular/router';
import {firstValueFrom, Observable, of} from 'rxjs';
import {UserModel} from "../../shared/data/models/user.model";
import {UserService} from "../../shared/data/services/user.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {ApiService} from "../../shared/data/services/api.service";

@Injectable({
    providedIn: 'root'
})
export class ProfileResolver  {

    constructor(
        public userService: UserService,
        private router: Router,
        private afAuth: AngularFireAuth,
        private api: ApiService,
    ) {
    }

    resolve(route: ActivatedRouteSnapshot): Promise<UserModel> {
        return new Promise(async (resolve, reject) => {
            // let user = await firstValueFrom(this.api.getUser(authUser.uid))
            // resolve(user)
        })
    }
}
